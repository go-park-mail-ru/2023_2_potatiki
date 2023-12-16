package http

import (
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/hub"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetNotifications(t *testing.T) {
	errVal := 00000
	testCases := []struct {
		Name     string
		Request  *http.Request
		Expected int
	}{
		{
			Name: "InvalidNotificationRequest/WrongUserID",
			Request: httptest.NewRequest("GET", "/notifications", nil).
				WithContext(context.WithValue(context.Background(), authmw.AccessTokenCookieName, errVal)),
			Expected: http.StatusUnauthorized,
		},
	}

	hub := &hub.Hub{}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			handler := NewNotificationsHandler(hub, logger.Set("local", os.Stdout))
			responseRecorder := httptest.NewRecorder()

			handler.GetNotifications(responseRecorder, tc.Request)

			assert.Equal(t, tc.Expected, responseRecorder.Code)
		})
	}
}
