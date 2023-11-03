package jwter

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
)

//go:generate mockgen -source jwter.go -destination ./mocks/jwt_mock.go -package mock

type Configer interface {
	GetTTL() time.Duration
	GetSecret() string
	GetIssuer() string
}

type JWTer interface {
	EncodeAuthToken(uuid.UUID) (string, time.Time, error)
	DecodeAuthToken(string) (uuid.UUID, error)
	EncodeCSRFToken(string) (string, time.Time, error)
	DecodeCSRFToken(string) (string, error)
}

type jwtManager struct {
	ttl    time.Duration
	secret string
	issuer string
}

func New(cfg Configer) *jwtManager {
	return &jwtManager{
		ttl:    cfg.GetTTL(),
		secret: cfg.GetSecret(),
		issuer: cfg.GetIssuer(),
	}
}

type authClaims struct {
	ID uuid.UUID `json:"id"` // Profile ID
	jwt.RegisteredClaims
}

func (j *jwtManager) EncodeAuthToken(ID uuid.UUID) (string, time.Time, error) {
	expirationTime := time.Now().UTC().Add(j.ttl)
	claims := &authClaims{
		ID: ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    j.issuer,
		},
	}
	return j.generateToken(claims)
}

func (j *jwtManager) DecodeAuthToken(tokenString string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authClaims{}, j.getKeyFunc())
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error happened in jwt.ParseWithClaims: %w", err)
	}

	if claims, ok := token.Claims.(*authClaims); ok && token.Valid {
		return claims.ID, nil
	}

	return uuid.UUID{}, errors.New("error in GetClaims, invalid token or Claims.(*Claims) not cast")
}

type csrfClaims struct {
	UserAgent string `json:"userAgent"` // request UserAgent
	jwt.RegisteredClaims
}

func (j *jwtManager) EncodeCSRFToken(UserAgent string) (string, time.Time, error) {
	expirationTime := time.Now().UTC().Add(j.ttl)
	claims := &csrfClaims{
		UserAgent: UserAgent,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    j.issuer,
		},
	}
	return j.generateToken(claims)
}

func (j *jwtManager) DecodeCSRFToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &csrfClaims{}, j.getKeyFunc())
	if err != nil {
		return "", fmt.Errorf("error happened in jwt.ParseWithClaims: %w", err)
	}

	if claims, ok := token.Claims.(*csrfClaims); ok && token.Valid {
		return claims.UserAgent, nil
	}

	return "", errors.New("error in GetClaims, invalid token or Claims.(*Claims) not cast")
}

func (j *jwtManager) generateToken(claims jwt.Claims) (string, time.Time, error) {
	expirationTime := time.Now().UTC().Add(j.ttl)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", time.Now(), err
	}

	return tokenStr, expirationTime, nil
}

func (j *jwtManager) getKeyFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.secret), nil
	}
}
