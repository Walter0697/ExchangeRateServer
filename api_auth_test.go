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

type AuthAPITestSuite struct {
	suite.Suite
	router *chi.Mux
}

func (suite *AuthAPITestSuite) SetupTest() {
	config.Init("config.test.toml")
	database.Init()
	database.Refresh()
	database.AutoMigration()
	seed.SeedDatabase()

	suite.router = getServer()
}

func (suite *AuthAPITestSuite) TestLoginAPI() {
	// Fail login
	willFailRequest := httptest.NewRequest("POST", "/auth/login", strings.NewReader(`
		{
			"username": "random",
			"password": "veryrandompass"
		}
	`))
	failRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(failRecorder, willFailRequest)
	assert.Equal(suite.T(), 401, failRecorder.Result().StatusCode)

	// Success login
	test.LoginTest(suite.Suite, suite.router)
}

func TestAuthAPI(t *testing.T) {
	suite.Run(t, new(AuthAPITestSuite))
}
