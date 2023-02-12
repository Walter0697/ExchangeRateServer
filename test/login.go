package test

import (
	"chaos/backend/config"
	"io/ioutil"
	"net/http/httptest"
	"strings"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"
)

func LoginTest(suite suite.Suite, router *chi.Mux) string {
	loginString := `{"username":"` + config.Data.Default.AdminUsername + `", "password": "` + config.Data.Default.AdminPassword + `"}`
	willSuccessRequest := httptest.NewRequest("POST", "/auth/login", strings.NewReader(loginString))
	successRecoder := httptest.NewRecorder()
	router.ServeHTTP(successRecoder, willSuccessRequest)
	assert.Equal(suite.T(), 200, successRecoder.Result().StatusCode)

	resBody, _ := ioutil.ReadAll(successRecoder.Result().Body)
	body := gjson.ParseBytes(resBody)

	token := body.Get("token").String()
	return token
}
