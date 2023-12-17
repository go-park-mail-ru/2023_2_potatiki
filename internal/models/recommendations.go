package models

import uuid "github.com/satori/go.uuid"

const (
	MinProductsCount          = 20
	MinCateggoriesCount       = 4
	ProductCountFromCategory  = 2
	ProductCountFromStatistic = 10
)

type UserActivity struct {
	Product  ProductStatisticSlice  `json:"product"`
	Category CategoryStatisticSlice `json:"category"`
}

//easyjson:json
type ProductStatisticSlice []ProductStatistic

//easyjson:json
type CategoryStatisticSlice []CategoryStatistic

type UserActivityStore struct {
	Product  ProductStatisticMap  `json:"product"`
	Category CategoryStatisticMap `json:"category"`
}

//easyjson:json
type ProductStatisticMap map[uuid.UUID]ProductStatistic

//easyjson:json
type CategoryStatisticMap map[int64]CategoryStatistic

type ProductStatistic struct {
	ProductID      uuid.UUID `json:"productId"`
	ActivityPoints int64     `json:"activityPoints"`
	IsBought       bool      `json:"isBought"`
	IsReviewed     bool      `json:"isReviewed"`
}

type CategoryStatistic struct {
	CategoryID     int64 `json:"categoryId"`
	ActivityPoints int64 `json:"activityPoints"`
}
