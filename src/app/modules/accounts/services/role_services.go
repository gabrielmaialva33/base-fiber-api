package services

import (
	"base-fiber-api/src/app/modules/accounts/interfaces"
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/shared/pkg"
)

type RoleRepository struct {
	rr interfaces.RoleInterface
}

func (r RoleRepository) List(pagination pkg.Pagination) (*pkg.Pagination, error) {
	return r.rr.List(pagination)
}

func (r RoleRepository) Get(id string) (*models.Role, error) {
	return r.rr.Get(id)
}

func (r RoleRepository) Store(role *models.Role) (*models.Role, error) {
	return r.rr.Store(role)
}

func (r RoleRepository) Edit(role *models.Role) (*models.Role, error) {
	return r.rr.Edit(role)
}

func (r RoleRepository) Delete(role *models.Role) error {
	return r.rr.Delete(role)
}

func (r RoleRepository) FindBy(field string, value string) (*models.Role, error) {
	return r.rr.FindBy(field, value)
}

func (r RoleRepository) FindManyBy(field []string, value string) (*models.Role, error) {
	return r.rr.FindManyBy(field, value)
}

var _ interfaces.RoleInterface = &RoleRepository{}
