package jwts

import (
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/golang-jwt/jwt"
	"time"
)

type ClaimsWithLogin struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

func MakeToken(user models.User) (string, error) {
	claims := ClaimsWithLogin{
		Login: user.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//secret, stated := os.LookupEnv("TOKEN_SECRET")
	//if !stated {
	//	return "", errors.New("NoSecretKey")
	//}
	str, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", errors.New("NoSecretKey")
	}
	return str, nil
}
