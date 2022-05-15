package db

import (
	"github.com/3011/chatroom-go/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open(config.Config.DBFileName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&user{})
	db.AutoMigrate(&message{})
}
