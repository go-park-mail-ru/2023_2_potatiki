package models

import (
	"log/slog"

	"github.com/google/uuid"
)

type Category struct {
	Id     uuid.UUID `json:"id"` //nolint:stylecheck
	Name   string    `json:"name"`
	Parent uuid.UUID `json:"parent"`
}

func (c *Category) LogValue() slog.Value {
	//nolint:lll
	// check https://betterstack.com/community/guides/logging/logging-in-go/#hiding-sensitive-fields-with-the-logvaluer-interface
	return slog.GroupValue(
		slog.String("id", c.Id.String()),
		slog.String("name", c.Name),
		slog.String("parent", c.Parent.String()),
	)
}
