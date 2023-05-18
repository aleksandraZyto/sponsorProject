package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name"`
	LoginDataID int
	LoginData   LoginData `json:"login_data"`
}

type LoginData struct {
	ID       int
	Username string `json:"username"`
	Password string `json:"password"`
}
