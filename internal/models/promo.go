package models

import "time"

//go:generate easyjson -all /home/scremyda/GolandProjects/2023_2_potatiki/internal/models/promo.go

//easyjson:json
type Promocode struct {
	Id       int64     `json:"id"`
	Discount int       `json:"discount"`
	Name     string    `json:"name"`
	Leftover int       `json:"-"`
	Deadline time.Time `json:"-"`
}
