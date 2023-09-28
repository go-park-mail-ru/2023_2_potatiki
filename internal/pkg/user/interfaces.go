package user

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

type UserUsecase interface {
	GetProfile(context.Context, uuid.UUID) (models.Profile, error)
}

type UserRepo interface {
	ReadUser(context.Context, uuid.UUID) (models.Profile, error)
}
