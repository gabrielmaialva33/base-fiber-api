package models

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	Id        string         `gorm:"primary_key;default:uuid_generate_v4()" json:"id"`
	Name      string         `gorm:"column:name;size:80;not null;unique;index;" json:"name" validate:"required,min=2,max=80,unique,omitempty"`
	Slug      string         `gorm:"column:slug;size:80;not null;unique;index;" json:"slug" validate:"required,min=2,max=80,unique,omitempty"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index,default:null" json:"-"`
}
