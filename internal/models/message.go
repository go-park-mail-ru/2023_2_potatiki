package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//easyjson:skip
type Message struct {
	UserID      uuid.UUID
	Created     time.Time
	MessageInfo string
}
