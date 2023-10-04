package usecase

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/mocks"
	"github.com/golang/mock/gomock"
)

func TestAuthUsecase_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockAuthRepo(ctrl)
	cfg := mock.NewMockAuthConfig(ctrl)
	repo.EXPECT().CreateUser(gomock.Any(), models.User{}).Return(models.Profile{}, nil)
	uc := NewAuthUsecase(repo, cfg)

	profile, token, exp, err := uc.SignUp(context.Background(), models.User{})
	fmt.Println(profile, token, exp, err)
}

func TestAuthUsecase_SignUpBadRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockAuthRepo(ctrl)
	cfg := mock.NewMockAuthConfig(ctrl)
	repo.EXPECT().CreateUser(gomock.Any(), models.User{}).Return(models.Profile{}, errors.New("bad request"))

	uc := NewAuthUsecase(repo, cfg)

	profile, token, exp, err := uc.SignUp(context.Background(), models.User{})
	fmt.Println(profile, token, exp, err)
}
