package seed

import (
	"chaos/backend/config"
	"chaos/backend/database"
	"chaos/backend/database/model"
	"chaos/backend/utility"
)

func SeedDefaultUser() {
	var user model.User
	user_count, err := user.AnyUserExist(database.Connection)
	if err != nil {
		panic(err)
	}

	// only seed the first default user when there is no user in the database
	if user_count != 0 {
		return
	}

	user.Username = config.Data.Default.AdminUsername
	user.Password = utility.GetEncrpytPassword(config.Data.Default.AdminPassword)
	user.IsActivated = true
	user.CreatedBy = nil
	user.UpdatedBy = nil

	if err := user.Create(database.Connection); err != nil {
		panic(err)
	}
}
