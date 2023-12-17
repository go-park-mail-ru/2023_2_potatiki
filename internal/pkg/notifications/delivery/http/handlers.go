package http

import (
	userConnections "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/hub"
	"github.com/gorilla/websocket"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"
	uuid "github.com/satori/go.uuid"
)

var (
	upgrader = websocket.Upgrader{}
)

type NotificationsHandler struct {
	hub userConnections.HubInterface
	log *slog.Logger
}

func NewNotificationsHandler(hub userConnections.HubInterface, log *slog.Logger) *NotificationsHandler {
	return &NotificationsHandler{
		log: log,
		hub: hub,
	}
}

func (h *NotificationsHandler) GetNotifications(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)

	userID, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
	if !ok {
		h.log.Error("failed cast uuid from context value")
		resp.JSONStatus(w, http.StatusUnauthorized)

		return
	}

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.log.Error("upgrade error:", err)
		return
	}
	h.log.Debug("connection upgraded: ", slog.Any("userID", userID))

	h.hub.AddClient(userID, connection)

	h.log.Debug("client disconnected: ", slog.Any("userID", userID))
}
