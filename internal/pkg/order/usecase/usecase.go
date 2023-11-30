package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address"
	addressRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address/repo"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart"
	cartRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/repo"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order"
	orderRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/repo"
	uuid "github.com/satori/go.uuid"
)

type OrderUsecase struct {
	repoOrder   order.OrderRepo
	repoCart    cart.CartRepo
	repoAddress address.AddressRepo
}

func NewOrderUsecase(repoOrder order.OrderRepo, repoCart cart.CartRepo, repoAddress address.AddressRepo) *OrderUsecase {
	return &OrderUsecase{
		repoOrder:   repoOrder,
		repoCart:    repoCart,
		repoAddress: repoAddress,
	}
}

func (uc *OrderUsecase) CreateOrder(ctx context.Context, userID uuid.UUID, deliveryTime, deliveryDate string) (models.Order, error) {
	address, err := uc.repoAddress.ReadCurrentAddress(ctx, userID)
	if err != nil {
		if errors.Is(err, addressRepo.ErrAddressNotFound) {
			return models.Order{}, err
		}
		err = fmt.Errorf("error happened in repoAddress.ReadCurrentAddressID: %w", err)

		return models.Order{}, err
	}

	cart, err := uc.repoCart.ReadCart(ctx, userID)
	if err != nil {
		if errors.Is(err, cartRepo.ErrCartNotFound) {
			return models.Order{}, err
		}
		err = fmt.Errorf("error happened in repoCart.ReadCart: %w", err)

		return models.Order{}, err
	}

	order, err := uc.repoOrder.CreateOrder(ctx, cart, address.Id, userID, 1, deliveryTime, deliveryDate)
	if err != nil {
		if errors.Is(err, orderRepo.ErrPoductNotFound) {
			return models.Order{}, err
		}
		err = fmt.Errorf("error happened in repo.CreateOrder: %w", err)

		return models.Order{}, err
	}
	order.Address = address

	err = uc.repoCart.DeleteCart(ctx, cart.Id)
	if err != nil {
		err = fmt.Errorf("error happened in repoCart.DeleteCart: %w", err)

		return models.Order{}, err
	}

	return order, nil
}

func (uc *OrderUsecase) GetCurrentOrder(ctx context.Context, userID uuid.UUID) (models.Order, error) {
	orderID, err := uc.repoOrder.ReadOrderID(ctx, userID)
	if err != nil {
		if errors.Is(err, orderRepo.ErrOrderNotFound) {
			return models.Order{}, err
		}
		err = fmt.Errorf("error happened in repoOrder.ReadOrderID: %w", err)

		return models.Order{}, err
	}

	order, err := uc.repoOrder.ReadOrder(ctx, orderID)
	if err != nil {
		if errors.Is(err, orderRepo.ErrPoductsInOrderNotFound) {
			return models.Order{}, err
		}
		err = fmt.Errorf("error happened in repoOrder.ReadOrder: %w", err)

		return models.Order{}, err
	}

	return order, nil
}

func (uc *OrderUsecase) GetOrders(ctx context.Context, userID uuid.UUID) ([]models.Order, error) {
	ordersID, err := uc.repoOrder.ReadOrdersID(ctx, userID)
	if err != nil {
		if errors.Is(err, orderRepo.ErrOrdersNotFound) {
			return []models.Order{}, err
		}
		err = fmt.Errorf("error happened in repoOrder.ReadOrder: %w", err)

		return []models.Order{}, err
	}

	orders := make([]models.Order, 0)
	for _, orderID := range ordersID {
		order, err := uc.repoOrder.ReadOrder(ctx, orderID)
		if err != nil {
			if errors.Is(err, orderRepo.ErrPoductsInOrderNotFound) {
				return []models.Order{}, err
			}
			err = fmt.Errorf("error happened in repoOrder.ReadOrder: %w", err)

			return []models.Order{}, err // ASK: Стоит ли выбрасывать пустой слайс при ошибке чтения?
		}
		orders = append(orders, order)
	}

	return orders, nil
}
