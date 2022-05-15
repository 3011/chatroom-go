package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID     uint
	Socket *websocket.Conn
}

type RawData struct {
	Type int         `json:"type"`
	Data interface{} `json:"data"`
}

type Message struct {
	ChatID int    `json:"chatid"`
	Text   string `json:"text"`
}

func (c *Client) read() {
	defer func() { // 避免忘记关闭，所以要加上close
		_ = c.Socket.Close()
	}()

	for {
		c.Socket.PongHandler()

		// msgType, p, err := c.Socket.ReadMessage() // 读取json格式，如果不是json格式，会报错
		// if err != nil {
		// 	fmt.Printf("err.Error(): %v\n", err.Error())
		// }

		// fmt.Printf("msgType: %v\n", msgType)
		// fmt.Printf("p: %v\n", p)

		// c.Socket.WriteMessage(websocket.TextMessage, p)

		var rawJson json.RawMessage
		rawData := RawData{
			Data: &rawJson,
		}

		// _,msg,_:=c.Socket.ReadMessage()
		err := c.Socket.ReadJSON(&rawData) // 读取json格式，如果不是json格式，会报错
		if err != nil {
			fmt.Printf("err.Error(): %v\n", err.Error())
		}
		switch rawData.Type {
		case 1:
			var msg Message
			if err := json.Unmarshal(rawJson, &msg); err != nil {
				fmt.Printf("err.Error(): %v\n", err.Error())
			}
			fmt.Printf("%v in %v: %v\n", strconv.Itoa(int(c.ID)), msg.ChatID, msg.Text)
		}
	}
}

func (c *Client) write() {
	defer func() { // 避免忘记关闭，所以要加上close
		_ = c.Socket.Close()
	}()

	for {
		c.Socket.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("15:04:05")))
		time.Sleep(1 * time.Second)
	}
}
