package handlers

import (
	"bytes"
	"chat-app/models"
	"chat-app/services"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ChatRoomCreatorStructMock struct {
	CookieUser string
}

var CreateChatRoomMock func(req *models.CreateChatRoomRequest) (models.ChatRoom, error)

func (m ChatRoomCreatorStructMock) CreateChatRoomWrapper(req *models.CreateChatRoomRequest) (models.ChatRoom, error) {
	return CreateChatRoomMock(req)
}

func TestCreateChatRoomHandler_HappyPath(t *testing.T) {
	CreateChatRoomMock = func(req *models.CreateChatRoomRequest) (models.ChatRoom, error) {
		return models.ChatRoom{}, nil
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

func TestCreateChatRoomHandler_BadRequest(t *testing.T) {
	CreateChatRoomMock = func(req *models.CreateChatRoomRequest) (models.ChatRoom, error) {
		return models.ChatRoom{}, nil
	}
	newChatRoomReq := &models.CreateChatRoomRequest{}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	jsonValue, _ := json.Marshal(newChatRoomReq)
	req, _ := http.NewRequest("POST", "/createChatRoom", bytes.NewBuffer(jsonValue))
	c.Request = req

	CreateChatRoomHandler(c, &services.ChatRoomCreatorStruct{CookieUser: "Ola"})

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestCreateChatRoomHandler_CreationError(t *testing.T) {
	CreateChatRoomMock = func(req *models.CreateChatRoomRequest) (models.ChatRoom, error) {
		return models.ChatRoom{}, errors.New("Error when saving the new chat room")
	}
	newChatRoomReq := &models.CreateChatRoomRequest{
		Title: "Some title",
	}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	jsonValue, _ := json.Marshal(newChatRoomReq)
	req, _ := http.NewRequest("POST", "/createChatRoom", bytes.NewBuffer(jsonValue))
	c.Request = req

	CreateChatRoomHandler(c, &ChatRoomCreatorStructMock{CookieUser: "Ola"})

	actualErr := recorder.Body.String()
	expectedErr := "\"Error when saving the new chat room\""

	assert.Equal(t, expectedErr, actualErr)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
