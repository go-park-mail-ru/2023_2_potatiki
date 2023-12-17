package models

import (
	uuid "github.com/satori/go.uuid"
)

//easyjson:json
type Cart struct {
	Id        uuid.UUID     `json:"-"`
	ProfileId uuid.UUID     `json:"-"`
	IsCurrent bool          `json:"-"`
	Products  []CartProduct `json:"products"`
}

//easyjson:json
type CartProduct struct {
	Quantity int64 `json:"quantity"`
	Product
}

//easyjson:json
type CartUpdate struct {
	Id        uuid.UUID           `json:"-"`
	ProfileId uuid.UUID           `json:"-"`
	IsCurrent bool                `json:"-"`
	Products  []CartProductUpdate `json:"productsInfo"`
}

//easyjson:json
type CartProductUpdate struct {
	Quantity int64     `json:"quantity"`
	Id       uuid.UUID `json:"productId"`
}

//easyjson:json
type CartProductDelete struct {
	Id uuid.UUID `json:"productId"`
}

//func (c *Cart) LogValue() slog.Value {
//	return slog.GroupValue(
//		slog.String("id", c.Id.String()),
//		slog.String("ProfileId", c.ProfileId.String()),
//	)
//}
