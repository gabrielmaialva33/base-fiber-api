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

type RoleRepo struct {
	db *gorm.DB
}

// NewRoleRepository returns a new instance of RoleRepo
func NewRoleRepository(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db}
}

// RoleRepo implements interfaces.RoleInterface
var _ interfaces.RoleInterface = &RoleRepo{}

func (r *RoleRepo) List(meta pkg.Meta) (*pkg.Pagination, error) {
	var roles models.Roles
	var fields = []string{"name", "slug", "description"}
	var pagination pkg.Pagination

	if err := r.db.Scopes(scopes.Paginate(roles, fields, &meta, r.db)).Find(&roles).Error; err != nil {
		return nil, err
	}

	pagination.SetMeta(meta)
	pagination.SetData(roles.PublicRoles())

	return &pagination, nil
}

func (r *RoleRepo) Get(id string) (*models.Role, error) {
	var role models.Role
	if err := r.db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepo) Store(role *models.Role) (*models.Role, error) {
	if err := r.db.Create(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepo) Edit(role *models.Role) (*models.Role, error) {
	if err := r.db.Clauses(clause.Returning{}).Where("id = ?", role.Id).Updates(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepo) Delete(role *models.Role) error {
	r.db.Where("id = ?", role.Id).Updates(&role)
	if err := r.db.Where("id = ?", role.Id).Delete(&role).Error; err != nil {
		return err
	}
	return nil
}

func (r *RoleRepo) FindBy(field string, value string) (*models.Role, error) {
	var role models.Role
	if err := r.db.Where(field+" = ?", value).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepo) FindManyBy(field []string, value string) (*models.Role, error) {
	var role models.Role
	for _, f := range field {
		r.db.Where(f+" = ?", value).First(&role)
		if role.Id != "" {
			return &role, nil
		}
	}
	return nil, errors.New("user not found")
}
