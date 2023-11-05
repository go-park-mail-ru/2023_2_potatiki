package models

import (
	"log/slog"

	"github.com/satori/go.uuid"
)

type Cart struct {
	Id        uuid.UUID     `json:"-"`
	ProfileId uuid.UUID     `json:"-"`
	IsCurrent bool          `json:"-"`
	Products  []CartProduct `json:"products"`
}

type CartProduct struct {
	Quantity int64 `json:"quantity"`
	Product
}

type CartUpdate struct {
	Id        uuid.UUID           `json:"-"`
	ProfileId uuid.UUID           `json:"-"`
	IsCurrent bool                `json:"-"`
	Products  []CartProductUpdate `json:"productsInfo"`
}

type CartProductUpdate struct {
	Quantity int64     `json:"quantity"`
	Id       uuid.UUID `json:"productId"`
}

type CartProductDelete struct {
	Id uuid.UUID `json:"productId"`
}

func (c *Cart) LogValue() slog.Value {
	//nolint:lll
	// check https://betterstack.com/community/guides/logging/logging-in-go/#hiding-sensitive-fields-with-the-logvaluer-interface
	return slog.GroupValue(
		slog.String("id", c.Id.String()),
		slog.String("ProfileId", c.ProfileId.String()),
	)
}
