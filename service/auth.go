package service

import (
	"chaos/backend/database"
	"chaos/backend/database/model"
	"chaos/backend/helper"
	"chaos/backend/utility"
)

type UnauthorizationError struct{}

func (m *UnauthorizationError) Error() string {
	return "wrong username or password"
}

// register a new user
func RegisterUser(username, password string) (*model.User, error) {
	var user model.User

	user.Username = username
	user.Password = utility.GetEncrpytPassword(password)
	user.IsActivated = true
	user.CreatedBy = nil
	user.UpdatedBy = nil

	if err := user.Create(database.Connection); err != nil {
		return nil, err
	}

	return &user, nil
}

// login function
func Login(username, password string) (string, error) {
	var user model.User
	user.Username = username

	if err := user.GetUserByUsername(database.Connection); err != nil {
		if utility.RecordNotFound(err) {
			return "", &UnauthorizationError{}
		}
		return "", err
	}

	if !utility.ComparePassword(user.Password, password) {
		return "", &UnauthorizationError{}
	}

	jwtToken, err := helper.GenerateToken(user.Username, user.ID)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

// validate token for middleware
func ValidateToken(token string) *model.User {
	jwtInfo, err := helper.ParseToken(token)
	if err != nil {
		return nil
	}

	var user model.User
	user.Username = jwtInfo.Username
	if err := user.GetUserByUsername(database.Connection); err != nil {
		return nil
	}

	return &user
}
