package services

import (
	"base-fiber-api/src/app/modules/accounts/interfaces"
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/shared/pkg"
)

type UserRepository struct {
	ur interfaces.UserInterface
}

func (u UserRepository) List(meta pkg.Meta) (*pkg.Pagination, error) {
	return u.ur.List(meta)
}

func (u UserRepository) Get(id string) (*models.User, error) {
	return u.ur.Get(id)
}

func (u UserRepository) Store(user *models.User) (*models.User, error) {
	return u.ur.Store(user)
}

func (u UserRepository) Edit(user *models.User) (*models.User, error) {
	return u.ur.Edit(user)
}

func (u UserRepository) Delete(user *models.User) error {
	return u.ur.Delete(user)
}

func (u UserRepository) FindBy(field string, value string) (*models.User, error) {
	return u.ur.FindBy(field, value)
}

func (u UserRepository) FindManyBy(field []string, value string) (*models.User, error) {
	return u.ur.FindManyBy(field, value)
}

var _ interfaces.UserInterface = &UserRepository{}
