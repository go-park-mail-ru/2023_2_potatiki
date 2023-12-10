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
	getActivity                   = "SELECT a.body from activities a WHERE a.user_id = $1"
	createActivity                = "INSERT INTO activities(user_id) VALUES ($1)"
	updateUserActivity            = "UPDATE activities SET body = $2 WHERE user_id = $1"
	getProductsFromParentCategory = `select p.id, p.name, p.description,     
		 p.price, p.imgsrc, COALESCE(AVG(cm.rating), 0) AS average_rating , p.category_id, 
		 ct.name as category_name, p.count_comments  
		 from product p     
		 LEFT JOIN comment cm ON p.id = cm.productID    
		 join category ct on ct.id = p.category_id     
		 where ct.id in (select c.id as id     
			 from category c     
			 where c.id <> $1 and c.parent in 
				  (select c1.id as id    
				 from category c1     
				 join category c2 on c1.id = c2.parent     
				 where c2.id = $1)
			 )  
		 GROUP BY
		 p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, ct.name, p.count_comments
		 ORDER BY
			COALESCE(AVG(cm.rating), 0)    
		 limit 20; `

	getProduct = `SELECT p.id, p.name, p.description, p.price, p.imgsrc, COALESCE(AVG(cm.rating), 0) AS average_rating,
       p.category_id, c.name AS category_name, p.count_comments
	FROM product p
	JOIN category c ON p.category_id = c.id
	LEFT JOIN comment cm ON p.id = cm.productID
	WHERE p.id = $1
	GROUP BY p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name, p.count_comments;
	`

	getProductsFromCategory = `
		SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.description,
		p.price,
		p.imgsrc,
		COALESCE(AVG(cm.rating), 0) AS average_rating,
		p.category_id,
		c.name AS category_name,
		p.count_comments
	FROM
		product p
	JOIN
		category c ON p.category_id = c.id
	LEFT JOIN
		comment cm ON p.id = cm.productID
	WHERE p.category_id = $1
	GROUP BY
		p.id, p.name, p.description, p.price, p.imgsrc, p.category_id, c.name, p.count_comments
	ORDER BY
   		COALESCE(AVG(cm.rating), 0)
	LIMIT
		$2;
	`
)

var (
	ErrActivityNotFound = errors.New("activity not found")
	ErrProductsNotFound = errors.New("products not found")
	ErrProductNotFound  = errors.New("product not found")
)

type RecommendationsRepo struct {
	db pgxtype.Querier
}

func NewRecommendationsRepo(db pgxtype.Querier) *RecommendationsRepo {
	return &RecommendationsRepo{
		db: db,
	}
}

func (r *RecommendationsRepo) ReadProductsFromCategory(ctx context.Context, categoryID int64) (models.ProductSlice, error) {
	productSlice := make(models.ProductSlice, 0, models.ProductCountFromCategory)
	rows, err := r.db.Query(ctx, getProductsFromCategory, categoryID, models.ProductCountFromCategory)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []models.Product{}, ErrProductsNotFound
		}
		err = fmt.Errorf("error happened in db.QueryContext: %w", err)

		return []models.Product{}, err
	}
	product := models.Product{}
	for rows.Next() {
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ImgSrc,
			&product.Rating,
			&product.Category.Id,
			&product.Category.Name,
			&product.CountComments,
		)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return []models.Product{}, err
		}
		productSlice = append(productSlice, product)
	}
	defer rows.Close()

	return productSlice, nil
}

func (r *RecommendationsRepo) ReadProductsFromCategories(ctx context.Context, categoryIDs models.CategoryIDs) (models.ProductSlice, error) {
	productSlice := make(models.ProductSlice, 0, models.MinProductsCount/2)
	for _, id := range categoryIDs {
		products, err := r.ReadProductsFromCategory(ctx, id)
		if err != nil {
			if !errors.Is(err, ErrProductNotFound) {
				err = fmt.Errorf("error happened in db.QueryContext: %w", err)

				return models.ProductSlice{}, err
			}
			break
		}
		productSlice = append(productSlice, products...)
	}

	return productSlice, nil
}

func (r *RecommendationsRepo) ReadRecommendations(ctx context.Context, productIDs models.ProductIDs) (models.ProductSlice, error) {
	productSlice := make(models.ProductSlice, 0, models.MinProductsCount/2)
	for _, id := range productIDs {
		product, err := r.ReadProduct(ctx, id)
		if err != nil {
			if !errors.Is(err, ErrProductNotFound) {
				err = fmt.Errorf("error happened in db.QueryContext: %w", err)

				return models.ProductSlice{}, err
			}
			break
		}
		productSlice = append(productSlice, product)
	}

	return productSlice, nil

}

func (r *RecommendationsRepo) ReadProduct(ctx context.Context, id uuid.UUID) (models.Product, error) {
	pr := models.Product{}
	err := r.db.QueryRow(ctx, getProduct, id).
		Scan(&pr.Id, &pr.Name, &pr.Description, &pr.Price, &pr.ImgSrc, &pr.Rating,
			&pr.Category.Id, &pr.Category.Name, &pr.CountComments)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Product{}, ErrProductNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return models.Product{}, err
	}

	return pr, nil
}

func (r *RecommendationsRepo) ReadProductsFromParentCategory(ctx context.Context, categoryID int64) (models.ProductSlice, error) {
	productSlice := make([]models.Product, 0, models.MinProductsCount)

	rows, err := r.db.Query(ctx, getProductsFromParentCategory, categoryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []models.Product{}, ErrProductsNotFound
		}
		err = fmt.Errorf("error happened in db.QueryContext: %w", err)

		return []models.Product{}, err
	}
	product := models.Product{}
	for rows.Next() {
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ImgSrc,
			&product.Rating,
			&product.Category.Id,
			&product.Category.Name,
			&product.CountComments,
		)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return []models.Product{}, err
		}
		productSlice = append(productSlice, product)
	}
	defer rows.Close()

	return productSlice, nil
}

func (r *RecommendationsRepo) CreateUserActivity(ctx context.Context, userID uuid.UUID) error {
	_, err := r.db.Exec(ctx, createActivity, userID)
	if err != nil {
		err = fmt.Errorf("error happened in Exec: %w", err)

		return err
	}

	return nil
}

func (r *RecommendationsRepo) ReadUserActivity(ctx context.Context, userID uuid.UUID) (models.UserActivityStore, error) {
	oldActivity := models.UserActivityStore{}
	var jsonData []byte
	err := r.db.QueryRow(ctx, getActivity, userID).Scan(&jsonData)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.UserActivityStore{}, ErrActivityNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return models.UserActivityStore{}, err
	}
	_ = oldActivity.UnmarshalJSON(jsonData)

	return oldActivity, nil
}

func (r *RecommendationsRepo) UpdateUserActivity(ctx context.Context, userID uuid.UUID, store models.UserActivityStore) error {

	_, err := r.db.Exec(ctx, updateUserActivity, userID, store)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return err
	}
	return nil
}
