package handlers

import (
	database "chat-app/database"
	models "chat-app/models"
	"encoding/base64"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	name     string `json: "name"`
	password string `json: "password"`
	username string `json: "username"`
}

type UserHandler interface {
	Register(req *RegisterRequest) models.User
}

type UserHandlerStruct struct{}

func (handler *UserHandlerStruct) Register(req *RegisterRequest) models.User {
	fmt.Println(req)
	return models.User{Name: "Mock", Password: "mock", Username: "mock"}
}

func RegistrationHandler(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user.Password = base64.StdEncoding.EncodeToString([]byte(user.Password))
	database.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}

func Home(c *fiber.Ctx) error {
	json := models.User{
		Name:     "haha",
		Password: "sdfsd",
		Username: "ola",
	}
	return c.Status(200).JSON(json)
}

func LoginHandler(c *fiber.Ctx) error {
	var user models.User

	payload := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error while parsing the request body: " + err.Error())
	}

	if err := database.DB.Db.Where("username = ?", payload.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Username not found")
	}

	encodedPass := base64.StdEncoding.EncodeToString([]byte(payload.Password))

	if string(user.Password) != encodedPass {
		return c.Status(401).SendString("Password incorrect")
	}

	return c.Status(200).JSON(user)
}
