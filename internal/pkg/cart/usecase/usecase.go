package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/repo"
	"github.com/google/uuid"
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
		cart, err = uc.repo.CreateCart(ctx, cart.ProfileId)
		if err != nil {
			err = fmt.Errorf("error happened in repo.GetCart: %w", err)

			return models.Cart{}, err
		}

		return cart, nil
	}

	return cart, nil
}

func (uc *CartUsecase) AddProduct(ctx context.Context, cart models.Cart, product models.CartProduct) (models.Cart, error) {
	c, err := uc.repo.CheckCart(ctx, cart.ProfileId)

	if errors.Is(err, repo.ErrCartNotFound) {
		c, err = uc.repo.CreateCart(ctx, cart.ProfileId)
		if err != nil {
			err = fmt.Errorf("error happened in repo.GetCart: %w", err)

			return models.Cart{}, err
		}

		c, err = uc.repo.AddProduct(ctx, c, product)

		return c, err
	}

	c, err = uc.repo.AddProduct(ctx, c, product)

	return cart, err
}

func (uc *CartUsecase) DeleteProduct(ctx context.Context, cart models.Cart, product models.CartProduct) (models.Cart, error) {
	c, err := uc.repo.CheckCart(ctx, cart.ProfileId)
	if err != nil {
		return models.Cart{}, err
	}

	c, err = uc.repo.DeleteProduct(ctx, c, product)

	return cart, err
}

func (uc *CartUsecase) UpdateCart(ctx context.Context, cart models.Cart) (models.Cart, error) {
	_, err := uc.repo.ReadCart(ctx, cart.ProfileId)
	if errors.Is(err, repo.ErrCartNotFound) {
		cart, err = uc.repo.CreateCart(ctx, cart.ProfileId)
		if err != nil {
			return models.Cart{}, err
		}
	}

	cart, err = uc.repo.UpdateCart(ctx, cart)
	if err != nil {
		err = fmt.Errorf("error happened in repo.UpdateCart: %w", err)

		return models.Cart{}, err
	}

	return cart, nil
}
