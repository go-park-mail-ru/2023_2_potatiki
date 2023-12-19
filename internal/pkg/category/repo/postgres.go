package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/jackc/pgtype/pgxtype"
)

const (
	getCategories = "SELECT id, name, COALESCE(parent, -1) FROM category ORDER BY id"
	/*getCategories = `SELECT t1.name AS lvl1, t2.name AS lvl2, t3.name AS lvl3
	FROM category AS t1
	LEFT OUTER JOIN category AS t2 ON t2.parent = t1.id
	LEFT OUTER JOIN category AS t3 ON t3.parent = t2.id`*/
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

func (r *CategoryRepo) ReadCategories(ctx context.Context) (models.CategoryTree, error) {
	categorySlice := make(models.CategoryTree, 0)
	rows, err := r.db.Query(ctx, getCategories)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
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
