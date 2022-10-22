package models

import (
	"base-fiber-api/src/app/shared/pkg"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	Id              string         `gorm:"primaryKey;default:uuid_generate_v4();" json:"id"`
	FirstName       string         `gorm:"column:first_name;size:80;not null;" json:"first_name" validate:"required,min=2,max=80,alpha,omitempty"`
	LastName        string         `gorm:"column:last_name;size:80;not null;" json:"last_name" validate:"required,min=2,max=80,alpha,omitempty"`
	Email           string         `gorm:"column:email;size:255;not null;unique;unique;index;" json:"email" validate:"required,email,max=247,unique,omitempty"`
	UserName        string         `gorm:"column:user_name;size:58;not null;unique;index;" json:"user_name" validate:"required,min=4,max=50,unique,omitempty"`
	Password        string         `gorm:"column:password;size:255;not null;" json:"password" form:"password" validate:"required,min=6,max=50,omitempty"`
	ConfirmPassword string         `gorm:"-" json:"confirm_password" form:"confirm_password" validate:"required,min=6,max=50,eqfield=Password,omitempty"`
	CreatedAt       time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;" json:"-"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;" json:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;index,default:null;" json:"-"`

	// Relations
	Roles []*Role `gorm:"many2many:user_roles;" json:"roles"`
}
type Users []User

type UserPublic struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	UserName  string `json:"user_name"`
}

// PublicUser returns a public representation of a user.
func (u *User) PublicUser() interface{} {
	return &UserPublic{
		Id:        u.Id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		UserName:  u.UserName,
	}
}

// PublicUsers returns a public representation of a user.
func (users Users) PublicUsers() []interface{} {
	result := make([]interface{}, len(users))
	for i, user := range users {
		result[i] = user.PublicUser()
	}
	return result
}

// BeforeSave hook executed before saving a User to the database.
func (u *User) BeforeSave(*gorm.DB) error {
	hash, err := pkg.CreateHash(u.Password, pkg.DefaultParams)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}

type Login struct {
	Uid      string `json:"uid" validate:"required"`
	Password string `json:"password" validate:"required"`
}
