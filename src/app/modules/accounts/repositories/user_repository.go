package repositories

import (
	"base-fiber-api/src/app/modules/accounts/interfaces"
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/shared/pkg"
	"base-fiber-api/src/app/shared/scopes"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserGorm struct {
	db *gorm.DB
}

func (u UserGorm) List(pagination pkg.Pagination) (*pkg.Pagination, error) {
	var users models.Users

	if err := u.db.Scopes(scopes.Paginate(users, &pagination, u.db)).Find(&users).Error; err != nil {
		return nil, err
	}
	pagination.Data = users.PublicUsers()

	return &pagination, nil
}

func (u UserGorm) Get(id string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserGorm) Store(user *models.User) (*models.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserGorm) Edit(user *models.User) (*models.User, error) {
	if err := u.db.Clauses(clause.Returning{}).Where("id = ?", user.Id).Updates(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserGorm) Delete(user *models.User) error {
	u.db.Where("id = ?", user.Id).Updates(&user)
	if err := u.db.Where("id = ?", user.Id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserGorm) FindBy(field string, value string) (*models.User, error) {
	var user models.User
	if err := u.db.Where(field+" = ?", value).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserGorm) FindManyBy(field []string, value string) (*models.User, error) {
	var user models.User
	for _, f := range field {
		u.db.Where(f+" = ?", value).First(&user)
		if user.Id != "" {
			return &user, nil
		}
	}
	return nil, errors.New("record not found")
}

func UserRepository(db *gorm.DB) *UserGorm {
	return &UserGorm{db}
}

var _ interfaces.UserInterface = &UserGorm{}
