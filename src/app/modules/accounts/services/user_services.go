package services

import (
	"base-fiber-api/src/app/modules/accounts/interfaces"
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/shared/pkg"
)

type UserRepository struct {
	ur interfaces.UserInterface
}

func (u UserRepository) List(pagination pkg.Pagination) (*pkg.Pagination, error) {
	return u.ur.List(pagination)
}

func (u UserRepository) Get(id string) (*models.User, error) {
	return u.ur.Get(id)
}

func (u UserRepository) Store(user *models.User) (*models.User, error) {
	return u.ur.Store(user)
}

func (u UserRepository) Edit(id string, user *models.User) (*models.User, error) {
	return u.ur.Edit(id, user)
}

func (u UserRepository) Delete(id string, user *models.User) error {
	return u.ur.Delete(id, user)
}

var _ interfaces.UserInterface = &UserRepository{}
