package models

import uuid "github.com/satori/go.uuid"

type Address struct {
	Id        uuid.UUID `json:"addressId"`
	ProfileId uuid.UUID `json:"-"`
	City      string    `json:"city"`
	Street    string    `json:"street"`
	House     string    `json:"house"`
	Flat      string    `json:"flat"`
	IsCurrent bool      `json:"addressIsCurrent"`
}

type AddressInfo struct {
	City   string `json:"city"`
	Street string `json:"street"`
	House  string `json:"house"`
	Flat   string `json:"flat"`
}

type AddressDelete struct {
	ProfileId uuid.UUID `json:"-"`
	Id        uuid.UUID `json:"addressId"`
}

type AddressMakeCurrent struct {
	ProfileId uuid.UUID `json:"-"`
	Id        uuid.UUID `json:"addressId"`
}
