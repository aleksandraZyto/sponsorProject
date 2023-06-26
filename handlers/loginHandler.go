package handlers

import (
	"chat-app/models"
	"chat-app/repos"
	"chat-app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandler(c *gin.Context) {
	req := new(models.RegisterRequest)
	repo := repos.UserRepositoryStruct{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	user, err := services.Register(req, repo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"Created user:": user})
	}
}

func LoginHandler(c *gin.Context) {
	req := new(models.LoginRequest)
	repo := repos.UserRepositoryStruct{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.Login(req, repo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	SetUserCookie(c, req.Username)
	c.JSON(http.StatusOK, gin.H{})
}
