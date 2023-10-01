package usecase

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
)

func (uc *AuthUsecase) CheckToken(r *http.Request) (string, error) {
	secret, stated := os.LookupEnv("JWT_SECRET")
	if !stated {
		return "", errors.New("NoSecretKey")
	}
	cookie, err := r.Cookie("Default")
	if err != nil {
		return "", err
	}
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	}

	return "", errors.New("token is invalid")
}
