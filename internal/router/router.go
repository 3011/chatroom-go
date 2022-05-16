package router

import (
	"net/http"

	"github.com/3011/chatroom-go/internal/api"
	"github.com/3011/chatroom-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery(), gin.Logger())

	v1 := r.Group("/")
	{
		r.LoadHTMLGlob("internal/templates/*")
		v1.GET("index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "Main website",
			})
		})
		v1.GET("ping", middleware.AuthMiddleware(), func(ctx *gin.Context) {
			userid, ok := ctx.Get("userid")
			if !ok {
				ctx.JSON(404, gin.H{"data": gin.H{"userid": userid}})
				return
			}
			ctx.JSON(200, gin.H{"data": gin.H{"userid": userid}})
		})
		// v1.POST("user/register", api.UserRegister)
		v1.GET("ws", api.WsHandler)
		v1.POST("signup", api.UserRegisterHandler)
		v1.POST("signin", api.UserLoginHandler)
		// v1.POST("login", api.UserLogin)
	}
	return r
}
