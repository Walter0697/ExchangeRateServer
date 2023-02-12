package helper

import (
	"chaos/backend/config"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTInfo struct {
	Username string
	UserId   uint
}

func GenerateToken(username string, id uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["userid"] = id
	ttl := 100000 * time.Second
	claims["exp"] = time.Now().UTC().Add(ttl).Unix()
	tokenString, err := token.SignedString([]byte(config.Data.App.JWT))
	if err != nil {
		log.Fatal("Error in generating key for " + username)
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenStr string) (*JWTInfo, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Data.App.JWT), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		id := claims["userid"].(float64)
		result := JWTInfo{
			Username: username,
			UserId:   uint(id),
		}
		return &result, nil
	} else {
		return nil, err
	}
}
