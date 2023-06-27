package handlers

import (
	"bytes"
	"chat-app/models"
	"chat-app/repos"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RegistererStructMock struct {
	mock.Mock
}

func (m *RegistererStructMock) WrapperRegister(_ *models.RegisterRequest, _ repos.UserRepository) (models.User, error) {
	return models.User{}, nil
}

func TestRegister_HappyPath(t *testing.T) {
	registerReq := &models.RegisterRequest{
		LoginData: models.LoginRequest{
			Username: "Olka",
			Password: "1111",
		},
		Name: "Alex",
	}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	jsonValue, _ := json.Marshal(registerReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	c.Request = req

	RegisterHandler(c, &RegistererStructMock{})
	assert.Equal(t, http.StatusCreated, recorder.Code)
	//fmt.Println("r body:", recorder.Body)
	//assert.Equal(t, "", recorder.Body)
}

// TODO: Sad path
//args := m.Called(req, repo) TODO: How to use this .Called
//m.On("WrapperRegister", registerReq, repo).Return("Olkaaaaa", errors.New("so,e error")) TODO: How to use this .Return
