package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	ID uuid.UUID `json:"id"` // Profile ID
	jwt.RegisteredClaims
}
