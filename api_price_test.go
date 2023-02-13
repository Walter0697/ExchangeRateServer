package main

import (
	"chaos/backend/config"
	"chaos/backend/database"
	"chaos/backend/seed"
	"chaos/backend/test"
	"io/ioutil"
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"
)

type APIPriceTestSuite struct {
	suite.Suite
	router *chi.Mux
}

func (suite *APIPriceTestSuite) SetupTest() {
	config.Init("config.test.toml")
	database.Init()
	database.Refresh()
	database.AutoMigration()
	seed.SeedDatabase()

	suite.router = getServer()
}

func (suite *APIPriceTestSuite) TestPriceTest() {
	total_number := 20 + rand.Intn(10)
	list, err := test.GenerateRandomPriceList(total_number)
	if err != nil {
		assert.Equal(suite.T(), "", err.Error())
	}

	// getting apikey token
	keyString := `{"identifier": "pricetest"}`
	api_token := test.APIKeyTest(suite.Suite, suite.router, keyString)

	// testing last api
	lastRequest := httptest.NewRequest("GET", "/price/last", nil)
	lastRequest.Header.Add("Authorization", api_token)

	lastRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(lastRecorder, lastRequest)
	assert.Equal(suite.T(), 200, lastRecorder.Result().StatusCode)

	lastresBody, _ := ioutil.ReadAll(lastRecorder.Result().Body)
	lastBody := gjson.ParseBytes(lastresBody)
	value_last := float64(lastBody.Get("value").Float())

	lastPrice := list[total_number-1]
	assert.Equal(suite.T(), lastPrice.Value, value_last)

	// testing selected time api
	middle_index := rand.Intn(total_number/2) + rand.Intn(total_number/10)
	middle_price := list[middle_index]
	increase_second := rand.Intn(60)
	request_time := middle_price.CreatedAt.Add(time.Second * time.Duration(increase_second))
	request_time_str := request_time.Format(time.RFC3339)

	timeRequest := httptest.NewRequest("GET", "/price/bytime/"+request_time_str, nil)
	timeRequest.Header.Add("Authorization", api_token)

	timeRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(timeRecorder, timeRequest)
	assert.Equal(suite.T(), 200, timeRecorder.Result().StatusCode)

	middle_price_next := list[middle_index+1]
	add_ratio := float64(increase_second) / float64(60)

	diff_value := middle_price_next.Value - middle_price.Value
	add_value := diff_value * add_ratio
	between := add_value + middle_price.Value

	timeResBody, _ := ioutil.ReadAll(timeRecorder.Result().Body)
	timeBody := gjson.ParseBytes(timeResBody)
	value_time := float64(timeBody.Get("value").Float())

	assert.Equal(suite.T(), between, value_time)

	// testing range api
	starting_index := rand.Intn(3)
	ending_index := total_number - 5 + rand.Intn(2)

	start := list[starting_index].CreatedAt.Format(time.RFC3339)
	end := list[ending_index].CreatedAt.Format(time.RFC3339)

	rangeRequest := httptest.NewRequest("GET", "/price/range/"+start+"/"+end, nil)
	rangeRequest.Header.Add("Authorization", api_token)

	rangerRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(rangerRecorder, rangeRequest)
	assert.Equal(suite.T(), 200, rangerRecorder.Result().StatusCode)

	resBody, _ := ioutil.ReadAll(rangerRecorder.Result().Body)
	body := gjson.ParseBytes(resBody)
	value_average := float64(body.Get("value").Float())

	var sum float64 = 0
	for i := starting_index; i <= ending_index; i++ {
		sum += list[i].Value
	}
	average := sum / float64(ending_index-starting_index+1)

	assert.Equal(suite.T(), average, value_average)
}

func TestPrice(t *testing.T) {
	suite.Run(t, new(APIPriceTestSuite))
}
