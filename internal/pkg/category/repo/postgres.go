package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/jackc/pgtype/pgxtype"
)

const (
	getCategories = "SELECT * FROM category"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type CategoryRepo struct {
	db pgxtype.Querier // TODO: add logger
}

func NewCategoryRepo(db pgxtype.Querier) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (r *CategoryRepo) ReadCategories(ctx context.Context) ([]models.Category, error) {
	var categorySlice []models.Category
	rows, err := r.db.Query(ctx, getCategories)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Category{}, ErrCategoryNotFound
		}
		err = fmt.Errorf("error happened in db.QueryContext: %w", err)

		return []models.Category{}, err
	}
	category := models.Category{}
	for rows.Next() {
		err = rows.Scan(&category.Id, &category.Name, &category.Parent)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return []models.Category{}, err
		}
		categorySlice = append(categorySlice, category)
	}
	defer rows.Close()

	return categorySlice, nil
}
