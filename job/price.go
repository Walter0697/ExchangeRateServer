package job

import (
	"chaos/backend/database"
	"chaos/backend/database/model"
	"chaos/backend/service"
	"fmt"
)

func FetchBTCUSDPair() {
	// var apiEnableSetting model.Setting
	// apiEnableSetting.Identifier = "fetch_enable"
	// if err := apiEnableSetting.GetByIdentifier(database.Connection); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// if !apiEnableSetting.BoolValue {
	// 	// if api enable is not true, don't fetch anything
	// 	return
	// }

	crypto := "BTC"
	currency := "USD"
	resp, err := service.GetCurrentPrice(crypto, currency)
	if err != nil {
		fmt.Println(err)
		return
	}

	var pricePair model.PricePair
	pricePair.Crypto = crypto
	pricePair.Currency = currency
	if err = pricePair.GetOrCreate(database.Connection); err != nil {
		fmt.Println(err)
		return
	}

	var price model.Price
	price.PairId = pricePair.ID
	price.Value = resp.USD
	if err = price.Create(database.Connection); err != nil {
		fmt.Println(err)
		return
	}
}
