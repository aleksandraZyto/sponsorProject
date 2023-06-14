package main

import (
	"chat-app/handlers"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	req := handlers.RegisterRequest{
		LoginData: handlers.LoginRequest{
			Username: "ale",
			Password: "1234",
		},
		Name: "aleksandra",
	}
	writer := makeRequest("POST", "/register", req)
	assert.Equal(t, http.StatusCreated, writer.Code)
}
