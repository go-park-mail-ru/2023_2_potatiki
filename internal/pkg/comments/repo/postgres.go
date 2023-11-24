package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
)

const (
	createComment = `
	INSERT INTO comment (id, productID, userID, pros, cons, comment, rating) VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	getComments = `
	SELECT id, productID,  pros, cons, comment, rating FROM comment WHERE productID = $1;
	`
)

var (
	ErrCommentNotFound = errors.New("comment not found")
)

type CommentsRepo struct {
	db pgxtype.Querier
}

func NewCommentsRepo(db pgxtype.Querier) *CommentsRepo {
	return &CommentsRepo{
		db: db,
	}
}

func (r *CommentsRepo) MakeComment(ctx context.Context, commentPayload models.CommentPayload) error {
	commentID := uuid.NewV4()
	_, err := r.db.Exec(ctx, createComment, commentID,
		commentPayload.ProductID,
		commentPayload.UserID,
		commentPayload.Pros,
		commentPayload.Cons,
		commentPayload.Comment,
		commentPayload.Rating,
	)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return err
	}

	return nil
}

func (r *CommentsRepo) ReadProductComments(ctx context.Context, productID uuid.UUID) ([]models.Comment, error) {
	rows, err := r.db.Query(ctx, getComments, productID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []models.Comment{}, ErrCommentNotFound
		}
		err = fmt.Errorf("error happened in db.Query: %w", err)

		return []models.Comment{}, err
	}
	comment := models.Comment{}
	commentsSlice := []models.Comment{}
	for rows.Next() {
		err = rows.Scan(
			&comment.ID,
			&comment.ProductID,
			&comment.Pros,
			&comment.Cons,
			&comment.Comment,
			&comment.Rating,
		)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return []models.Comment{}, err
		}
		commentsSlice = append(commentsSlice, comment)
	}
	defer rows.Close()

	return commentsSlice, nil
}
