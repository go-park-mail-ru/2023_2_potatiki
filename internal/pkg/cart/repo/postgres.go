package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgtype/pgxtype"
)

const (
	getCart            = "SELECT id FROM cart WHERE ProfileId=$1 and isCurrent = true;"
	createCart         = "INSERT INTO cart(Id, ProfileId, isCurrent) VALUES($1, $2, true);"
	updateCartProducts = "UPDATE shopping_cart_item SET quantity=$3 WHERE cartId=$1 and productId=$2 and quantity<>$3;"
	getProducts        = "SELECT p.Id , p.name_product, p.Description, p.Price, p.ImgSrc, p.Rating, sc.Quantity " +
		"FROM shopping_cart_item sc JOIN product p ON sc.productId=p.id WHERE p.id=$1 and sc.quantity>0;"
)

var (
	ErrCartNotFound    = errors.New("cart not found")
	ErrPoductsNotFound = errors.New("products not found")
)

type CartRepo struct {
	db pgxtype.Querier
}

func NewCartRepo(db pgxtype.Querier) *CartRepo {
	return &CartRepo{
		db: db,
	}
}

func (r *CartRepo) CreateCart(ctx context.Context, userID uuid.UUID) error {
	cartID := uuid.New()
	_, err := r.db.Exec(ctx, createCart, cartID, userID)
	if err != nil {
		err = fmt.Errorf("error happened in rows.Scan: %w", err)

		return err
	}

	return nil
}

func (r *CartRepo) ReadCart(ctx context.Context, userID uuid.UUID) (models.Cart, error) {
	cart := models.Cart{}
	err := r.db.QueryRow(ctx, getCart, userID).Scan(&cart.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Cart{}, ErrCartNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return models.Cart{}, err
	}

	cart, err = r.ReadCartProducts(ctx, cart)

	return cart, err
}

func (r *CartRepo) UpdateCart(ctx context.Context, cart models.Cart) (models.Cart, error) {
	err := r.db.QueryRow(ctx, getCart, cart.ProfileId).Scan(&cart.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Cart{}, ErrCartNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return models.Cart{}, err
	}

	for _, product := range cart.Products {
		_, err := r.db.Exec(ctx, updateCartProducts, cart.Id, product.Id, product.Quantity)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return cart, err
		}
	}

	cart, err = r.ReadCartProducts(ctx, cart)

	return cart, err
}

func (r *CartRepo) ReadCartProducts(ctx context.Context, cart models.Cart) (models.Cart, error) {
	rows, err := r.db.Query(ctx, getProducts, cart.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return cart, ErrPoductsNotFound
		}
		err = fmt.Errorf("error happened in db.QueryContext: %w", err)

		return cart, err
	}

	product := models.CartProduct{}
	for rows.Next() {
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ImgSrc,
			&product.Rating,
			&product.Quantity,
		)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return cart, err
		}
		cart.Products = append(cart.Products, product)
	}
	defer rows.Close()

	return cart, nil
}
