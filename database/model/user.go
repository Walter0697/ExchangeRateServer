package model

import "gorm.io/gorm"

type User struct {
	ObjectBase
	Username    string `json:"username"`
	Password    string `json:"-"`
	Role        string `json:"role"`
	IsActivated bool   `json:"is_activated"`
}

func (user *User) Create(db *gorm.DB) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (user *User) Update(db *gorm.DB) error {
	if err := db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (user *User) GetUserByUsername(db *gorm.DB) error {
	if err := db.Where("username = ?", user.Username).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (user *User) GetUserById(db *gorm.DB) error {
	if err := db.Where("id = ?", user.ID).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (user *User) AnyUserExist(db *gorm.DB) (int64, error) {
	var count int64
	if err := db.Model(user).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
