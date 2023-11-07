package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address/repo"
	uuid "github.com/satori/go.uuid"
)

type AddressUsecase struct {
	repo address.AddressRepo
}

func NewAddressUsecase(repoAddress address.AddressRepo) *AddressUsecase {
	return &AddressUsecase{
		repo: repoAddress,
	}
}

func (uc *AddressUsecase) AddAddress(ctx context.Context, userID uuid.UUID, addressInfo models.AddressInfo) (models.Address, error) {
	addressInfo.Sanitize()

	address, err := uc.repo.CreateAddress(ctx, userID, addressInfo)
	if err != nil {
		err = fmt.Errorf("error happened in repo.CreateAddress: %w", err)

		return models.Address{}, err
	}

	return address, nil
}

func (uc *AddressUsecase) UpdateAddress(ctx context.Context, addressInfo models.Address) (models.Address, error) {
	addressInfo.Sanitize()

	err := uc.repo.UpdateAddress(ctx, addressInfo)
	if err != nil {
		err = fmt.Errorf("error happened in repo.UpdateAddress: %w", err)

		return models.Address{}, err
	}

	address, err := uc.repo.ReadAddress(ctx, addressInfo.ProfileId, addressInfo.Id)
	if err != nil {
		if errors.Is(err, repo.ErrAddressNotFound) {
			return models.Address{}, err
		}
		err = fmt.Errorf("error happened in repo.ReadAddress: %w", err)

		return models.Address{}, err
	}

	return address, nil
}

func (uc *AddressUsecase) DeleteAddress(ctx context.Context, addressInfo models.AddressDelete) error {
	err := uc.repo.DeleteAddress(ctx, addressInfo)
	if err != nil {
		if errors.Is(err, repo.ErrNoCurrentAddressNotFound) {
			return err
		}
		err = fmt.Errorf("error happened in repo.DeleteAddress: %w", err)

		return err
	}

	return nil
}

func (uc *AddressUsecase) MakeCurrentAddress(ctx context.Context, addressInfo models.AddressMakeCurrent) error {
	err := uc.repo.MakeCurrentAddress(ctx, addressInfo)
	if err != nil {
		if errors.Is(err, repo.ErrCurrentAddressNotFound) {
			return err
		}
		err = fmt.Errorf("error happened in repo.MakeCurrentAddress: %w", err)

		return err
	}

	return nil
}

func (uc *AddressUsecase) GetCurrentAddress(ctx context.Context, userID uuid.UUID) (models.Address, error) {
	address, err := uc.repo.ReadCurrentAddress(ctx, userID)
	if err != nil {
		if errors.Is(err, repo.ErrAddressNotFound) {
			return models.Address{}, err
		}
		err = fmt.Errorf("error happened in repo.ReadCurrentAddress: %w", err)

		return models.Address{}, err
	}

	return address, nil
}

func (uc *AddressUsecase) GetAllAddresses(ctx context.Context, userID uuid.UUID) ([]models.Address, error) {
	address, err := uc.repo.ReadAllAddresses(ctx, userID)
	if err != nil {
		if errors.Is(err, repo.ErrAddressesNotFound) {
			return []models.Address{}, err
		}
		err = fmt.Errorf("error happened in repo.ReadAllAddresses: %w", err)

		return []models.Address{}, err
	}

	return address, nil
}
