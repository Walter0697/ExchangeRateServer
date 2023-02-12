package model

import (
	"time"

	"gorm.io/gorm"
)

type Price struct {
	ObjectBase
	Value  float64   `json:"value"`
	Pair   PricePair `gorm:"foreignKey:pair_id;references:id"`
	PairId uint
}

func (p *Price) QueryByPair(db *gorm.DB) *gorm.DB {
	return db.Where("pair_id = ?", p.PairId)
}

func (p *Price) Create(db *gorm.DB) error {
	if err := db.Create(p).Error; err != nil {
		return err
	}

	return nil
}

func (p *Price) FindByLatest(db *gorm.DB) error {
	base := p.QueryByPair(db)
	if err := base.Last(p).Error; err != nil {
		return err
	}

	return nil
}

func (p *Price) FindPreviousByTime(db *gorm.DB) error {
	base := p.QueryByPair(db)
	if err := base.Where("created_at <= ?", p.CreatedAt.Format(time.RFC3339)).Last(p).Error; err != nil {
		return err
	}

	return nil
}

func (p *Price) FindNextByTime(db *gorm.DB) error {
	base := p.QueryByPair(db)
	if err := base.Where("created_at >= ?", p.CreatedAt.Format(time.RFC3339)).First(p).Error; err != nil {
		return err
	}

	return nil
}
