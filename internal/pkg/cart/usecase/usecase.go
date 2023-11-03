package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/repo"
	"github.com/satori/go.uuid"
)

type CartUsecase struct {
	repo cart.CartRepo
}

func NewCartUsecase(repo cart.CartRepo) *CartUsecase {
	return &CartUsecase{
		repo: repo,
	}
}

func (uc *CartUsecase) GetCart(ctx context.Context, id uuid.UUID) (models.Cart, error) {
	cart, err := uc.repo.ReadCart(ctx, id)

	if errors.Is(err, repo.ErrCartNotFound) {
		cart.Id, err = uc.repo.CreateCart(ctx, id)
		cart.IsCurrent = true
		if err != nil {
			err = fmt.Errorf("error happened in repo.GetCart: %w", err)

			return models.Cart{}, err
		}

		return cart, nil
	}

	if err != nil {
		return models.Cart{}, err
	}

	return cart, nil
}

func (uc *CartUsecase) AddProduct(ctx context.Context, cart models.Cart, product models.CartProductUpdate) (models.Cart, error) {
	userID := cart.ProfileId
	cartChecked, err := uc.repo.CheckCart(ctx, userID)

	if errors.Is(err, repo.ErrCartNotFound) {
		cart.Id, err = uc.repo.CreateCart(ctx, userID)
		cart.IsCurrent = true
		if err != nil {
			err = fmt.Errorf("error happened in repo.GetCart: %w", err)

			return models.Cart{}, err
		}

		cart, err = uc.repo.AddProduct(ctx, cart, product)

		return cart, err
	}

	if err != nil {
		return models.Cart{}, err
	}

	cart, err = uc.repo.AddProduct(ctx, cartChecked, product)

	return cart, err
}

func (uc *CartUsecase) DeleteProduct(ctx context.Context, cart models.Cart, product models.CartProductDelete) (models.Cart, error) {
	cart, err := uc.repo.CheckCart(ctx, cart.ProfileId)
	if err != nil {
		return models.Cart{}, err
	}

	cart, err = uc.repo.DeleteProduct(ctx, cart, product)

	return cart, err
}

func (uc *CartUsecase) UpdateCart(ctx context.Context, cart models.Cart) (models.Cart, error) {
	_, err := uc.repo.ReadCart(ctx, cart.ProfileId)
	if errors.Is(err, repo.ErrCartNotFound) {
		cart.Id, err = uc.repo.CreateCart(ctx, cart.ProfileId)
		cart.IsCurrent = true
		if err != nil {
			return models.Cart{}, err
		}
	}

	if err != nil {
		return models.Cart{}, err
	}

	cart, err = uc.repo.UpdateCart(ctx, cart)
	if err != nil {
		err = fmt.Errorf("error happened in repo.UpdateCart: %w", err)

		return models.Cart{}, err
	}

	return cart, nil
}
