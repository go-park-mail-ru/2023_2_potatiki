package auth

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/auth_mock.go -package mock

type AuthUsecase interface {
	SignIn(context.Context, *models.SignInPayload) (*models.Profile, string, time.Time, error)
	SignUp(context.Context, *models.SignUpPayload) (*models.Profile, string, time.Time, error)
	CheckAuth(context.Context, uuid.UUID) (*models.Profile, error)
}

type AuthAuther interface {
	GenerateToken(*models.Profile) (string, time.Time, error)
	GetClaims(string) (*models.Claims, error)
}

type AuthConfig interface {
	GetAccessExpirationTime() time.Duration
	GetJwtAccess() string
}
