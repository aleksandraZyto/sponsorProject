package services

import (
	"chat-app/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	userHandler := &handlers.UserHandlerStruct{}
	req := new(handlers.RegisterRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	user, err := userHandler.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"Created user:": user})
	}
}

func LoginHandler(c *gin.Context) {
	handler := &handlers.UserHandlerStruct{}
	req := new(handlers.LoginRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	handlerErr := handler.Login(req)
	if handlerErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": handlerErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
