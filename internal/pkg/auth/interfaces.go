package auth

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/auth_mock.go -package mock

type AuthUsecase interface {
	SignIn(context.Context, models.User) (models.Profile, string, time.Time, error)
	SignUp(context.Context, models.User) (models.Profile, string, time.Time, error)
	CheckToken(context.Context, string) (uuid.UUID, error)
}

type AuthAuther interface {
	GenerateToken(*models.Profile) (string, time.Time, error)
	GetClaims(string) (*models.Claims, error)
}

type AuthConfig interface {
	GetAccessExpirationTime() time.Duration
	GetJwtAccess() string
}
