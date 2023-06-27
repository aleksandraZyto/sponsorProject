package handlers

import (
	"bytes"
	"chat-app/models"
	"chat-app/repos"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RegistererStructMock struct{}
type LoggerStructMock struct{}

var RegisterMock func(req *models.RegisterRequest, repo repos.UserRepository) (models.User, error)
var LoginMock func(req *models.LoginRequest, repo repos.UserRepository) error

func (rsm RegistererStructMock) WrapperRegister(req *models.RegisterRequest, repo repos.UserRepository) (models.User, error) {
	return RegisterMock(req, repo)
}
func (lsm LoggerStructMock) WrapperLogin(req *models.LoginRequest, repo repos.UserRepository) error {
	return LoginMock(req, repo)
}

func TestRegisterHandler_HappyPath(t *testing.T) {
	expectedNewUser := models.User{
		Username:        "Olka",
		EncodedPassword: "1111",
		Name:            "Alex",
	}
	RegisterMock = func(req *models.RegisterRequest, repo repos.UserRepository) (models.User, error) {
		return expectedNewUser, nil
	}
	registerReq := &models.RegisterRequest{
		LoginData: models.LoginRequest{
			Username: "Olka",
			Password: "1111",
		},
		Name: "Alex",
	}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	jsonValue, _ := json.Marshal(registerReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	c.Request = req

	RegisterHandler(c, &RegistererStructMock{})

	actualNewUser := models.User{}
	json.Unmarshal([]byte(recorder.Body.String()), &actualNewUser)

	assert.Equal(t, expectedNewUser, actualNewUser)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestRegisterHandler_TakenUsername(t *testing.T) {
	RegisterMock = func(req *models.RegisterRequest, repo repos.UserRepository) (models.User, error) {
		return models.User{}, errors.New("error when creating a user, username might be already taken")
	}
	registerReq := &models.RegisterRequest{
		LoginData: models.LoginRequest{
			Username: "Olka",
			Password: "1111",
		},
		Name: "Alex",
	}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	jsonValue, _ := json.Marshal(registerReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	c.Request = req

	RegisterHandler(c, &RegistererStructMock{})

	actualErr := recorder.Body.String()
	expectedErr := "\"error when creating a user, username might be already taken\""

	assert.Equal(t, expectedErr, actualErr)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestRegisterHandler_BadRequest(t *testing.T) {
	RegisterMock = func(req *models.RegisterRequest, repo repos.UserRepository) (models.User, error) {
		return models.User{}, nil
	}
	registerReq := &models.LoginRequest{
		Username: "Olka",
		Password: "1111",
	}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	jsonValue, _ := json.Marshal(registerReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	c.Request = req

	RegisterHandler(c, &RegistererStructMock{})

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestLoginHandler_HappyPath(t *testing.T) {
	LoginMock = func(req *models.LoginRequest, repo repos.UserRepository) error {
		return nil
	}
	loginReq := &models.LoginRequest{
		Username: "Olka",
		Password: "1111",
	}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	jsonValue, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	c.Request = req

	LoginHandler(c, &LoggerStructMock{})

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestLoginHandler_InvalidPassPath(t *testing.T) {
	LoginMock = func(req *models.LoginRequest, repo repos.UserRepository) error {
		return errors.New("invalid password")
	}
	loginReq := &models.LoginRequest{
		Username: "Olka",
		Password: "x",
	}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	jsonValue, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	c.Request = req

	LoginHandler(c, &LoggerStructMock{})

	actualErr := recorder.Body.String()
	expectedErr := "\"invalid password\""

	assert.Equal(t, expectedErr, actualErr)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestLoginHandler_BadRequest(t *testing.T) {
	LoginMock = func(req *models.LoginRequest, repo repos.UserRepository) error {
		return nil
	}
	loginReq := &models.LoginRequest{
		Username: "Olka",
	}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	jsonValue, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	c.Request = req

	LoginHandler(c, &LoggerStructMock{})

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
