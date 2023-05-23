package main

import (
	database "chat-app/database"
	"chat-app/handlers"
)

func main() {
	database.ConnectDb()

	app := CreateApp(&handlers.UserHandlerStruct{})
	app.Listen(":3000")
}
