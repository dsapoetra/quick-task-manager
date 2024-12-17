package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userId uint, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
