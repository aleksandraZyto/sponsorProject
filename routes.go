package main

import (
	handlers "chat-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func CreateApp(handler handlers.UserHandler) *fiber.App {
	app := fiber.New()

	app.Post("/register", func(c *fiber.Ctx) error {
		req := new(handlers.RegisterRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(400).SendString("Invalid payload")
		}

		user := handler.Register(req)
		return c.Status(200).JSON(user)
	})

	return app
}
