package models

import "github.com/google/uuid"

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
