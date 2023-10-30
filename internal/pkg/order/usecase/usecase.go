package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart"
	repoCart "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/repo"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/repo"
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
		if errors.Is(err, repoCart.ErrCartNotFound) {
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
		if errors.Is(err, repo.ErrPoductNotFound) {
			return models.Order{}, err
		}
		err = fmt.Errorf("error happened in repo.CreateOrder: %w", err)

		return models.Order{}, err
	}

	return order, nil
}

func (uc *OrderUsecase) GetCurrentOrder(ctx context.Context, userID uuid.UUID) (models.Order, error) {

	orderID, err := uc.repoOrder.ReadOrderID(ctx, userID, "Pending")
	if err != nil {
		if errors.Is(err, repo.ErrOrderNotFound) {
			// TODO: implement
		}
		err = fmt.Errorf("error happened in repoOrder.ReadOrderID: %w", err)

		return models.Order{}, err
	}

	order, err := uc.repoOrder.ReadOrder(ctx, orderID)
	if err != nil {
		if errors.Is(err, repo.ErrPoductsInOrderNotFound) {
			// TODO: implement
		}
		err = fmt.Errorf("error happened in repoOrder.ReadOrder: %w", err)

		return models.Order{}, err
	}

	return order, nil
}

func (uc *OrderUsecase) GetOrders(ctx context.Context, userID uuid.UUID) ([]models.Order, error) {
	ordersID, err := uc.repoOrder.ReadOrdersID(ctx, userID)
	if err != nil {
		if errors.Is(err, repo.ErrOrdersNotFound) {
			// TODO: implement
		}
		err = fmt.Errorf("error happened in repoOrder.ReadOrder: %w", err)

		return []models.Order{}, err
	}

	var orders []models.Order
	for _, orderID := range ordersID {
		order, err := uc.repoOrder.ReadOrder(ctx, orderID)
		if err != nil {
			if errors.Is(err, repo.ErrPoductsInOrderNotFound) {
				// TODO: implement
			}
			err = fmt.Errorf("error happened in repoOrder.ReadOrder: %w", err)

			return []models.Order{}, err // ASK: Стоит ли выбрасывать пустой слайс при ошибке чтения?
		}
		orders = append(orders, order)
	}

	return orders, nil
}
