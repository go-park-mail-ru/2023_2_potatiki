package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type Order struct {
	Id           uuid.UUID `json:"id"`
	Status       string    `json:"status"`
	DeliveryDate string    `json:"deliveryDate"`
	DeliveryTime string    `json:"deliveryTime"`
	CreationAt   time.Time `json:"creationDate"`
	Address
	Products []OrderProduct `json:"products"`
}

type OrderProduct struct {
	Quantity int64 `json:"quantity"`
	Product
}

type OrderInfo struct {
	DeliveryAtDate string `json:"deliveryDate"`
	DeliveryAtTime string `json:"deliveryTime"`
}
