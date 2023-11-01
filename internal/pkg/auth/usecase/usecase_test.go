package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mockAuth "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/mocks"
	mockProfile "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/profile/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/hasher"
	"github.com/golang/mock/gomock"
)

func TestAuthUsecase_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockProfile.NewMockProfileRepo(ctrl)
	cfg := mockAuth.NewMockAuthConfig(ctrl)
	cfg.EXPECT().GetAccessExpirationTime().Return(time.Second)
	cfg.EXPECT().GetJwtAccess().Return("")

	_ = &models.Profile{
		Id:           uuid.NewV4(),
		Login:        "iudsbfiwhdbfi",
		Description:  "",
		ImgSrc:       "default.png",
		PasswordHash: hasher.HashPass("hafikyagdfiaysgf"),
	}

	repo.EXPECT().CreateProfile(gomock.Any(), gomock.Any()).Return(nil)
	uc := NewAuthUsecase(repo, cfg)

	profile, token, _, err := uc.SignUp(context.Background(), &models.SignUpPayload{
		Login:    "iudsbfiwhdbfi",
		Phone:    "+79912245678",
		Password: "hafikyagdfiaysgf",
	})
	assert.Nil(t, err)
	assert.NotNil(t, profile)
	assert.NotEmpty(t, token)
}

func TestAuthUsecase_SignUpBadRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockProfile.NewMockProfileRepo(ctrl)
	cfg := mockAuth.NewMockAuthConfig(ctrl)
	cfg.EXPECT().GetAccessExpirationTime().Return(time.Second)
	cfg.EXPECT().GetJwtAccess().Return("")

	repo.EXPECT().CreateProfile(gomock.Any(), gomock.Any()).Return(errors.New("bad request"))

	uc := NewAuthUsecase(repo, cfg)

	profile, token, _, err := uc.SignUp(context.Background(), &models.SignUpPayload{
		Login:    "iudsbfiwhdbfi",
		Phone:    "+79912245678",
		Password: "hafikyagdfiaysgf",
	})
	assert.NotNil(t, err)
	assert.NotNil(t, profile)
	assert.Empty(t, token)
}

func TestAuthUsecase_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockProfile.NewMockProfileRepo(ctrl)
	cfg := mockAuth.NewMockAuthConfig(ctrl)
	cfg.EXPECT().GetAccessExpirationTime().Return(time.Second)
	cfg.EXPECT().GetJwtAccess().Return("")

	payload := &models.SignInPayload{
		Login:    "iudsbfiwhdbfi",
		Password: "hafikyagdfiaysgf",
	}
	Id := uuid.NewV4()

	p := &models.Profile{
		Id:           uuid.NewV4(),
		Login:        "iudsbfiwhdbfi",
		Description:  "",
		ImgSrc:       "default.png",
		PasswordHash: hasher.HashPass("hafikyagdfiaysgf"),
	}

	repo.EXPECT().GetProfileIdByLogin(gomock.Any(), payload.Login).Return(Id, nil)
	repo.EXPECT().ReadProfile(gomock.Any(), Id).Return(p, nil)
	uc := NewAuthUsecase(repo, cfg)

	profile, token, _, err := uc.SignIn(context.Background(), payload)
	assert.Nil(t, err)
	assert.NotNil(t, profile)
	assert.NotEmpty(t, token)
}

func TestAuthUsecase_SigInBadRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockProfile.NewMockProfileRepo(ctrl)
	cfg := mockAuth.NewMockAuthConfig(ctrl)
	cfg.EXPECT().GetAccessExpirationTime().Return(time.Second)
	cfg.EXPECT().GetJwtAccess().Return("")

	payload := &models.SignInPayload{
		Login:    "iudsbfiwhdbfi",
		Password: "hafikyagdfiaysgf",
	}
	Id := uuid.NewV4()

	_ = &models.Profile{
		Id:           uuid.NewV4(),
		Login:        "iudsbfiwhdbfi",
		Description:  "",
		ImgSrc:       "default.png",
		PasswordHash: hasher.HashPass("hafikyagdfiaysgf"),
	}

	repo.EXPECT().GetProfileIdByLogin(gomock.Any(), payload.Login).Return(Id, errors.New("bad"))
	uc := NewAuthUsecase(repo, cfg)

	profile, token, _, err := uc.SignIn(context.Background(), payload)
	assert.NotNil(t, err)
	assert.NotNil(t, profile)
	assert.Empty(t, token)
}
