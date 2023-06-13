package main

import (
	handlers "chat-app/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TODO: return correct status codes
func RegisterHandler(c *gin.Context) {
	handler := &handlers.UserHandlerStruct{}
	req := new(handlers.RegisterRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := handler.Register(req)
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

//TODO: return correct status codes
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
