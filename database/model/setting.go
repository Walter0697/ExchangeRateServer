package model

import (
	"gorm.io/gorm"
)

type Setting struct {
	ObjectBase
	Identifier string  `json:"identifier"`
	StrValue   string  `json:"str_value"`
	BoolValue  bool    `json:"bool_value"`
	IntValue   uint    `json:"int_value"`
	FloatValue float64 `json:"float_value"`
}

func (s *Setting) GetOrCreate(db *gorm.DB) error {
	if err := db.Where("identifier = ?", s.Identifier).FirstOrCreate(s).Error; err != nil {
		return err
	}

	return nil
}

func (s *Setting) GetByIdentifier(db *gorm.DB) error {
	if err := db.Where("identifier = ?", s.Identifier).First(s).Error; err != nil {
		return err
	}

	return nil
}

func (s *Setting) Update(db *gorm.DB) error {
	if err := db.Save(s).Error; err != nil {
		return err
	}

	return nil
}
