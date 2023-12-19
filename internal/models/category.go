package models

//go:generate easyjson -all /home/scremyda/GolandProjects/2023_2_potatiki/internal/models/category.go

//easyjson:json
type Category struct {
	Id     int64  `json:"categoryId"`
	Name   string `json:"categoryName"`
	Parent int64  `json:"categoryParent,omitempty"`
}

const MAX_LEVEL_CATEGORY = 3

//easyjson:json
type CategoryBranch [MAX_LEVEL_CATEGORY]string

//easyjson:json
type CategoryTree []Category

//func (c *Category) LogValue() slog.Value {
//	return slog.GroupValue(
//		//slog.Int64("id", c.Id),
//		slog.String("name", c.Name),
//		//slog.Int64("parent", c.Parent.String()),
//	)
//}
