package models

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	gorm.Model

	Id        string         `gorm:"primaryKey;default:uuid_generate_v4()" json:"id"`
	Name      string         `gorm:"column:name;size:80;not null;unique;index;" json:"name" validate:"required,min=2,max=80,unique,omitempty"`
	Slug      string         `gorm:"column:slug;size:80;not null;unique;" json:"slug" validate:"required,min=2,max=80,unique,omitempty"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index,default:null" json:"-"`

	// Relations
	Users []*User `gorm:"many2many:user_roles;" json:"users"`
}

type Roles []Role

type RolePublic struct {
	Id   string `json:"id"`
	Slug string `json:"slug"`
}

func (r *Role) PublicRole() interface{} {
	return &RolePublic{
		Id:   r.Id,
		Slug: r.Slug,
	}
}

func (roles Roles) PublicRoles() []interface{} {
	result := make([]interface{}, len(roles))
	for i, role := range roles {
		result[i] = role.PublicRole()
	}
	return result
}
