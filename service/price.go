package service

import (
	"chaos/backend/database/model"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type BTCUSDResponse struct {
	USD float64 `json:"USD"`
}

const (
	BTCUSDLink string = "https://min-api.cryptocompare.com/data/price"
)

func GetPairLink(crypto, currency string) string {
	return BTCUSDLink + "?fsym=" + crypto + "&tsyms=" + currency
}

// getting current btc usd price by public api
func GetCurrentPrice(crypto, currency string) (*BTCUSDResponse, error) {
	link := GetPairLink(crypto, currency)

	body, err := GetRequest(link)
	if err != nil {
		return nil, err
	}

	var resp BTCUSDResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func GetPriceList(db *gorm.DB, pp model.PricePair, start time.Time, end time.Time) ([]model.Price, error) {
	query := db.Where("pair_id = ?", pp.ID)
	query = query.Where("created_at >= ?", start.Format(time.RFC3339))
	query = query.Where("created_at <= ?", end.Format(time.RFC3339))

	var priceList []model.Price
	if err := query.Find(&priceList).Error; err != nil {
		return priceList, err
	}
	return priceList, nil
}
