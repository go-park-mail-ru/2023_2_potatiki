package models

import (
	"log/slog"

	"github.com/google/uuid"
)

type Cart struct {
	Id        uuid.UUID     `json:"-"` //nolint:stylecheck
	ProfileId uuid.UUID     `json:"-"` //nolint:stylecheck
	IsCurrent bool          `json:"-"` //nolint:stylecheck
	Products  []CartProduct `json:"products"`
}

type CartProduct struct {
	Quantity int64 `json:"quantity"`
	Product
}

func (c *Cart) LogValue() slog.Value {
	//nolint:lll
	// check https://betterstack.com/community/guides/logging/logging-in-go/#hiding-sensitive-fields-with-the-logvaluer-interface
	return slog.GroupValue(
		slog.String("id", c.Id.String()),
		slog.String("ProfileId", c.ProfileId.String()),
	)
}
