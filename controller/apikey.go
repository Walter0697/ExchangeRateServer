package controller

import (
	"chaos/backend/database"
	"chaos/backend/database/model"
	"chaos/backend/middleware"
	"chaos/backend/utility"
	"encoding/json"
	"net/http"
)

type GenerateAPIBody struct {
	Identifier string `json:"apikey"`
}

type APIKeyResultRespond struct {
	BaseRespond
	Token string `json:"token"`
}

// generate api key by user
// Authorization: user
func GenerateAPIkey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.ForContext(ctx)

	var body GenerateAPIBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		ErrorResp(w, r, http.StatusBadRequest, err)
		return
	}

	token := utility.RandStringRunes(20)
	var apikey model.APIKey
	apikey.Identifier = body.Identifier
	apikey.Token = token
	apikey.IsEnabled = true
	apikey.CreatedBy = user
	apikey.UpdatedBy = user

	if err := apikey.Create(database.Connection); err != nil {
		ErrorResp(w, r, http.StatusInternalServerError, err)
		return
	}

	result := APIKeyResultRespond{
		Token: apikey.Token,
	}

	JSON(w, r, 200, &result)
}
