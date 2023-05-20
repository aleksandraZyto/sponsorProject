package main

import (
	handlers "chat-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/hello", handlers.Home)
	app.Post("/register", handlers.RegistrationHandler)
	app.Post("/login", handlers.LoginHandler)
}
