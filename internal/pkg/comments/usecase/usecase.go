package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/comments"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/comments/repo"
	uuid "github.com/satori/go.uuid"
)

type CommentsUsecase struct {
	repo comments.CommentsRepo
}

func NewCommentsUsecase(repoComments comments.CommentsRepo) *CommentsUsecase {
	return &CommentsUsecase{
		repo: repoComments,
	}
}

func (uc *CommentsUsecase) CreateComment(ctx context.Context, commentPayload models.CommentPayload) error {
	err := uc.repo.MakeComment(ctx, commentPayload)
	if err != nil {
		err = fmt.Errorf("error happened in repo.MakeComment: %w", err)

		return err
	}

	return nil
}

func (uc *CommentsUsecase) GetProductComments(ctx context.Context, productID uuid.UUID) ([]models.Comment, error) {
	comments, err := uc.repo.ReadProductComments(ctx, productID)
	if err != nil {
		if errors.Is(err, repo.ErrCommentNotFound) {
			return []models.Comment{}, err
		}
		err = fmt.Errorf("error happened in repo.ReadProductComments: %w", err)

		return []models.Comment{}, err
	}

	return comments, nil
}
