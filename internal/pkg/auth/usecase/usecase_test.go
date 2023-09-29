package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestAuthUsecase_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock.NewMockAuthRepo(ctrl)
	repo.EXPECT().CreateUser(gomock.Any(), models.User{}).Return(models.Profile{}, nil)

	uc := NewAuthUsecase(repo)

	profile, err := uc.SignUp(context.Background(), models.User{})
	fmt.Println(profile, err)
}

func TestAuthUsecase_SignUpBadRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock.NewMockAuthRepo(ctrl)
	repo.EXPECT().CreateUser(gomock.Any(), models.User{}).Return(models.Profile{}, errors.New("bad request"))

	uc := NewAuthUsecase(repo)

	profile, err := uc.SignUp(context.Background(), models.User{})
	fmt.Println(profile, err)
}
