package test

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"
)

func APIKeyTest(suite suite.Suite, router *chi.Mux, keyString string) string {
	token := LoginTest(suite, router)

	withTokenInHeaderReq := httptest.NewRequest("POST", "/apikey", strings.NewReader(keyString))
	withTokenInHeaderReq.Header.Add("Authorization", "Bearer "+token)
	withAuthRecorder := httptest.NewRecorder()
	router.ServeHTTP(withAuthRecorder, withTokenInHeaderReq)
	assert.Equal(suite.T(), 200, withAuthRecorder.Result().StatusCode)

	resBody, _ := ioutil.ReadAll(withAuthRecorder.Result().Body)
	body := gjson.ParseBytes(resBody)
	api_token := body.Get("token").String()

	return api_token
}
