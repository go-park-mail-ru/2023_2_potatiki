package usecase

import (
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order"
)

type OrderUsecase struct {
	repo order.OrderRepo
}

func NewOrderUsecase(repo order.OrderRepo) *OrderUsecase {
	return &OrderUsecase{
		repo: repo,
	}
}

//func (uc *OrderUsecase) GetOrder(ctx context.Context, id uuid.UUID) (models.Order, error) {
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
//
//func (uc *OrderUsecase) CreateOrder(ctx context.Context, id uuid.UUID) (models.Order, error) {
//	order, err := uc.repo.CreateOrder(ctx, id)
//	if err != nil {
//		err = fmt.Errorf("error happened in repo.CreateOrder: %w", err)
//
//		return models.Order{}, err
//	}
//
//	return order, nil
//}
