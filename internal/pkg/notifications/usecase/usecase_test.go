package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/notifications"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"testing"
)

func TestNewNotificationsUsecase(t *testing.T) {
	type args struct {
		repoNotifications notifications.NotificationsRepo
	}
	tests := []struct {
		name string
		args args
		want *NotificationsUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotificationsUsecase(tt.args.repoNotifications); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotificationsUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationsUsecase_GetDayNotifications(t *testing.T) {
	type fields struct {
		repo notifications.NotificationsRepo
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &NotificationsUsecase{
				repo: tt.fields.repo,
			}
			got, err := uc.GetDayNotifications(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDayNotifications() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDayNotifications() got = %v, want %v", got, tt.want)
			}
		})
	}
}
