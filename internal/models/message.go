package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//go:generate easyjson /home/scremyda/GolandProjects/2023_2_potatiki/internal/models/message.go

//easyjson:skip
type Message struct {
	UserID      uuid.UUID
	Created     time.Time
	MessageInfo string
	Type        string
	OrderID     uuid.UUID
}

//easyjson:json
type MessageSlice []Message
