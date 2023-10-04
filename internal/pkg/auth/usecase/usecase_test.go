package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/mocks"
	"github.com/golang/mock/gomock"
)

func TestAuthUsecase_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockAuthRepo(ctrl)
	cfg := mock.NewMockAuthConfig(ctrl)
	cfg.EXPECT().GetAccessExpirationTime().Return(time.Second)
	cfg.EXPECT().GetJwtAccess().Return("")

	repo.EXPECT().CreateUser(gomock.Any(), models.User{
		Login:        "iudsbfiwhdbfi",
		PasswordHash: "hafikyagdfiaysgf",
	}).Return(models.Profile{}, nil)
	uc := NewAuthUsecase(repo, cfg)

	profile, token, _, err := uc.SignUp(context.Background(), models.User{
		Login:        "iudsbfiwhdbfi",
		PasswordHash: "hafikyagdfiaysgf",
	})
	assert.Nil(t, err)
	assert.NotNil(t, profile)
	assert.NotEmpty(t, token)
}

func TestAuthUsecase_SignUpBadRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockAuthRepo(ctrl)
	cfg := mock.NewMockAuthConfig(ctrl)
	cfg.EXPECT().GetAccessExpirationTime().Return(time.Second)
	cfg.EXPECT().GetJwtAccess().Return("")

	repo.EXPECT().CreateUser(gomock.Any(), models.User{
		Login:        "iudsbfiwhdbfi",
		PasswordHash: "hafikyagdfiaysgf",
	}).Return(models.Profile{}, errors.New("bad request"))

	uc := NewAuthUsecase(repo, cfg)

	profile, token, _, err := uc.SignUp(context.Background(), models.User{
		Login:        "iudsbfiwhdbfi",
		PasswordHash: "hafikyagdfiaysgf",
	})
	assert.NotNil(t, err)
	assert.NotNil(t, profile)
	assert.Empty(t, token)
}

func TestAuthUsecase_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockAuthRepo(ctrl)
	cfg := mock.NewMockAuthConfig(ctrl)
	cfg.EXPECT().GetAccessExpirationTime().Return(time.Second)
	cfg.EXPECT().GetJwtAccess().Return("")

	repo.EXPECT().CheckUser(gomock.Any(), models.User{
		Login:        "iudsbfiwhdbfi",
		PasswordHash: "hafikyagdfiaysgf",
	}).Return(models.Profile{}, nil)
	uc := NewAuthUsecase(repo, cfg)

	profile, token, _, err := uc.SignIn(context.Background(), models.User{
		Login:        "iudsbfiwhdbfi",
		PasswordHash: "hafikyagdfiaysgf",
	})
	assert.Nil(t, err)
	assert.NotNil(t, profile)
	assert.NotEmpty(t, token)
}

func TestAuthUsecase_SigInBadRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockAuthRepo(ctrl)
	cfg := mock.NewMockAuthConfig(ctrl)
	cfg.EXPECT().GetAccessExpirationTime().Return(time.Second)
	cfg.EXPECT().GetJwtAccess().Return("")

	repo.EXPECT().CheckUser(gomock.Any(), models.User{
		Login:        "iudsbfiwhdbfi",
		PasswordHash: "hafikyagdfiaysgf",
	}).Return(models.Profile{}, errors.New("bad request"))
	uc := NewAuthUsecase(repo, cfg)

	profile, token, _, err := uc.SignIn(context.Background(), models.User{
		Login:        "iudsbfiwhdbfi",
		PasswordHash: "hafikyagdfiaysgf",
	})
	assert.NotNil(t, err)
	assert.NotNil(t, profile)
	assert.Empty(t, token)
}
