package database

import (
	"base-fiber-api/src/app/modules/accounts/interfaces"
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/modules/accounts/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repositories struct {
	User interfaces.UserInterface
	db   *gorm.DB
}

var DB *gorm.DB

func NewRepositories(dsn string) *Repositories {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(" -> Could not connect to the database")
	}

	DB = database
	
	return &Repositories{
		User: repositories.UserRepository(database),
		db:   database,
	}
}

func (r Repositories) Migrate() {
	r.db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	r.db.AutoMigrate(&models.User{})
}

func (r Repositories) Drop() {
	r.db.Exec("DROP EXTENSION IF EXISTS \"uuid-ossp\";")
	r.db.Migrator().DropTable(&models.User{})
}

func (r Repositories) Seed() {
	//TODO implement me
}
