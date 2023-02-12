package database

import (
	"chaos/backend/config"
	"chaos/backend/database/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func Connect(host, port, db, username, password string) error {
	var err error
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Hong_Kong",
		host, username, password, db, port)

	Connection, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		return err
	}

	return nil
}

func Init() {
	err := Connect(
		config.Data.DB.Host,
		config.Data.DB.Port,
		config.Data.DB.Name,
		config.Data.DB.Username,
		config.Data.DB.Password,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func AutoMigration() {
	Connection.AutoMigrate(&model.User{})
	Connection.AutoMigrate(&model.APIKey{})
	Connection.AutoMigrate(&model.PricePair{})
	Connection.AutoMigrate(&model.Price{})
}

func Refresh() {
	Connection.Migrator().DropTable(&model.User{})
	Connection.Migrator().DropTable(&model.APIKey{})
	Connection.Migrator().DropTable(&model.PricePair{})
	Connection.Migrator().DropTable(&model.Price{})
}
