package controller

import (
	"chaos/backend/database"
	"chaos/backend/database/model"
	"fmt"
	"net/http"
	"time"
)

func disableFetch() {
	var setting model.Setting
	setting.Identifier = "fetch_enable"
	if err := setting.GetOrCreate(database.Connection); err != nil {
		fmt.Println(err)
		return
	}

	setting.BoolValue = false
	if err := setting.Update(database.Connection); err != nil {
		fmt.Println(err)
		return
	}
}

func EnableFetch(w http.ResponseWriter, r *http.Request) {
	var setting model.Setting
	setting.Identifier = "fetch_enable"
	if err := setting.GetByIdentifier(database.Connection); err != nil {
		ErrorResp(w, r, http.StatusInternalServerError, err)
		return
	}

	if !setting.BoolValue {
		time.AfterFunc(30*time.Minute, disableFetch)
		setting.BoolValue = true
		if err := setting.Update(database.Connection); err != nil {
			ErrorResp(w, r, http.StatusInternalServerError, err)
			return
		}
	}

	var result BasicRespond
	result.Message = "ok"

	JSON(w, r, 200, &result)
}
