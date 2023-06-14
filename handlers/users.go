package handlers

import (
	database "chat-app/database"
	models "chat-app/models"
	"encoding/base64"
	"errors"
)

type RegisterRequest struct {
	LoginData LoginRequest `json:loginData`
	Name      string       `json:"name"`
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
	req.LoginData.Password = base64.StdEncoding.EncodeToString([]byte(req.LoginData.Password))
	user := models.User{
		Name:     req.Name,
		Username: req.LoginData.Username,
		Password: req.LoginData.Password,
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
