package jwter

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -source jwter.go -destination ./mocks/jwt_mock.go -package mock

type Configer interface {
	GetTTL() time.Duration
	GetSecret() string
	GetIssuer() string
}

type JWTer interface {
	GenerateToken(*models.Profile) (string, time.Time, error)
	GetClaims(string) (*models.Claims, error)
}
type JWTManager struct {
	ttl    time.Duration
	secret string
	issuer string
}

func New(cfg Configer) *JWTManager {
	return &JWTManager{
		ttl:    cfg.GetTTL(),
		secret: cfg.GetSecret(),
		issuer: cfg.GetIssuer(),
	}
}

func (j *JWTManager) GenerateToken(profile *models.Profile) (string, time.Time, error) {
	expirationTime := time.Now().UTC().Add(j.ttl)

	claims := &models.Claims{
		ID: profile.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    j.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", time.Now(), err
	}

	return tokenStr, expirationTime, nil
}

func (j *JWTManager) getKeyFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.secret), nil
	}
}

func (j *JWTManager) GetClaims(tokenString string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, j.getKeyFunc())
	if err != nil {
		err = fmt.Errorf("error happened in jwt.ParseWithClaims: %w", err)

		return nil, err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("error in GetClaims, invalid token or Claims.(*Claims) not cast")
}
