package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        string         `gorm:"primary_key;default:uuid_generate_v4()" json:"id"`
	FirstName string         `gorm:"column:first_name;size:80;not null" json:"first_name"`
	LastName  string         `gorm:"column:last_name;size:80;not null" json:"last_name"`
	Email     string         `gorm:"column:email;size:255;not null;unique;unique;index;" json:"email"`
	UserName  string         `gorm:"column:user_name;size:80;not null;unique;index;" json:"user_name"`
	Password  string         `gorm:"column:password;size:255;not null" json:"-"`
	IsDeleted bool           `gorm:"column:is_deleted;not null;default:false" json:"-"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;default:null" json:"-"`
}
