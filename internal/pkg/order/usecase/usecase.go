package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address"
	addressRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address/repo"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/promo"

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
	repoPromo   promo.PromoRepo
}

func NewOrderUsecase(
	repoOrder order.OrderRepo,
	repoCart cart.CartRepo,
	repoAddress address.AddressRepo,
	repoPromo promo.PromoRepo) *OrderUsecase {
	return &OrderUsecase{
		repoOrder:   repoOrder,
		repoCart:    repoCart,
		repoAddress: repoAddress,
		repoPromo:   repoPromo,
	}
}

func (uc *OrderUsecase) CreateOrder(
	ctx context.Context,
	userID uuid.UUID,
	deliveryTime,
	deliveryDate,
	promocodeName string) (models.Order, error) {
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

	promocodeID := int64(-1)
	if promocodeName != "" {
		promocode, err := uc.repoPromo.UsePromocode(ctx, promocodeName)
		if err != nil {
			return models.Order{}, err
		}
		if time.Now().After(promocode.Deadline) {
			return models.Order{}, promo.ErrPromocodeExpired
		}
		if promocode.Leftover < 1 {
			return models.Order{}, promo.ErrPromocodeLeftout
		}
		promocodeID = promocode.Id
	}

	order, err := uc.repoOrder.CreateOrder(ctx, cart, address.Id, userID, 1, promocodeID, deliveryTime, deliveryDate)
	if err != nil {
		if errors.Is(err, orderRepo.ErrPoductNotFound) {
			return models.Order{}, err
		}
		err = fmt.Errorf("error happened in repo.CreateOrder: %w", err)

		return models.Order{}, err
	}
	order.Address = address
	order.PomocodeName = promocodeName
	order.DeliveryDate = deliveryDate
	order.DeliveryTime = deliveryTime

	err = uc.repoCart.DeleteCart(ctx, cart.Id)
	if err != nil {
		err = fmt.Errorf("error happened in repoCart.DeleteCart: %w", err)

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
		return nil, fmt.Errorf("error happened in repoOrder.ReadOrder: %w", err)
	}

	orders := make([]models.Order, len(ordersID))
	for i, orderID := range ordersID {
		orders[i], err = uc.repoOrder.ReadOrder(ctx, orderID)
		if err != nil {
			if errors.Is(err, orderRepo.ErrPoductsInOrderNotFound) {
				return []models.Order{}, err
			}
			return nil, fmt.Errorf("error happened in repoOrder.ReadOrder: %w", err)
		}
	}

	return orders, nil
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
