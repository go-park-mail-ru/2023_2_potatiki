package models

import uuid "github.com/satori/go.uuid"

//go:generate easyjson -all /home/scremyda/GolandProjects/2023_2_potatiki/internal/models/recommendations.go

const (
	MinProductsCount          = 20
	MinCateggoriesCount       = 4
	ProductCountFromCategory  = 2
	ProductCountFromStatistic = 10
)

//easyjson:json
type UserActivity struct {
	Product  ProductStatisticSlice  `json:"product"`
	Category CategoryStatisticSlice `json:"category"`
}

//easyjson:json
type ProductStatisticSlice []ProductStatistic

//easyjson:json
type CategoryStatisticSlice []CategoryStatistic

//easyjson:json
type UserActivityStore struct {
	Product  ProductStatisticMap  `json:"product"`
	Category CategoryStatisticMap `json:"category"`
}

//easyjson:json
type ProductStatisticMap map[uuid.UUID]ProductStatistic

//easyjson:json
type CategoryStatisticMap map[int64]CategoryStatistic

//easyjson:json
type ProductStatistic struct {
	ProductID      uuid.UUID `json:"productId"`
	ActivityPoints int64     `json:"activityPoints"`
	IsBought       bool      `json:"isBought"`
	IsReviewed     bool      `json:"isReviewed"`
}

//easyjson:json
type CategoryStatistic struct {
	CategoryID     int64 `json:"categoryId"`
	ActivityPoints int64 `json:"activityPoints"`
}
