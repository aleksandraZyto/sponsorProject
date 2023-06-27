package handlers

import (
	"bytes"
	"chat-app/models"
	"chat-app/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ChatRoomCreatorStructMock struct{}

var CreateChatRoomMock func(req *models.CreateChatRoomRequest) models.ChatRoom

func (m ChatRoomCreatorStructMock) CreateChatRoomWrapper(req *models.CreateChatRoomRequest) models.ChatRoom {
	return CreateChatRoomMock(req)
}

func TestCreateChatRoomHandler(t *testing.T) {
	CreateChatRoomMock = func(req *models.CreateChatRoomRequest) models.ChatRoom {
		return models.ChatRoom{}
	}
	expectedNewRoom := models.ChatRoom{}
	newChatRoomReq := &models.CreateChatRoomRequest{
		Title: "Some title",
	}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	jsonValue, _ := json.Marshal(newChatRoomReq)
	req, _ := http.NewRequest("POST", "/createChatRoom", bytes.NewBuffer(jsonValue))
	c.Request = req

	CreateChatRoomHandler(c, &services.ChatRoomCreatorStruct{CookieUser: "Ola"})

	actualNewRoom := models.ChatRoom{}
	json.Unmarshal([]byte(recorder.Body.String()), &actualNewRoom)

	assert.Equal(t, expectedNewRoom, actualNewRoom)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}
