package db

import (
	"golang.org/x/crypto/bcrypt"
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

	// 加密密码
	hedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false
	} else {
		password = string(hedPassword)
	}

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
		// 解密密码
		if bcrypt.CompareHashAndPassword([]byte(returnUser.Password), []byte(password)) == nil {
			return returnUser.ID
		}
	}
	return 0
}
