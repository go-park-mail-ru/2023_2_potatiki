package models

//easyjson:json
type Promocode struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Discount int    `json:"discount"`
}
