package services

import "chat-app/models"

type ChatRoomCreatorStruct struct {
	CookieUser string
}

type ChatRoomCreator interface {
	CreateChatRoomWrapper(req *models.CreateChatRoomRequest) (models.ChatRoom, error)
}

func (crcs ChatRoomCreatorStruct) CreateChatRoomWrapper(_ *models.CreateChatRoomRequest) (models.ChatRoom, error) { // TODO: Return an explicit error
	// CreateChatRoom()
	return models.ChatRoom{}, nil
}
