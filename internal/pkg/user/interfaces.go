package user

import (
	"context"
	"github.com/google/uuid"
)

type UserUsecase interface {
	// GetProfile(context.Context, uuid.UUID) (models.Profile, error) // take from pkg auth
	UpdatePhoto(context.Context, uuid.UUID) error
	UpdatePassword(context.Context, uuid.UUID, string) error
	UpdateDescription(context.Context, uuid.UUID, string) error
}

type UserRepo interface {
	// ReadProfile(context.Context, uuid.UUID) (models.Profile, error) // take from pkg auth
	UpdatePhoto(context.Context, uuid.UUID) error
	UpdatePassword(context.Context, uuid.UUID, string) error
	UpdateDescription(context.Context, uuid.UUID, string) error
}