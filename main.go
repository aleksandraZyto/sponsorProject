package main

import (
	database "chat-app/database"
	"chat-app/handlers"
	"log"
	"os"
	"time"
)

func main() {
	log.Print("App started")
	err := database.ConnectDb()
	if err != nil {
		time.Sleep(10 * time.Second)
		os.Exit(1)
	}

	app := CreateApp(&handlers.UserHandlerStruct{})
	app.Listen(":3000")
}
