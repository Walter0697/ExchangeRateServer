package controller

import (
	"chaos/backend/database"
	"chaos/backend/database/model"
	"chaos/backend/service"
	"chaos/backend/utility"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type MissingDateError struct{}

func (m *MissingDateError) Error() string {
	return "missing requested date"
}

type PriceResultRespond struct {
	BaseRespond
	USD float64 `json:"usd"`
}

// helper function for getting price pair
// default to be BTC to USD
func GetPricePairByRequest(r *http.Request) (*model.PricePair, error) {
	cryptoStr := "BTC"
	currencyStr := "USD"
	crypto := r.URL.Query().Get("crypto")
	currency := r.URL.Query().Get("currency")

	if crypto != "" {
		cryptoStr = crypto
	}
	if currency != "" {
		currencyStr = currency
	}

	var pricePair model.PricePair
	pricePair.Crypto = cryptoStr
	pricePair.Currency = currencyStr
	if err := pricePair.GetOrCreate(database.Connection); err != nil {
		return nil, err
	}

	return &pricePair, nil
}

// get latest price
// Authorization: apikey
func GetPriceByLatest(w http.ResponseWriter, r *http.Request) {
	pricePair, err := GetPricePairByRequest(r)
	if err != nil {
		ErrorResp(w, r, http.StatusBadRequest, err)
		return
	}

	var price model.Price
	price.PairId = pricePair.ID
	if err := price.FindByLatest(database.Connection); err != nil {
		ErrorResp(w, r, http.StatusInternalServerError, err)
		return
	}

	result := PriceResultRespond{
		USD: price.Value,
	}

	JSON(w, r, 200, &result)
}

// get price by time
// Authorization: apikey
func GetPriceByTime(w http.ResponseWriter, r *http.Request) {
	datetime := chi.URLParam(r, "time")
	if datetime == "" {
		ErrorResp(w, r, http.StatusBadRequest, &MissingDateError{})
		return
	}

	pricePair, err := GetPricePairByRequest(r)
	if err != nil {
		ErrorResp(w, r, http.StatusBadRequest, err)
		return
	}

	requestedTime, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		ErrorResp(w, r, http.StatusBadRequest, err)
		return
	}

	var previousPrice model.Price
	previousPrice.PairId = pricePair.ID
	previousPrice.CreatedAt = requestedTime

	var nextPrice model.Price
	nextPrice.PairId = pricePair.ID
	nextPrice.CreatedAt = requestedTime

	if err := previousPrice.FindPreviousByTime(database.Connection); err != nil {
		if utility.RecordNotFound(err) {
			ErrorResp(w, r, http.StatusNotFound, err)
			return
		}
		ErrorResp(w, r, http.StatusInternalServerError, err)
		return
	}

	if err := nextPrice.FindNextByTime(database.Connection); err != nil {
		if utility.RecordNotFound(err) {
			ErrorResp(w, r, http.StatusNotFound, err)
			return
		}
		ErrorResp(w, r, http.StatusInternalServerError, err)
		return
	}

	diff_mid := requestedTime.Sub(previousPrice.CreatedAt)
	diff_whole := nextPrice.CreatedAt.Sub(previousPrice.CreatedAt)

	ratio := diff_mid.Seconds() / diff_whole.Seconds()

	diff_value := nextPrice.Value - previousPrice.Value
	add_value := diff_value * ratio
	between := add_value + previousPrice.Value

	result := PriceResultRespond{
		USD: between,
	}

	JSON(w, r, 200, &result)
}

// get average price by range
// Authorization: apikey
func GetAverageByRange(w http.ResponseWriter, r *http.Request) {
	start := chi.URLParam(r, "start")
	if start == "" {
		ErrorResp(w, r, http.StatusBadRequest, &MissingDateError{})
		return
	}

	end := chi.URLParam(r, "end")
	if end == "" {
		ErrorResp(w, r, http.StatusBadRequest, &MissingDateError{})
		return
	}

	startTime, err := time.Parse(time.RFC3339, start)
	if err != nil {
		ErrorResp(w, r, http.StatusBadRequest, err)
		return
	}

	endTime, err := time.Parse(time.RFC3339, end)
	if err != nil {
		ErrorResp(w, r, http.StatusBadRequest, err)
		return
	}

	pricePair, err := GetPricePairByRequest(r)
	if err != nil {
		ErrorResp(w, r, http.StatusBadRequest, err)
		return
	}

	priceList, err := service.GetPriceList(database.Connection, *pricePair, startTime, endTime)
	if err != nil {
		ErrorResp(w, r, http.StatusInternalServerError, err)
		return
	}

	var sum float64 = 0
	for _, price := range priceList {
		sum += price.Value
	}

	average := sum / float64(len(priceList))

	result := PriceResultRespond{
		USD: float64(average),
	}

	JSON(w, r, 200, &result)
}
