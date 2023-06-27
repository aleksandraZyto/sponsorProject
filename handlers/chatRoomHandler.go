package handlers

import (
	"chat-app/models"
	"chat-app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateChatRoomReceiver(c *gin.Context) {
	cookieUser, err := GetUserCookie(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not logged in. " + err.Error()})
		return
	}
	CreateChatRoomHandler(c, services.ChatRoomCreatorStruct{CookieUser: cookieUser})
}

func CreateChatRoomHandler(c *gin.Context, chatRoomCreator services.ChatRoomCreator) {
	req := new(models.CreateChatRoomRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	chatRoom := chatRoomCreator.CreateChatRoomWrapper(req)
	c.JSON(http.StatusCreated, chatRoom)
}
