package main

import (
	handler "chat-app/handlers"
	service "chat-app/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func serveApplication() {
	router := gin.Default()
	router.POST("/register", service.RegisterRequestReceiver)
	router.POST("/login", service.LoginRequestReceiver)
	router.POST("/createChatRoom", handler.CreateChatRoomReceiver)

	router.Run(":3000")
	log.Println("Server running on port 3000")
}
