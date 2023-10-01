package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func MakeToken(user models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Login,
		"exp":      time.Now().Add(time.Hour * 6).Unix(),
	})

	secret, stated := os.LookupEnv("JWT_SECRET")
	if !stated {
		return "", errors.New("NoSecretKey")
	}
	str, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", errors.New("NoSecretKey")
	}
	return str, nil
}
