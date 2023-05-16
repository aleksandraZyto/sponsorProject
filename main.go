package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	LoginData LoginData `json:"loginData"`
	Name      string    `json:"name"`
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
		c.String(http.StatusBadRequest, "Something went wrong")
		return
	}
	c.String(http.StatusOK, "success")
}

func loginHandler(c *gin.Context) {
	var loginData LoginData
	loginErr := c.ShouldBindJSON(&loginData)
	fmt.Println(loginErr)
	if loginErr != nil {
		c.String(http.StatusBadRequest, "Something went wrong")
		return
	}
	c.String(http.StatusOK, "success")
}
