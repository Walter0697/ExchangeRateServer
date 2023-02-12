package main

import (
	"chaos/backend/config"
	"chaos/backend/database"
	"chaos/backend/seed"
	"chaos/backend/test"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type APIKeyApiTestSuite struct {
	suite.Suite
	router *chi.Mux
}

func (suite *APIKeyApiTestSuite) SetupTest() {
	config.Init("config.test.toml")
	database.Init()
	database.Refresh()
	database.AutoMigration()
	seed.SeedDatabase()

	suite.router = getServer()
}

func (suite *APIKeyApiTestSuite) TestLoginAPI() {
	keyString := `{"identifier": "testing"}`

	// generate api key without bearer token
	withoutAuthReq := httptest.NewRequest("POST", "/apikey", strings.NewReader(keyString))
	noAuthRecoder := httptest.NewRecorder()
	suite.router.ServeHTTP(noAuthRecoder, withoutAuthReq)
	assert.Equal(suite.T(), 401, noAuthRecoder.Result().StatusCode)

	test.APIKeyTest(suite.Suite, suite.router, keyString)
}

func TestAPIkey(t *testing.T) {
	suite.Run(t, new(APIKeyApiTestSuite))
}
