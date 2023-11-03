package models

import "github.com/satori/go.uuid"

type Order struct {
	Id       uuid.UUID      `json:"id"` //nolint:stylecheck
	Status   int            `json:"statusId"`
	Products []OrderProduct `json:"products"`
}

type OrderProduct struct {
	Quantity int64 `json:"quantity"`
	Product
}
