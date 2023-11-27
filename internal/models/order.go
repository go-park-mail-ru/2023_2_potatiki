package models

import "github.com/satori/go.uuid"

type Order struct {
	Id     uuid.UUID `json:"id"`
	Status int64     `json:"statusId"`
	Address
	Products []OrderProduct `json:"products"`
}

type OrderProduct struct {
	Quantity int64 `json:"quantity"`
	Product
}
