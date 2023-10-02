package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Auther struct {
	ttl    time.Duration
	secret string
}

func NewAuther(cfg config.Auther) *Auther {
	return &Auther{
		ttl:    cfg.AccessExpirationTime,
		secret: cfg.JwtAccess,
	}
}

type claims struct {
	// User ID
	ID uuid.UUID `json:"id"`

	jwt.RegisteredClaims
}

func (a *Auther) generateToken(profile *models.Profile) (string, time.Time, error) {
	expirationTime := time.Now().UTC().Add(a.ttl)

	claims := &claims{
		profile.Id,
		jwt.RegisteredClaims{
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

func (a *Auther) getClaims(tokenString string) (*claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, a.getKeyFunc())
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("error in GetClaims, invalid token or Claims.(*Claims) not cast")
	}
}
