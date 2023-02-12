package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type ObjectBase struct {
	BaseModel
	CreatedBy  *User `gorm:"foreignKey:created_uid;references:id"`
	CreatedUID *uint
	UpdatedAt  time.Time `json:"updatedAt"`
	UpdatedBy  *User     `gorm:"foreignKey:updated_uid;references:id"`
	UpdatedUID *uint
}
