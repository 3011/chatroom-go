package api

import (
	"github.com/3011/chatroom-go/internal/db"
	"github.com/3011/chatroom-go/internal/jwt"
	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if len(username) < 3 || len(password) < 6 {
		responseFail(ctx, "The length of the username or password is illegal", nil)
		return
	}

	if !db.UserRegister(username, password) {
		responseFail(ctx, "Registration failed", nil)
		return
	}
	responseSuccess(ctx, "Registration succeeded", nil)
}

func UserLoginHandler(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if len(username) < 3 && len(password) < 6 {
		responseFail(ctx, "The length of the username or password is illegal", nil)
		return
	}
	id := db.UserLogin(username, password)
	if id == 0 {
		responseFail(ctx, "Login failed", nil)
		return
	}

	token, err := jwt.SignToken(id)
	if err != nil {
		responseFail(ctx, "Login failed", nil)
		return
	}

	claims, err := jwt.ParseToken(token)
	if err != nil {
		responseFail(ctx, "Login failed", nil)
		return
	}

	responseSuccess(ctx, "Login succeeded", gin.H{"user": gin.H{"userid": (*claims)["jti"]}})
	return
}
