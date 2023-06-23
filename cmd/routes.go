package main

import (
	handler "chat-app/handlers"
	service "chat-app/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func serveApplication() {
	router := gin.Default()
	router.POST("/register", service.RegisterHandler)
	router.POST("/login", service.LoginHandler)
	router.POST("/createChatRoom", handler.CreateChatRoom)

	router.Run(":3000")
	log.Println("Server running on port 3000")
}
