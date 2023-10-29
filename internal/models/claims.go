package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	// User ID
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}
