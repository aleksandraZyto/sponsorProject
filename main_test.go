package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHelloHandler(t *testing.T) {
	r := SetUpRouter()
	mockResponse := "\"Hello\""
	r.GET("/hello", helloHandler)

	req, _ := http.NewRequest("GET", "/hello", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	responseData, _ := ioutil.ReadAll(recorder.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestLoginHandler(t *testing.T) {
	loginData := LoginData{Username: "ola", Password: "123"}
	jsonValue, _ := json.Marshal(loginData)

	r := SetUpRouter()
	r.POST("/login", loginHandler)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestRegisterHandler(t *testing.T) {
	user := User{
		LoginData: LoginData{Username: "olazyto", Password: "123"},
		Name:      "ola",
	}
	jsonValue, _ := json.Marshal(user)

	r := SetUpRouter()
	r.POST("/register", registerHandler)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
