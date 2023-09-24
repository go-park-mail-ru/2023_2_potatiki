package auth

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks/auth_mock.go -package=mock

type AuthUsecase interface {
	Login(context.Context, models.User) (models.Profile, error)
	SingUp(context.Context, models.User) (models.Profile, error)
	GetProfile(context.Context, uuid.UUID) (models.Profile, error)
}

type AuthRepo interface {
	CreateUser(context.Context, models.User) (models.Profile, error)
	CheckUser(context.Context, models.User) (models.Profile, error)
	ReadProfile(context.Context, uuid.UUID) (models.Profile, error)
}
