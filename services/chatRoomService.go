package services

import (
	models "chat-app/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateChatRoom(c *gin.Context) {
	cookieUser, err := GetUserCookie(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not logged in. " + err.Error()})
	}
	// add chat room to db
	chatRoom := models.ChatRoom{
		Participants: nil,
		Messages:     nil,
		Creator:      cookieUser,
	}
	fmt.Println(chatRoom)
}
