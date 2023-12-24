package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	uuid "github.com/satori/go.uuid"
)

const (
	getNotifications = `
	SELECT user_id, created, message_info
	FROM messages
	WHERE user_id = $1;
	`
)

var (
	ErrNotificationsNotFound = errors.New("notifications not found")
)

type NotificationsRepo struct {
	db pgxtype.Querier
}

func NewNotificationsRepo(db pgxtype.Querier) *NotificationsRepo {
	return &NotificationsRepo{
		db: db,
	}
}

func (r *NotificationsRepo) ReadDayNotifications(ctx context.Context, userID uuid.UUID) ([]models.Message, error) {
	rows, err := r.db.Query(ctx, getNotifications, userID)
	if err != nil {
		err = fmt.Errorf("error happened in db.Query: %w", err)

		return []models.Message{}, err
	}
	message := models.Message{}
	messagesSlice := []models.Message{}
	for rows.Next() {
		err = rows.Scan(
			&message.UserID,
			&message.Created,
			&message.MessageInfo,
		)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return []models.Message{}, err
		}
		messagesSlice = append(messagesSlice, message)
	}
	defer rows.Close()

	if len(messagesSlice) == 0 {
		return []models.Message{}, ErrNotificationsNotFound
	}

	return messagesSlice, nil
}
