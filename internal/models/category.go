package models

type Category struct {
	Id     int64  `json:"categoryId"`
	Name   string `json:"categoryName"`
	Parent int64  `json:"categoryParent,omitempty"`
}

const MAX_LEVEL_CATEGORY = 3

type CategoryBranch [MAX_LEVEL_CATEGORY]string

type CategoryTree []Category

//func (c *Category) LogValue() slog.Value {
//	return slog.GroupValue(
//		//slog.Int64("id", c.Id),
//		slog.String("name", c.Name),
//		//slog.Int64("parent", c.Parent.String()),
//	)
//}
