package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	models "chat-app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/valyala/fasthttp"
)

func TestHelloHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/hello", Home)

	resp, err := app.Test(httptest.NewRequest("GET", "/hello", nil))

	utils.AssertEqual(t, nil, err, "Error")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}

func TestRegistrationHandler(t *testing.T) {
	app := fiber.New()
	app.Post("/register", Home)
	password := "123"
	reqBody := models.User{Username: "ola", Password: password, Name: "Aleksadra"}
	jsonValue, _ := json.Marshal(&reqBody)

	resp, err := app.Test(httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue)))
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	fmt.Println("string(c.Response().Body())")
	fmt.Println(string(c.Response().String()))

	utils.AssertEqual(t, nil, err, "Error")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}
