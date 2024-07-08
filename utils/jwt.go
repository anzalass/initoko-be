package utils

import (
	"initoko/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userId string, email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"role":    role,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(3 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accesToken, err := token.SignedString([]byte(config.InitConfig().Secret))
	if err != nil {
		return "", err
	}
	return accesToken, nil
}

func ValidateToken(tokenstring string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.InitConfig().Secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
