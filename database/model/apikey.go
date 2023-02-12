package model

import (
	"time"

	"gorm.io/gorm"
)

type APIKey struct {
	ObjectBase
	Identifier string    `json:"identifier"`
	Token      string    `json:"token"`
	IsEnabled  bool      `json:"is_enabled"`
	ExpiredAt  time.Time `json:"expiredAt"`
}

func (key *APIKey) Create(db *gorm.DB) error {
	if err := db.Create(key).Error; err != nil {
		return err
	}

	return nil
}

func (key *APIKey) Update(db *gorm.DB) error {
	if err := db.Save(key).Error; err != nil {
		return err
	}

	return nil
}

func (key *APIKey) GetKeyByToken(db *gorm.DB) error {
	if err := db.Where("token = ?", key.Token).First(key).Error; err != nil {
		return err
	}

	return nil
}
