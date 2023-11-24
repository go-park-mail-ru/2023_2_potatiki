package comments

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/comments_mock.go -package mock

type CommentsUsecase interface {
	CreateComment(context.Context, models.CommentPayload) error
	GetProductComments(context.Context, uuid.UUID) ([]models.Comment, error)
}

type CommentsRepo interface {
	MakeComment(context.Context, models.CommentPayload) error
	ReadProductComments(context.Context, uuid.UUID) ([]models.Comment, error)
}
