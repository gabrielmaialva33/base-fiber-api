package database

import (
	"base-fiber-api/src/app/modules/accounts/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(" -> Could not connect to the database")
	}

	database.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	database.AutoMigrate(&models.User{})

	DB = database
}
