package repositories

import (
	"base-fiber-api/src/app/modules/accounts/interfaces"
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/shared/pkg"
	"base-fiber-api/src/app/shared/scopes"
	"gorm.io/gorm"
)

type Gorm struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *Gorm {
	return &Gorm{db}
}

var _ interfaces.UserInterface = &Gorm{}

func (g Gorm) List(pagination pkg.Pagination) (*pkg.Pagination, error) {
	var users []models.User

	if err := g.db.Scopes(scopes.Paginate(users, &pagination, g.db)).Find(&users).Error; err != nil {
		return nil, err
	}
	pagination.Data = users

	return &pagination, nil
}

func (g Gorm) Get(id string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (g Gorm) Store(user *models.User) (*models.User, error) {
	if err := g.db.Debug().Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (g Gorm) Edit(model *models.User) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (g Gorm) Delete(model *models.User) error {
	//TODO implement me
	panic("implement me")
}
