package main

import (
	"chaos/backend/config"
	"chaos/backend/database"
	"chaos/backend/seed"
	"chaos/backend/test"
	"io/ioutil"
	"math/rand"
	"net/http/httptest"
	"strings"
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
	list, err := test.GenerateRandomPriceList(10)
	if err != nil {
		assert.Equal(suite.T(), "", err.Error())
	}

	// getting apikey token
	keyString := `{"identifier": "pricetest"}`
	api_token := test.APIKeyTest(suite.Suite, suite.router, keyString)

	// testing range api
	starting_index := rand.Intn(3)
	ending_index := 7 + rand.Intn(2)

	start := list[starting_index].CreatedAt.Format(time.RFC3339)
	end := list[ending_index].CreatedAt.Format(time.RFC3339)

	rangeRequest := httptest.NewRequest("GET", "/price/range/"+start+"/"+end, strings.NewReader(""))
	rangeRequest.Header.Add("Authorization", api_token)

	rangerRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(rangerRecorder, rangeRequest)
	assert.Equal(suite.T(), 200, rangerRecorder.Result().StatusCode)

	resBody, _ := ioutil.ReadAll(rangerRecorder.Result().Body)
	body := gjson.ParseBytes(resBody)
	usd_average := float64(body.Get("usd").Float())

	var sum float64 = 0
	for i := starting_index; i < ending_index; i++ {
		sum += list[i].Value
	}
	average := sum / float64(ending_index-starting_index)

	assert.Equal(suite.T(), average, usd_average)
}

func TestPrice(t *testing.T) {
	suite.Run(t, new(APIPriceTestSuite))
}
