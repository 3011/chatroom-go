package main

import (
	"github.com/3011/chatroom-go/internal/config"
	"github.com/3011/chatroom-go/internal/db"
	"github.com/3011/chatroom-go/internal/router"
)

func main() {
	config.Init()
	db.InitDB()
	r := router.NewRouter()
	_ = r.Run(config.Config.HTTPPort)
}
