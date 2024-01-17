package repo

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"testing"
)

func TestNewNotificationsRepo(t *testing.T) {
	type args struct {
		db pgxtype.Querier
	}
	tests := []struct {
		name string
		args args
		want *NotificationsRepo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotificationsRepo(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotificationsRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationsRepo_ReadDayNotifications(t *testing.T) {
	type fields struct {
		db pgxtype.Querier
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
			r := &NotificationsRepo{
				db: tt.fields.db,
			}
			got, err := r.ReadDayNotifications(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDayNotifications() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDayNotifications() got = %v, want %v", got, tt.want)
			}
		})
	}
}
