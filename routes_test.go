package main

import (
	"bytes"
	"chat-app/database"
	"chat-app/handlers"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"chat-app/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()
	teardown()

	os.Exit(exitCode)
}

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

func setup() {
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Fatal("Error loading .env.test file")
	}

	database.ConnectDb()
	database.DB.Db.AutoMigrate(&models.User{})
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}

func router() *gin.Engine {
	router := gin.Default()

	router.POST("/register", RegisterHandler)
	router.POST("/login", LoginHandler)

	return router
}

func teardown() {
	migrator := database.DB.Db.Migrator()
	migrator.DropTable(&models.User{})
}
