package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//easyjson:json
type Order struct {
	Id           uuid.UUID `json:"id"`
	Status       string    `json:"status"`
	DeliveryDate string    `json:"deliveryDate"`
	DeliveryTime string    `json:"deliveryTime"`
	PomocodeName string    `json:"promocodeName"`
	CreationAt   time.Time `json:"creationDate"`
	Address
	Products []OrderProduct `json:"products"`
}

//easyjson:json
type OrderSlice []Order

//easyjson:json
type OrderProduct struct {
	Quantity int64 `json:"quantity"`
	Product        //`json:"product"`
}

//easyjson:json
type OrderInfo struct {
	DeliveryAtDate string `json:"deliveryDate"`
	DeliveryAtTime string `json:"deliveryTime"`
	PromocodeName  string `json:"promocodeName,omitempty"`
}
