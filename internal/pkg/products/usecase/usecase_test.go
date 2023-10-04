package usecase

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestProducts_GetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockProductsRepo(ctrl)
	repo.EXPECT().ReadProduct(gomock.Any(), models.Product{}).Return(models.Product{}, nil)

	uc := NewProductsUsecase(repo)

	product, err := uc.GetProduct(context.Background(), uuid.UUID{})
	fmt.Println(product, err)
}

func TestAuthUsecase_SignUpBadRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockProductsRepo(ctrl)
	repo.EXPECT().ReadProduct(gomock.Any(), models.User{}).Return(models.Product{}, errors.New("bad request"))

	uc := NewProductsUsecase(repo)

	product, err := uc.GetProduct(context.Background(), uuid.UUID{})
	fmt.Println(product, err)
}
