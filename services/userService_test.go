package services

import (
	"chat-app/models"
	"reflect"
	"testing"
)

type Repo struct {
	numberOfCalls int
	calledWith    *models.RegisterRequest
}

func (repo *Repo) AddUser(req *models.RegisterRequest) (models.User, error) {
	repo.numberOfCalls++
	repo.calledWith = req
	user := models.User{
		Name:     "Aleksandra",
		Username: "ola",
		Password: "123",
	}
	return user, nil
}

func (repo *Repo) GetUser(_ string) (models.User, error) {
	repo.numberOfCalls++
	user := models.User{
		Name:     "",
		Username: "",
		Password: "",
	}
	return user, nil
}

func TestUserHandlerStruct_RegisterHappyPath(t *testing.T) {
	repository := &Repo{}
	user := UserHandlerStruct{}
	req := models.RegisterRequest{}
	expectedNewUser := models.User{
		Name:     "Aleksandra",
		Username: "ola",
		Password: "123",
	}

	actualNewUser, err := user.Register(&req, repository)

	if err != nil {
		t.Errorf("Error is not nil")
	}
	if !reflect.DeepEqual(expectedNewUser, actualNewUser) {
		t.Errorf("Expected and actual user are not the same")
	}
	if repository.calledWith != &req {
		t.Errorf("Register method expected to have been with %v", req)
	}
	if repository.numberOfCalls != 1 {
		t.Errorf("Register method expected to have been called once, but it was called %d times", repository.numberOfCalls)
	}
}
