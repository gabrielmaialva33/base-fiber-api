package repositories

import (
	"base-fiber-api/src/app/modules/accounts/interfaces"
	"base-fiber-api/src/app/modules/accounts/models"
	"gorm.io/gorm"
)

type Gorm struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *Gorm {
	return &Gorm{db}
}

var _ interfaces.UserInterface = &Gorm{}

func (g Gorm) List() ([]models.User, error) {
	//TODO implement me
	panic("implement me")
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
