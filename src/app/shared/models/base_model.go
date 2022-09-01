package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        string         `gorm:"primary_key;default:uuid_generate_v4()" json:"id"`
	IsDeleted bool           `gorm:"column:is_deleted;not null;default:false" json:"-"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;default:null" json:"-"`
}
