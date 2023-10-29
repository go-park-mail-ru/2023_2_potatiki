package auth

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/auth_mock.go -package mock

type AuthUsecase interface {
	SignIn(context.Context, *models.User) (*models.Profile, string, time.Time, error)
	SignUp(context.Context, *models.User) (*models.Profile, string, time.Time, error)
}

type AuthAuther interface {
	GenerateToken(*models.Profile) (string, time.Time, error)
	GetClaims(string) (*models.Claims, error)
}

type AuthConfig interface {
	GetAccessExpirationTime() time.Duration
	GetJwtAccess() string
}
