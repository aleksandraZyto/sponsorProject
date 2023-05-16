package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	username string `json:"username"`
	password string `json:"password"`
	name     string `json:"name"`
}

func main() {
	router := gin.Default()
	router.POST("/register", registerHandler)
	router.POST("/login", loginHandler)

	router.Run("localhost:8080")
}

func registerHandler(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	c.String(http.StatusOK)
}

func loginHandler(c *gin.Context) {
	var username string
	var password string
	usernameErr := c.BindJSON(&username)
	passwordErr := c.BindJSON(&password)
	if usernameErr != nil || passwordErr != nil {
		return
	}
	c.String(http.StatusOK)
}
