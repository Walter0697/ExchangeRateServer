package controller

import (
	"chaos/backend/service"
	"encoding/json"
	"errors"
	"net/http"
)

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRespond struct {
	BaseRespond
	Username string `json:"username"`
	Token    string `json:"token"`
}

// login with username and password
func Login(w http.ResponseWriter, r *http.Request) {
	var body LoginBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		ErrorResp(w, r, http.StatusBadRequest, err)
		return
	}

	token, err := service.Login(body.Username, body.Password)
	if err != nil {
		if errors.Is(err, &service.UnauthorizationError{}) {
			ErrorResp(w, r, http.StatusUnauthorized, err)
		} else {
			ErrorResp(w, r, http.StatusInternalServerError, err)
		}
		return
	}

	var result LoginRespond
	result.Token = token
	result.Username = body.Username

	JSON(w, r, 200, &result)
}
