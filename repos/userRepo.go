package repos

import (
	"chat-app/database"
	"chat-app/models"
	"errors"
)

type UserRepository interface {
	GetUser(username string) (models.User, error) // TODO: Name "encodedPassword"
	AddUser(req *models.RegisterRequest) (models.User, error)
}

type UserRepositoryStruct struct{}

func (repo UserRepositoryStruct) GetUser(username string) (models.User, error) {
	user := models.User{}
	if err := database.DB.Db.Where(
		"username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo UserRepositoryStruct) AddUser(req *models.RegisterRequest) (models.User, error) {
	user := models.User{
		Name:            req.Name,
		Username:        req.LoginData.Username,
		EncodedPassword: req.LoginData.Password,
	}

	if gormErr := database.DB.Db.Create(&user); gormErr.Error != nil {
		err := errors.New("error when creating a user, username might be already taken")
		return user, err
	}

	return user, nil
}
