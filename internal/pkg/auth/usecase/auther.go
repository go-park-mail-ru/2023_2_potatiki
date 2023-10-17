package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/golang-jwt/jwt/v5"
)

type Auther struct {
	ttl    time.Duration
	secret string
}

func NewAuther(cfg auth.AuthConfig) *Auther {
	return &Auther{
		ttl:    cfg.GetAccessExpirationTime(), //cfg.AccessExpirationTime,
		secret: cfg.GetJwtAccess(),            //cfg.JwtAccess,
	}
}

func (a *Auther) GenerateToken(profile *models.Profile) (string, time.Time, error) {
	expirationTime := time.Now().UTC().Add(a.ttl)

	claims := &models.Claims{
		ID: profile.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return "", time.Now(), err
	}

	return tokenStr, expirationTime, nil
}

func (a *Auther) getKeyFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.secret), nil
	}
}

func (a *Auther) GetClaims(tokenString string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, a.getKeyFunc())
	if err != nil {
		err = fmt.Errorf("error happened in jwt.ParseWithClaims: %w", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("error in GetClaims, invalid token or Claims.(*Claims) not cast")
}
