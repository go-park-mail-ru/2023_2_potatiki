package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Message struct {
	UserID      uuid.UUID
	Created     time.Time
	MessageInfo string
}
