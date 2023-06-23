package services

import (
	"chat-app/models"
	"chat-app/repos"
	"errors"
)

type UserHandler interface {
	Register(req *models.RegisterRequest) models.User
	Login(req *models.LoginRequest) error // why LR has to be a pointer? apparently non pointer breaks body parsing:/
}

type UserHandlerStruct struct{}

func (handler *UserHandlerStruct) Register(req *models.RegisterRequest) (models.User, error) {
	req.LoginData.Password = Encode(req.LoginData.Password)

	return repos.AddUser(req)
}

func (handler *UserHandlerStruct) Login(req *models.LoginRequest) error {
	user, err := repos.GetUser(req.Username)
	if err != nil {
		return err
	}

	encodedPass := Encode(req.Password)
	if string(user.Password) != encodedPass {
		return errors.New("invalid password")
	}
	return nil
}
