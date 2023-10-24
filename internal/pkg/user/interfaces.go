package user

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

type UserUsecase interface {
	// GetProfile(context.Context, uuid.UUID) (models.Profile, error) // take from pkg auth
	UpdatePhoto(context.Context, uuid.UUID, uuid.UUID) error
	UpdateInfo(context.Context, uuid.UUID, models.ProfileInfo) error
}

type UserRepo interface {
	// ReadProfile(context.Context, uuid.UUID) (models.Profile, error) // take from pkg auth
	UpdatePhoto(context.Context, uuid.UUID) error
	UpdateInfo(context.Context, uuid.UUID, models.UserInfo) error
}
