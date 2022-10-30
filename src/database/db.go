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
	Role interfaces.RoleInterface
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
		Role: repositories.RoleRepository(database),
		db:   database,
	}
}

func (r Repositories) Migrate() {
	r.db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	r.db.AutoMigrate(&models.User{}, &models.Role{})
}

func (r Repositories) Drop() {
	r.db.Exec("DROP EXTENSION IF EXISTS \"uuid-ossp\";")
	r.db.Migrator().DropTable(&models.User{}, &models.Role{})
}

func (r Repositories) Seed() {
	roles := []models.Role{
		{
			Name:        "root",
			Slug:        "Root",
			Description: "A root user has all permissions",
		},
		{
			Name:        "admin",
			Slug:        "Admin",
			Description: "An admin user has all permissions except root",
		},
		{
			Name:        "user",
			Slug:        "User",
			Description: "A user has limited permissions",
		},
		{
			Name:        "guest",
			Slug:        "Guest",
			Description: "A guest user has no permissions",
		},
	}
	users := []models.User{
		{
			FirstName: "Root",
			LastName:  "System",
			Email:     "root@go.com",
			UserName:  "root",
			Password:  "123456",
		},
		{
			FirstName: "Admin",
			LastName:  "System",
			Email:     "admin@go.com",
			UserName:  "admin",
			Password:  "123456",
		},
		{
			FirstName: "Gabriel",
			LastName:  "Maia",
			Email:     "maia@go.com",
			UserName:  "maia",
			Password:  "123456",
		},
	}

	r.db.Create(&roles)
	r.db.Create(&users)
}
