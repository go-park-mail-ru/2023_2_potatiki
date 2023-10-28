package user

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/auth_mock.go -package mock

type UserUsecase interface {
	GetProfile(context.Context, uuid.UUID) (models.Profile, error)
	UpdatePhoto(ctx context.Context, userID uuid.UUID, filePhotoByte []byte, fileType string) error
	UpdateInfo(context.Context, uuid.UUID, models.ProfileInfo) error
}

type UserRepo interface {
	ReadProfile(context.Context, uuid.UUID) (models.Profile, error)
	CreateUser(context.Context, models.User) (models.Profile, error)
	CheckUser(context.Context, models.User) (models.Profile, error)
	UpdatePhoto(context.Context, uuid.UUID, string) error
	UpdateInfo(context.Context, uuid.UUID, models.UserInfo) error
}
