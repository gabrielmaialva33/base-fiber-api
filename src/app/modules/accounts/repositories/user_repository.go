package repositories

import (
	"base-fiber-api/src/app/modules/accounts/interfaces"
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/shared/pkg"
	"base-fiber-api/src/app/shared/scopes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Gorm struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *Gorm {
	return &Gorm{db}
}

var _ interfaces.UserInterface = &Gorm{}

func (g Gorm) List(pagination pkg.Pagination) (*pkg.Pagination, error) {
	var users models.Users

	if err := g.db.Scopes(scopes.Paginate(users, &pagination, g.db)).Find(&users).Error; err != nil {
		return nil, err
	}
	pagination.Data = users.PublicUsers()

	return &pagination, nil
}

func (g Gorm) Get(id string) (*models.User, error) {
	var user models.User
	if err := g.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (g Gorm) Store(user *models.User) (*models.User, error) {
	if err := g.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (g Gorm) Edit(id string, model *models.User) (*models.User, error) {
	if err := g.db.Clauses(clause.Returning{}).Where("id = ?", id).Updates(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (g Gorm) Delete(id string, model *models.User) error {
	g.db.Where("id = ?", id).Updates(&model)
	if err := g.db.Where("id = ?", id).Delete(&model).Error; err != nil {
		return err
	}
	return nil
}

func (g Gorm) FindBy(field string, value string) (*models.User, error) {
	var user models.User
	if err := g.db.Where(field+" = ?", value).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
