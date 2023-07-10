package main

import (
	"chat-app/database"
	"log"
	"os"
	"time"
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
