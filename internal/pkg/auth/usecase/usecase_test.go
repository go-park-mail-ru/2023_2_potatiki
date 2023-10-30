package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mockAuth "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/mocks"
	mockUser "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/user/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/hasher"
	"github.com/golang/mock/gomock"
)

func TestAuthUsecase_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockUser.NewMockUserRepo(ctrl)
	cfg := mockAuth.NewMockAuthConfig(ctrl)
	cfg.EXPECT().GetAccessExpirationTime().Return(time.Second)
	cfg.EXPECT().GetJwtAccess().Return("")

	_ = &models.Profile{
		Id:           uuid.New(),
		Login:        "iudsbfiwhdbfi",
		Description:  "",
		ImgSrc:       "default.png",
		PasswordHash: hasher.HashPass("hafikyagdfiaysgf"),
	}

	repo.EXPECT().CreateProfile(gomock.Any(), gomock.Any()).Return(nil)
	uc := NewAuthUsecase(repo, cfg)

	profile, token, _, err := uc.SignUp(context.Background(), &models.User{
		Login:    "iudsbfiwhdbfi",
		Password: "hafikyagdfiaysgf",
	})
	assert.Nil(t, err)
	assert.NotNil(t, profile)
	assert.NotEmpty(t, token)
}

func TestAuthUsecase_SignUpBadRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockUser.NewMockUserRepo(ctrl)
	cfg := mockAuth.NewMockAuthConfig(ctrl)
	cfg.EXPECT().GetAccessExpirationTime().Return(time.Second)
	cfg.EXPECT().GetJwtAccess().Return("")

	repo.EXPECT().CreateProfile(gomock.Any(), gomock.Any()).Return(errors.New("bad request"))

	uc := NewAuthUsecase(repo, cfg)

	profile, token, _, err := uc.SignUp(context.Background(), &models.User{
		Login:    "iudsbfiwhdbfi",
		Password: "hafikyagdfiaysgf",
	})
	assert.NotNil(t, err)
	assert.NotNil(t, profile)
	assert.Empty(t, token)
}

func TestAuthUsecase_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockUser.NewMockUserRepo(ctrl)
	cfg := mockAuth.NewMockAuthConfig(ctrl)
	cfg.EXPECT().GetAccessExpirationTime().Return(time.Second)
	cfg.EXPECT().GetJwtAccess().Return("")

	user := &models.User{
		Login:    "iudsbfiwhdbfi",
		Password: "hafikyagdfiaysgf",
	}
	Id := uuid.New()

	p := &models.Profile{
		Id:           uuid.New(),
		Login:        "iudsbfiwhdbfi",
		Description:  "",
		ImgSrc:       "default.png",
		PasswordHash: hasher.HashPass("hafikyagdfiaysgf"),
	}

	repo.EXPECT().GetProfileIdByUser(gomock.Any(), user).Return(Id, nil)
	repo.EXPECT().ReadProfile(gomock.Any(), Id).Return(p, nil)
	uc := NewAuthUsecase(repo, cfg)

	profile, token, _, err := uc.SignIn(context.Background(), user)
	assert.Nil(t, err)
	assert.NotNil(t, profile)
	assert.NotEmpty(t, token)
}

func TestAuthUsecase_SigInBadRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockUser.NewMockUserRepo(ctrl)
	cfg := mockAuth.NewMockAuthConfig(ctrl)
	cfg.EXPECT().GetAccessExpirationTime().Return(time.Second)
	cfg.EXPECT().GetJwtAccess().Return("")

	user := &models.User{
		Login:    "iudsbfiwhdbfi",
		Password: "hafikyagdfiaysgf",
	}
	Id := uuid.New()

	_ = &models.Profile{
		Id:           uuid.New(),
		Login:        "iudsbfiwhdbfi",
		Description:  "",
		ImgSrc:       "default.png",
		PasswordHash: hasher.HashPass("hafikyagdfiaysgf"),
	}

	repo.EXPECT().GetProfileIdByUser(gomock.Any(), user).Return(Id, errors.New("bad"))
	uc := NewAuthUsecase(repo, cfg)

	profile, token, _, err := uc.SignIn(context.Background(), user)
	assert.NotNil(t, err)
	assert.NotNil(t, profile)
	assert.Empty(t, token)
}
