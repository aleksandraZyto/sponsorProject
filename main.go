package main

import (
	"fmt"
	"net/http"

	models "sponsorProject/models"

	"github.com/gin-gonic/gin"
)

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
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Something went wrong")
		return
	}
	c.String(http.StatusOK, "success")
}

func loginHandler(c *gin.Context) {
	var loginData models.LoginData
	loginErr := c.ShouldBindJSON(&loginData)
	fmt.Println(loginErr)
	if loginErr != nil {
		c.String(http.StatusBadRequest, "Something went wrong")
		return
	}
	c.String(http.StatusOK, "success")
}
