package address

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/address_mock.go -package mock

type AddressUsecase interface {
	AddAddress(context.Context, uuid.UUID, models.AddressPayload) (models.Address, error)
	UpdateAddress(context.Context, models.Address) (models.Address, error)
	DeleteAddress(context.Context, models.AddressDelete) error
	MakeCurrentAddress(context.Context, models.AddressMakeCurrent) error
	GetCurrentAddress(context.Context, uuid.UUID) (models.Address, error)
	GetAllAddresses(context.Context, uuid.UUID) ([]models.Address, error)
}

type AddressRepo interface {
	CreateAddress(context.Context, uuid.UUID, models.AddressPayload) (models.Address, error)
	UpdateAddress(context.Context, models.Address) error
	DeleteAddress(context.Context, models.AddressDelete) error
	MakeCurrentAddress(context.Context, models.AddressMakeCurrent) error
	ReadAddress(context.Context, uuid.UUID, uuid.UUID) (models.Address, error)
	ReadCurrentAddress(context.Context, uuid.UUID) (models.Address, error)
	ReadAllAddresses(context.Context, uuid.UUID) ([]models.Address, error)
	ReadCurrentAddressID(context.Context, uuid.UUID) (uuid.UUID, error)
}
