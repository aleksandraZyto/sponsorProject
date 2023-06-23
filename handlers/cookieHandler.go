package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetUserCookie(c *gin.Context, username string) {
	_, err := GetUserCookie(c) // TODO: Do we need this, what is this for
	if err != nil {
		c.SetCookie("user", username, 3600, "/", "127.0.0.1", false, false)
	}
}

func GetUserCookie(c *gin.Context) (string, error) {
	return c.Cookie("user")
}
