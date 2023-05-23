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

	app.Post("/login", func(c *fiber.Ctx) error {
		req := new(handlers.LoginRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(400).SendString("Invalid payload")
		}
		handlerErr := handler.Login(req)
		if handlerErr != nil {
			return c.Status(500).SendString(handlerErr.Error())
		}
		return c.Status(200).SendString("Login successful")
	})

	return app
}
