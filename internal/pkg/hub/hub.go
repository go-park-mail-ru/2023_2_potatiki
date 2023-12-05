package hub

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type Hub struct {
	connect       map[*websocket.Conn]uuid.UUID
	currentOffset time.Time
	repo          order.OrderRepo
	m             sync.RWMutex
}

func NewHub(repo order.OrderRepo) *Hub {
	return &Hub{
		repo:          repo,
		currentOffset: time.Now(),
		connect:       make(map[*websocket.Conn]uuid.UUID),
		m:             sync.RWMutex{},
	}
}

func (h *Hub) AddClient(userID uuid.UUID, client *websocket.Conn) {

	h.connect[client] = userID

	go func() {
		for {
			_, _, err := client.NextReader()
			if err != nil {
				_ = client.Close()
				return
			}
		}
	}()

	client.SetCloseHandler(func(code int, text string) error {
		h.m.Lock()
		delete(h.connect, client) //TODO: add sync.Map
		h.m.Unlock()
		return nil
	})

}

func (h *Hub) Run(ctx context.Context) {
	t := time.NewTicker(30 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			for connect, userID := range h.connect {
				messages, _ := h.repo.GetUpdates(ctx, userID, h.currentOffset)
				for _, message := range messages {
					err := connect.WriteJSON(message)
					if err != nil {
						continue
					}
				}
			}
			h.currentOffset = h.currentOffset.Add(30 * time.Second)
		case <-ctx.Done():
			return
		}

	}
}
