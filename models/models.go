package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	//UserId   string    `json:"userId" gorm:"unique"`
	Username string    `json:"username" gorm:"unique"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Messages []Message `json:"messages" gorm:"foreignKey:Sender"`
}

type Message struct {
	gorm.Model
	Sender uint
	Text   string    `json:"text"`
	SentAt time.Time `json:"sentAt"`
}

type ChatRoom struct {
	gorm.Model
	//Participantqs []User    `json:"participants" gorm:"many2many:chatRoom_participants;"`
	//Messages     []Message `json:"messages" gorm:"many2many:chatRoom_messages;"`
}
