package services

import (
	"chat-app/models"
	"chat-app/repos"
	"errors"
)

type Registerer interface {
	WrapperRegister(req *models.RegisterRequest, repo repos.UserRepository) (models.User, error)
}
type Logger interface {
	WrapperLogin(req *models.LoginRequest, repo repos.UserRepository) error
}

type RegistererStruct struct{}
type LoggerStruct struct{}

func (r RegistererStruct) WrapperRegister(req *models.RegisterRequest, repo repos.UserRepository) (models.User, error) {
	return Register(req, repo)
}
func (l LoggerStruct) WrapperLogin(req *models.LoginRequest, repo repos.UserRepository) error {
	return Login(req, repo)
}

func Register(req *models.RegisterRequest, repo repos.UserRepository) (models.User, error) {
	req.LoginData.Password = Encode(req.LoginData.Password)
	return repo.AddUser(req)
}

func Login(req *models.LoginRequest, repo repos.UserRepository) error {
	user, err := repo.GetUser(req.Username)
	if err != nil {
		return err
	}

	encodedPass := Encode(req.Password)
	if string(user.EncodedPassword) != encodedPass {
		return errors.New("invalid password")
	}
	return nil
}
