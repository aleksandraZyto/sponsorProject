package main

import (
	service "chat-app/services"
	"github.com/gin-gonic/gin"
	"log"
)

func serveApplication() {
	router := gin.Default()
	router.POST("/register", service.RegisterHandler)
	router.POST("/login", service.LoginHandler)
	router.POST("/createChatRoom", service.CreateChatRoom)

	router.Run(":3000")
	log.Println("Server running on port 3000")
}
