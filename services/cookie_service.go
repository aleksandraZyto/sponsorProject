package services

import (
	"github.com/gin-gonic/gin"
)

func SetCookieHandler(c *gin.Context) {
	cookie, err := c.Cookie("user")
	if err != nil {
		c.SetCookie("user", "yangyanxing", 3600, "/", "127.0.0.1", false, false)
		c.JSON(200, gin.H{"msg": "cookie set successfully"})
	} else {
		c.JSON(200, gin.H{"msg": cookie})
	}
}

func GetCookieHandler(c *gin.Context) {

}
