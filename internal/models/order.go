package models

import "github.com/google/uuid"

type Order struct {
	Id        uuid.UUID     `json:"id"`        //nolint:stylecheck
	ProfileId uuid.UUID     `json:"profileId"` //nolint:stylecheck
	IsCurrent bool          `json:"isCurrent"` //nolint:stylecheck
	Products  []CartProduct `json:"products"`
}

type CartProduct struct {
	Quantity int64 `json:"quantity"`
	Product
}
