package api

import (
	"net/http"
	"strconv"

	"github.com/3011/chatroom-go/internal/db"
	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if db.UserRegister(username, password) {
		c.JSON(200, "success")
		return
	}
	http.NotFound(c.Writer, c.Request)
	return
}

func UserLoginHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	id := db.UserLogin(username, password)
	if id == 0 {
		c.JSON(200, "Login fail")

		return
	}
	c.JSON(200, "Logined: "+strconv.Itoa(int(id)))
	return
}
