package models

import "gorm.io/gorm"

type LoginData struct {
	gorm.Model
	Username string `json:"username" gorm:"text;not null;default:null`
	Password string `json:"password" gorm:"text;not null;default:null`
}

type User struct {
	gorm.Model
	LoginData LoginData `json:"loginData" gorm:"text;not null;default:null`
	Name      string    `json:"name" gorm:"text;not null;default:null`
}
