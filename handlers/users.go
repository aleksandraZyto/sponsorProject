package handlers

import (
	database "chat-app/database"
	models "chat-app/models"
	"encoding/base64"
	"errors"
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