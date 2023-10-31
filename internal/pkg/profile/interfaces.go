package profile

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/auth_mock.go -package mock

type ProfileUsecase interface {
	GetProfile(context.Context, uuid.UUID) (*models.Profile, error)
	UpdatePhoto(ctx context.Context, userID uuid.UUID, filePhotoByte []byte, fileType string) error
	UpdateInfo(context.Context, uuid.UUID, *models.ProfileInfo) error
}

type ProfileRepo interface {
	CreateProfile(context.Context, *models.Profile) error
	ReadProfile(context.Context, uuid.UUID) (*models.Profile, error)
	GetProfileIdByLogin(context.Context, string) (uuid.UUID, error)
	UpdateProfile(context.Context, *models.Profile) error
	UpdatePhoto(context.Context, uuid.UUID, string) error
}
