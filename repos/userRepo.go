package repos

import (
	"chat-app/database"
	"chat-app/models"
	"errors"
)

func GetUser(username string) (models.User, error) {
	user := models.User{}
	if err := database.DB.Db.Where(
		"username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func AddUser(req *models.RegisterRequest) (models.User, error) {
	user := models.User{
		Name:     req.Name,
		Username: req.LoginData.Username,
		Password: req.LoginData.Password,
	}

	if gormErr := database.DB.Db.Create(&user); gormErr.Error != nil {
		err := errors.New("error when creating a user, username might be already taken")
		return user, err
	}

	return user, nil
}
