package main

import (
	database "chat-app/database"
	"chat-app/handlers"
)

func main() {
	database.ConnectDb()
	// app := fiber.New()
	// SetupRoutes(app)
	// app.Listen(":3000")

	alterApp := CreateApp(&handlers.UserHandlerStruct{})
	alterApp.Listen(":3000")
}
