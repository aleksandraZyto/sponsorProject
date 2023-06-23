package services

import (
	"chat-app/models"
	"testing"
)

type Repo struct {
	numberOfCalls int
}

func (repo *Repo) AddUser(req *models.RegisterRequest) (models.User, error) {
	repo.numberOfCalls++
	user := models.User{
		Name:     "",
		Username: "",
		Password: "",
	}
	return user, nil
}

func (repo *Repo) GetUser(username string) (models.User, error) {
	repo.numberOfCalls++
	user := models.User{
		Name:     "",
		Username: "",
		Password: "",
	}
	return user, nil
}

func TestUserHandlerStruct_Register(t *testing.T) {
	user := UserHandlerStruct{}
	req := &models.RegisterRequest{}
	repository := &Repo{}

	user.Register(req, repository)

	if repository.numberOfCalls == 0 {
		t.Errorf("Expected dog to bark once, but it barked %d times", repository.numberOfCalls)
	}
}
