package handlers

import (
	"chat-app/models"
	"chat-app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandler(c *gin.Context) {
	userService := &services.UserHandlerStruct{}
	req := new(models.RegisterRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	user, err := userService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"Created user:": user})
	}
}

func LoginHandler(c *gin.Context) {
	userService := &services.UserHandlerStruct{}
	req := new(models.LoginRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userService.Login(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	SetUserCookie(c, req.Username)
	c.JSON(http.StatusOK, gin.H{})
}
