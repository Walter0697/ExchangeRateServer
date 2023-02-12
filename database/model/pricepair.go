package model

import "gorm.io/gorm"

type PricePair struct {
	ObjectBase
	Crypto   string  `json:"crypto"`
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

func (pp *PricePair) GetOrCreate(db *gorm.DB) error {
	if err := db.Where("crypto = ? AND currency = ?", pp.Crypto, pp.Currency).FirstOrCreate(pp).Error; err != nil {
		return err
	}

	return nil
}
