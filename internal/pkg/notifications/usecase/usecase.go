package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/notifications"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/notifications/repo"
	uuid "github.com/satori/go.uuid"
)

type NotificationsUsecase struct {
	repo notifications.NotificationsRepo
}

func NewNotificationsUsecase(repoNotifications notifications.NotificationsRepo) *NotificationsUsecase {
	return &NotificationsUsecase{
		repo: repoNotifications,
	}
}

func (uc *NotificationsUsecase) GetDayNotifications(ctx context.Context, userID uuid.UUID) ([]models.Message, error) {
	notifications, err := uc.repo.ReadDayNotifications(ctx, userID)
	if err != nil {
		if errors.Is(err, repo.ErrNotificationsNotFound) {
			return []models.Message{}, err
		}
		err = fmt.Errorf("error happened in repo.ReadDayNotifications: %w", err)

		return []models.Message{}, err
	}

	return notifications, nil
}
