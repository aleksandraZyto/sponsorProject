package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username  string      `json:"username" gorm:"unique"`
	Name      string      `json:"name"`
	Password  string      `json:"password"`
	Messages  []Message   `json:"messages" gorm:"foreignKey:Sender"`
	ChatRooms []*ChatRoom `json:"chatRooms" gorm:"many2many:user_chatRooms;"`
}

type Message struct {
	gorm.Model
	Sender   uint
	Text     string    `json:"text"`
	SentAt   time.Time `json:"sentAt"`
	ChatRoom uint
}

type ChatRoom struct {
	gorm.Model
	Participants []*User   `json:"participants" gorm:"many2many:user_chatRooms;"`
	Messages     []Message `json:"messages" gorm:"foreignKey:ChatRoom"`
}
