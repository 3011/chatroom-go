package api

import (
	"net/http"
	"strconv"

	"github.com/3011/chatroom-go/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WsHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	// if err != nil {
	// 	// http.NewRequestWithContext()
	// 	http.NotFound(c.Writer, c.Request)
	// 	return
	// }

	id := db.UserLogin(username, password)
	if id == 0 {
		http.NotFound(c.Writer, c.Request)
		return
	}

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { // CheckOrigin解决跨域问题
			return true
		}}).Upgrade(c.Writer, c.Request, nil) // 升级成ws协议

	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	conn.WriteMessage(websocket.TextMessage, []byte("Hi!"+strconv.Itoa(int(id))))
	// 创建一个用户实例
	client := &Client{
		ID:     id,
		Socket: conn,
	}
	// 用户注册到用户管理上
	// Manager.Register <- client
	go client.read()
	go client.write()
}
