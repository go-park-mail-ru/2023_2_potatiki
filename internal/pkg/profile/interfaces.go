package profile

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/auth_mock.go -package mock

type ProfileUsecase interface {
	GetProfile(context.Context, uuid.UUID) (*models.Profile, error)
	UpdatePhoto(ctx context.Context, userID uuid.UUID, filePhotoByte []byte, fileType string) (*models.Profile, error)
	UpdateData(context.Context, uuid.UUID, *models.UpdateProfileDataPayload) (*models.Profile, error)
}

type ProfileRepo interface {
	CreateProfile(context.Context, *models.Profile) error
	ReadProfile(context.Context, uuid.UUID) (*models.Profile, error)
	GetProfileIdByLogin(context.Context, string) (uuid.UUID, error)
	UpdateProfile(context.Context, *models.Profile) error
	UpdatePhoto(context.Context, uuid.UUID, string) error
}

type ProfileConfig interface {
	GetPhotosFilePath() string
}

var (
	ErrorForbiddenExtension = errors.New("this file extension is not allowed")
	ErrBadUpdateData        = errors.New("incorrect update data profile payload")
)
