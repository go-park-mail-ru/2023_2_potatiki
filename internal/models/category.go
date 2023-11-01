package models

import (
	"log/slog"
)

type Category struct {
	Id     int64  `json:"id"` //nolint:stylecheck
	Name   string `json:"name"`
	Parent int64  `json:"parent"`
}

const MAX_LEVEL_CATEGORY = 3

type CategoryBranch [MAX_LEVEL_CATEGORY]string

type CategoryTree []Category

func (c *Category) LogValue() slog.Value {
	return slog.GroupValue(
		//slog.Int64("id", c.Id),
		slog.String("name", c.Name),
		//slog.Int64("parent", c.Parent.String()),
	)
}
