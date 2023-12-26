package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//go:generate easyjson -all /home/scremyda/GolandProjects/2023_2_potatiki/internal/models/order.go

//easyjson:json
type Order struct {
	Id           uuid.UUID `json:"id"`
	Status       string    `json:"status"`
	StatusId     int64     `json:"_"`
	DeliveryDate string    `json:"deliveryDate"`
	DeliveryTime string    `json:"deliveryTime"`
	PomocodeName string    `json:"promocodeName"`
	CreationAt   time.Time `json:"creationDate"`
	Address      `json:"address"`
	Products     []OrderProduct `json:"products"`
}

//easyjson:json
type OrderSlice []Order

//easyjson:json
type OrderProduct struct {
	Quantity int64 `json:"quantity"`
	Product
}

//easyjson:json
type OrderInfo struct {
	DeliveryAtDate string `json:"deliveryDate"`
	DeliveryAtTime string `json:"deliveryTime"`
	PromocodeName  string `json:"promocodeName,omitempty"`
}
