package seed

import (
	"chaos/backend/config"
	"chaos/backend/database"
	"chaos/backend/database/model"
)

func SeedDefaultSetting() {
	var apiEnableSetting model.Setting
	apiEnableSetting.Identifier = "fetch_enable"
	if config.Data.App.Environment == "production" {
		apiEnableSetting.BoolValue = true
	} else {
		apiEnableSetting.BoolValue = false
	}

	if err := apiEnableSetting.GetOrCreate(database.Connection); err != nil {
		panic(err)
	}
}
