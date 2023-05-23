package main

import (
	handlers "chat-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func CreateApp(handler handlers.UserHandler) *fiber.App {
	app := fiber.New()

	app.Post("/register", func(c *fiber.Ctx) error {
		req := new(handlers.RegisterRequest)
		handler.Register(req)
		return c.Status(200).SendString("Done")
	})

	return app
}

func SetupRoutes(app *fiber.App) {
	app.Get("/hello", handlers.Home)
	app.Post("/register", handlers.RegistrationHandler)
	app.Post("/login", handlers.LoginHandler)
}
