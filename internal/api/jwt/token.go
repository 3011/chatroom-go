package jwt

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret = []byte("adadad")

// 生成Token
func SignToken(userID uint) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 1).Unix(),
		Id:        strconv.Itoa(int(userID)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 检验Token
func ParseToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	} else {
		return nil, errors.New("claims error")
	}

}
