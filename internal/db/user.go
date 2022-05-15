package db

import (
	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Name     string
	// Avatar   string
}

func UserRegister(username string, password string) bool {
	result := db.Debug().Limit(1).Where("username = ?", username).Find(&user{})
	if result.RowsAffected == 0 {
		newUser := user{
			Username: username,
			Password: password,
		}
		result := db.Create(&newUser)
		return result.RowsAffected == 1
	}
	return false
}

func UserLogin(username string, password string) uint {
	var returnUser user
	result := db.Debug().Limit(1).Where("username = ?", username).Find(&returnUser)
	if result.RowsAffected == 1 {
		if returnUser.Password == password {
			return returnUser.ID
		}
	}
	return 0
}
