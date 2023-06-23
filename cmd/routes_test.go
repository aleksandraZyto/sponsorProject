package main

import (
	"chat-app/services"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterHappyPath(t *testing.T) {
	req := services.RegisterRequest{
		LoginData: services.LoginRequest{
			Username: "ale",
			Password: "1234",
		},
		Name: "aleksandra",
	}
	writer := makeRequest("POST", "/register", req)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestRegisterBadRequest(t *testing.T) {
	req := services.RegisterRequest{
		LoginData: services.LoginRequest{
			Username: "ale",
		},
		Name: "aleksandra",
	}
	writer := makeRequest("POST", "/register", req)
	assert.Equal(t, http.StatusBadRequest, writer.Code)
}

// run without teardown
func TestLoginHappyPath(t *testing.T) {
	req := services.LoginRequest{
		Username: "ale",
		Password: "1234",
	}

	writer := makeRequest("POST", "/login", req)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestLoginBadRequest(t *testing.T) {
	req := services.LoginRequest{
		Password: "1234",
	}

	writer := makeRequest("POST", "/login", req)
	assert.Equal(t, http.StatusBadRequest, writer.Code)
}
