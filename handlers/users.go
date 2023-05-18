package handlers

import (
	database "chat-app/database"
	models "chat-app/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Its working")
}

func RegistrationHandler(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fmt.Println("This is the user:", user)
	database.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}

func LoginHandler(c *fiber.Ctx) error {
	var user models.User

	payload := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := database.DB.Db.Where("username = ?", payload.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Username not found")
	}

	if payload.Password != user.Password {
		return c.Status(401).SendString("Password incorrect")
	}

	fmt.Println("ussseeer:", user)

	return c.Status(200).JSON(user)
}
