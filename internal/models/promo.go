package models

import "time"

//easyjson:json
type Promocode struct {
	Id       int64     `json:"id"`
	Discount int       `json:"discount"`
	Name     string    `json:"name"`
	Leftover int       `json:"-"`
	Deadline time.Time `json:"-"`
}
