package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id              string         `gorm:"primary_key;default:uuid_generate_v4()" json:"id"`
	FirstName       string         `gorm:"column:first_name;size:80;not null" json:"first_name" validate:"required,min=2,max=80,alpha"`
	LastName        string         `gorm:"column:last_name;size:80;not null" json:"last_name" validate:"required,min=2,max=80,alpha"`
	Email           string         `gorm:"column:email;size:255;not null;unique;unique;index;" json:"email" validate:"required,email,max=255,unique"`
	UserName        string         `gorm:"column:user_name;size:50;not null;unique;index;" json:"user_name" validate:"required,min=4,max=50,unique"`
	Password        string         `gorm:"column:password;size:255;not null" json:"-" form:"password" validate:"required,min=6,max=50"`
	ConfirmPassword string         `gorm:"-" json:"-" form:"confirm_password" validate:"required,min=6,max=50,eqfield=Password"`
	IsDeleted       bool           `gorm:"column:is_deleted;not null;default:false" json:"-"`
	CreatedAt       time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;default:null" json:"-"`
}

type UserEdit struct {
	FirstName       string `gorm:"column:first_name;size:80;not null" json:"first_name" validate:"min=2,max=80,alpha"`
	LastName        string `gorm:"column:last_name;size:80;not null" json:"last_name" validate:"min=2,max=80,alpha"`
	Email           string `gorm:"column:email;size:255;not null;unique;unique;index;" json:"email" validate:"email,max=255,unique"`
	UserName        string `gorm:"column:user_name;size:50;not null;unique;index;" json:"user_name" validate:"min=4,max=50,unique"`
	Password        string `gorm:"column:password;size:255;not null" json:"-" form:"password" validate:"min=6,max=50"`
	ConfirmPassword string `gorm:"-" json:"-" form:"confirm_password" validate:"min=6,max=50,eqfield=Password"`
}
