package services

import (
	"chat-app/models"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type Repo struct {
	numberOfCallsAddUser int
	numberOfCallsGetUser int
	calledWithAddUser    *models.RegisterRequest
	calledWithGetUser    string
}

func (repo *Repo) AddUser(req *models.RegisterRequest) (models.User, error) {
	repo.numberOfCallsAddUser++
	repo.calledWithAddUser = req
	user := models.User{
		Name:            "Aleksandra",
		Username:        "ola",
		EncodedPassword: "123",
	}
	return user, nil
}

func (repo *Repo) GetUser(username string) (models.User, error) {
	repo.numberOfCallsGetUser++
	repo.calledWithGetUser = username
	user := models.User{
		Name:            "Aleksandra",
		Username:        "ola",
		EncodedPassword: "NDU2",
	}
	return user, nil
}

func TestUserHandlerStruct_RegisterHappyPath(t *testing.T) {
	repository := &Repo{}
	user := UserHandlerStruct{}
	req := models.RegisterRequest{}
	expectedNewUser := models.User{
		Name:            "Aleksandra",
		Username:        "ola",
		EncodedPassword: "123",
	}

	actualNewUser, err := user.Register(&req, repository)

	if err != nil {
		t.Errorf("Error is not nil: %v", err.Error())
	}
	if !reflect.DeepEqual(expectedNewUser, actualNewUser) {
		t.Errorf("Expected and actual user are not the same")
	}
	if repository.calledWithAddUser != &req {
		t.Errorf("Register method expected to have been with %v", req)
	}
	if repository.numberOfCallsAddUser != 1 {
		t.Errorf("Register method expected to have been called once, but it was called %d times", repository.numberOfCallsAddUser)
	}
}

func TestUserHandlerStruct_LoginHappyPath(t *testing.T) {
	repository := &Repo{}
	user := UserHandlerStruct{}
	req := models.LoginRequest{
		Username: "olka",
		Password: "456",
	}

	err := user.Login(&req, repository)

	if err != nil {
		t.Errorf("Error is not nil: %v", err.Error())
	}
	if repository.calledWithGetUser != req.Username {
		t.Errorf("Login method expected to have been with %v", req.Username)
	}
	if repository.numberOfCallsGetUser != 1 {
		t.Errorf("Login method expected to have been called once, but it was called %d times", repository.numberOfCallsAddUser)
	}
}

func TestUserHandlerStruct_LoginInvalidPasswordPath(t *testing.T) {
	repository := &Repo{}
	user := UserHandlerStruct{}
	req := models.LoginRequest{
		Username: "olka",
		Password: "123",
	}

	err := user.Login(&req, repository)
	fmt.Println("Err: ", err)

	if err != nil {
		assert.Equal(t, err.Error(), "invalid password", "The error messages should be the same.")
	} else {
		t.Errorf("Error should indicate wrong password, but instead got no error")
	}
}
