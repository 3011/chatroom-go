package db

import "gorm.io/gorm"

type message struct {
	gorm.Model

	ChatID int
	FromID int

	ReplyID int
	Text    string
}

func SendMessage() {
	// db.Create()
}
