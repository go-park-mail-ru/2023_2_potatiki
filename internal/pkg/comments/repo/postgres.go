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
	SELECT c.id, c.productID, c.pros, c.cons, c.comment, c.rating, p.login, c.creation_at
	FROM comment c
	JOIN profile p ON c.userID = p.id
	WHERE c.productID = $1;
	`

	getComment = `
	SELECT c.id, c.productID, c.pros, c.cons, c.comment, c.rating, p.login, c.creation_at
	FROM comment c
	JOIN profile p ON c.userID = p.id
	WHERE c.userID = $1
	  AND c.productID = $2;
	`

	countOfCommentsToProduct = `
	SELECT COUNT(*), c.id, c.productID, c.pros, c.cons, c.comment, c.rating, p.login, c.creation_at
	FROM comment c
	JOIN profile p ON c.userID = p.id
	WHERE c.userID = $1 AND c.productID = $2
	GROUP BY c.id, c.productID, c.pros, c.cons, c.comment, c.rating, p.login, c.creation_at;
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

func (r *CommentsRepo) ReadCountOfCommentsToProduct(ctx context.Context, userID, productID uuid.UUID) (int, models.Comment, error) {
	countComments := 0
	comment := models.Comment{}
	err := r.db.QueryRow(ctx, countOfCommentsToProduct, userID, productID).
		Scan(&countComments, &comment.ID, &comment.ProductID, &comment.Pros, &comment.Cons,
			&comment.Comment, &comment.Rating, &comment.UserName, &comment.CreationDate)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return countComments, models.Comment{}, ErrCommentNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return countComments, models.Comment{}, err
	}

	return countComments, comment, nil
}

func (r *CommentsRepo) MakeComment(ctx context.Context, commentPayload models.CommentPayload) (models.Comment, error) {
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

		return models.Comment{}, err
	}

	comment := models.Comment{}
	err = r.db.QueryRow(ctx, getComment, commentPayload.UserID, commentPayload.ProductID).
		Scan(&comment.ID, &comment.ProductID, &comment.Pros, &comment.Cons,
			&comment.Comment, &comment.Rating, &comment.UserName, &comment.CreationDate)
	if err != nil {
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return models.Comment{}, err
	}
	return comment, nil
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
			&comment.UserName,
			&comment.CreationDate,
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
