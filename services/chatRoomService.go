package services

import "chat-app/models"

type ChatRoomCreatorStruct struct {
	CookieUser string
}

type ChatRoomCreator interface {
	CreateChatRoomWrapper(req *models.CreateChatRoomRequest) models.ChatRoom
}

func (crcs ChatRoomCreatorStruct) CreateChatRoomWrapper(_ *models.CreateChatRoomRequest) models.ChatRoom { //TODO: Implement error
	return models.ChatRoom{}
}
