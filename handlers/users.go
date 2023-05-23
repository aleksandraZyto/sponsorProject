package handlers

import (
	database "chat-app/database"
	models "chat-app/models"
	"encoding/base64"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"` // these have to be capital, otherwise body parser wont work
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserHandler interface {
	Register(req *RegisterRequest) models.User
	Login(req *LoginRequest) error // why LG has to be a pointer? apparently non pointer breaks body parsing:/
}

type UserHandlerStruct struct{}

// IMPLEMENT ERROR TO RETURN
func (handler *UserHandlerStruct) Register(req *RegisterRequest) models.User {
	req.Password = base64.StdEncoding.EncodeToString([]byte(req.Password))
	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}
	database.DB.Db.Create(&user)
	return user
}

func (handler *UserHandlerStruct) Login(req *LoginRequest) error {
	user := models.User{}
	if err := database.DB.Db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return err
	}

	encodedPass := base64.StdEncoding.EncodeToString([]byte(req.Password))

	if string(user.Password) != encodedPass {
		return errors.New("Invalid password")
	}
	return nil
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
