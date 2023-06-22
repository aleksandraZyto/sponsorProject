package handlers

import (
	database "chat-app/database"
	models "chat-app/models"
	"encoding/base64"
	"errors"
)

type RegisterRequest struct {
	LoginData LoginRequest `json:"loginData" binding:"required"`
	Name      string       `json:"name" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserHandler interface {
	Register(req *RegisterRequest) models.User
	Login(req *LoginRequest) error // why LR has to be a pointer? apparently non pointer breaks body parsing:/
}

type UserHandlerStruct struct{}

// IMPLEMENT ERROR TO RETURN
func (handler *UserHandlerStruct) Register(req *RegisterRequest) (models.User, error) {
	req.LoginData.Password = base64.StdEncoding.EncodeToString([]byte(req.LoginData.Password))
	user := models.User{
		Name:     req.Name,
		Username: req.LoginData.Username,
		Password: req.LoginData.Password,
	}
	gormErr := database.DB.Db.Create(&user)
	if gormErr.Error != nil {
		err := errors.New("error when creating a user, username might be already taken")
		return user, err
	}
	return user, nil
}

func (handler *UserHandlerStruct) Login(req *LoginRequest) error {
	user := models.User{}
	if err := database.DB.Db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return err
	}

	encodedPass := base64.StdEncoding.EncodeToString([]byte(req.Password))

	if string(user.Password) != encodedPass {
		return errors.New("invalid password")
	}
	return nil
}
