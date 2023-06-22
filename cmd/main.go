package main

import (
	database "chat-app/database"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("App started")
	err := database.ConnectDb()
	if err != nil {
		time.Sleep(10 * time.Second)
		os.Exit(1)
	}
	serveApplication()
}

func serveApplication() {
	router := gin.Default()
	router.POST("/register", RegisterHandler)
	router.POST("/login", LoginHandler)

	router.Run(":3000")
	log.Println("Server running on port 3000")
}
