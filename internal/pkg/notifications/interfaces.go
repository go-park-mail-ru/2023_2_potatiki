package notifications

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/notifications_mock.go -package mock

type NotificationsUsecase interface {
	GetDayNotifications(context.Context, uuid.UUID) ([]models.Message, error)
}

type NotificationsRepo interface {
	ReadDayNotifications(context.Context, uuid.UUID) ([]models.Message, error)
}
