package api

import (
	"net/http"
	"strconv"

	"github.com/3011/chatroom-go/internal/api/jwt"
	"github.com/3011/chatroom-go/internal/db"
	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if len(username) < 3 || len(password) < 6 {
		c.JSON(200, "username or password err")
		return
	}

	if db.UserRegister(username, password) {
		c.JSON(200, "success")
		return
	}
	http.NotFound(c.Writer, c.Request)
	return
}

func UserLoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if len(username) < 3 && len(password) < 6 {
		c.JSON(200, "Login fail")
		return
	}
	id := db.UserLogin(username, password)
	if id == 0 {
		c.JSON(200, "Login fail")
		return
	}
	token, err := jwt.SignToken(id)
	if err != nil {
		c.JSON(200, "Login fail")
		return
	}
	cla, err := jwt.ParseToken(token)
	if err != nil {
		c.JSON(200, "Login fail token")
		return
	}
	c.JSON(200, (*cla)["exp"])

	c.JSON(200, "Logined: "+strconv.Itoa(int(id))+"\n"+token)

	// println((*cl).Id)
	return
}
