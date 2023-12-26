package ticker

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order"
	"time"
)

//go:generate mockgen -source hub.go -destination ./mocks/hub_mock.go -package mock

type StatusTickerInterface interface {
	Run(context.Context)
}

type StatusTicker struct {
	currentOffset time.Time
	repo          order.OrderRepo
}

func NewStatusTicker(repo order.OrderRepo) *StatusTicker {
	return &StatusTicker{
		repo:          repo,
		currentOffset: time.Now(),
	}
}

func (h *StatusTicker) Run(ctx context.Context) {
	t := time.NewTicker(30 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:

			_ = h.repo.SetOrderStatus(ctx, h.currentOffset)

			h.currentOffset = h.currentOffset.Add(30 * time.Second)

		case <-ctx.Done():
			return
		}
	}
}
