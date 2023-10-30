package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/repo"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order"
	"github.com/google/uuid"
)

type OrderUsecase struct {
	repoOrder order.OrderRepo
	repoCart  cart.CartRepo
}

func NewOrderUsecase(repoOrder order.OrderRepo, repoCart cart.CartRepo) *OrderUsecase {
	return &OrderUsecase{
		repoOrder: repoOrder,
		repoCart:  repoCart,
	}
}

func (uc *OrderUsecase) CreateOrder(ctx context.Context, id uuid.UUID) (models.Order, error) {
	cart, err := uc.repoCart.ReadCart(ctx, id)
	if err != nil {
		if errors.Is(err, repo.ErrCartNotFound) {
			return models.Order{}, err
		}
		err = fmt.Errorf("error happened in repoCart.ReadCart: %w", err)

		return models.Order{}, err
	}

	err = uc.repoCart.DeleteCart(ctx, cart.Id)
	if err != nil {
		err = fmt.Errorf("error happened in repoCart.DeleteCart: %w", err)

		return models.Order{}, err
	}

	order, err := uc.repoOrder.CreateOrder(ctx, cart, id)
	if err != nil {
		err = fmt.Errorf("error happened in repo.CreateOrder: %w", err)

		return models.Order{}, err
	}

	return order, nil
}

//
//func (uc *OrderUsecase) GetCurrentOrder(ctx context.Context, id uuid.UUID) (models.Order, error) {
//
//	order, err := uc.repoOrder.ReadCurrentOrder(ctx, id)
//	if err != nil {
//		err = fmt.Errorf("error happened in repo.CreateOrder: %w", err)
//
//		return models.Order{}, err
//	}
//
//	return order, nil
//}

//func (uc *OrderUsecase) GetOrder(ctx context.Context, id uuid.UUID) (models.Order, error) {
//	cart, err := uc.repo.ReadCart(ctx, id)
//	order, err := uc.repo.ReadOrder(ctx, id)
//	if err != nil {
//		if errors.Is(err, repo.ErrOrderNotFound) {
//			// TODO: implement
//		}
//		err = fmt.Errorf("error happened in repo.ReadOrder: %w", err)
//
//		return models.Order{}, err
//	}
//
//	return order, nil
//}
