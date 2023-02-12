package test

import (
	"chaos/backend/database"
	"chaos/backend/database/model"
	"math/rand"
	"time"
)

func GenerateRandomPrice(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func GenerateRandomPriceList(num int) ([]model.Price, error) {
	var priceList []model.Price

	crypto := "BTC"
	currency := "USD"

	var pricePair model.PricePair
	pricePair.Crypto = crypto
	pricePair.Currency = currency
	if err := pricePair.GetOrCreate(database.Connection); err != nil {
		return priceList, err
	}

	now := time.Now()
	current := now.Add(-time.Minute * time.Duration(num+1))
	for i := 0; i < num; i++ {
		value := GenerateRandomPrice(21500, 22000)
		var price model.Price
		price.PairId = pricePair.ID
		price.Value = value
		price.CreatedAt = current
		if err := price.Create(database.Connection); err != nil {
			return priceList, err
		}

		priceList = append(priceList, price)
		current = current.Add(time.Minute)
	}
	return priceList, nil
}
