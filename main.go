package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	username string `json:"username"`
	password string `json:"password"`
}

type User struct {
	loginData LoginData `json:"loginData"`
	name      string    `json:"name"`
}

func main() {
	router := gin.Default()
	router.POST("/register", registerHandler)
	router.POST("/login", loginHandler)
	router.GET("/hello", helloHandler)

	router.Run("localhost:8080")
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello")
}

func registerHandler(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	c.String(http.StatusOK, "success")
}

func loginHandler(c *gin.Context) {
	var loginData LoginData
	loginErr := c.ShouldBindJSON(&loginData)
	fmt.Println(loginErr)
	if loginErr != nil {
		return
	}
	c.String(http.StatusOK, "success")
}
