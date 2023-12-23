package hub

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

//go:generate mockgen -source hub.go -destination ./mocks/hub_mock.go -package mock

type HubInterface interface {
	AddClient(uuid.UUID, *websocket.Conn)
	Run(context.Context)
}

type Hub struct {
	connect       sync.Map
	currentOffset time.Time
	repo          order.OrderRepo
}

func NewHub(repo order.OrderRepo) *Hub {
	return &Hub{
		repo:          repo,
		currentOffset: time.Now(),
	}
}

func (h *Hub) AddClient(userID uuid.UUID, client *websocket.Conn) {
	h.connect.Store(client, userID)

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
		h.connect.Delete(client)
		return nil
	})
}

func (h *Hub) Run(ctx context.Context) {
	t := time.NewTicker(8 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			h.connect.Range(func(key, value interface{}) bool {
				connect := key.(*websocket.Conn)
				userID := value.(uuid.UUID)

				messages, _ := h.repo.GetUpdates(ctx, userID, h.currentOffset)
				for _, message := range messages {
					err := connect.WriteJSON(message)
					if err != nil {
						continue
					}
				}

				return true
			})

			h.currentOffset = h.currentOffset.Add(8 * time.Second)

		case <-ctx.Done():
			return
		}
	}
}
