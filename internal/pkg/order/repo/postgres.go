package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgtype/pgxtype"
)

const (
	createOrder     = "INSERT INTO order_info (id, profile_id, delivery_at) VALUES ($1, $2, $3);"
	createOrderItem = "INSERT INTO order_item (id, order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4);"
	getProductInfo  = "SELECT name, description, price, imgsrc, rating  FROM product WHERE id=$1;"
	getCurrentOrder = "SELECT p.id AS product_id, p.name AS product_name, p.description AS product_description, p.imgsrc AS product_imgscr, p.rating AS product_rating, oi.quantity AS product_quantity, oi.price AS product_price " +
		"FROM order_item oi " +
		"JOIN product p ON oi.product_id = p.id " +
		"WHERE oi.order_id = $1;"
	getCurrentOrderID = "SELECT oi.id AS order_id " +
		"FROM order_info oi " +
		"JOIN status s ON oi.status_id = s.id " +
		"WHERE oi.profile_id = $1 AND s.name = $2;"
	getOrdersID = "SELECT id AS order_id FROM order_info WHERE profile_id=$1;"
	getCart     = "SELECT id FROM cart WHERE Profile_id=$1 and is_current = true;"
	getProducts = "SELECT p._id , p.name_product, p.Description, p.Price, p.ImgSrc, p.Rating, sc.Quantity " +
		"FROM shopping_cart_item sc JOIN product p ON sc.product_id=p.id WHERE p.id=$1;"
	getProduct = "SELECT id FROM product where id=$1;"
	addProduct = "insert into shopping_cart_item(cart_id, product_id, quantity) values ($1, $2, $3)" +
		" ON CONFLICT ON CONSTRAINT uq_shopping_cart_item_cart_id_product_id " +
		"do update set quantity=$3 WHERE shopping_cart_item.cart_id=$1 and shopping_cart_item.product_id=$2;"
	deleteProduct = "DELETE FROM shopping_cart_item WHERE cart_id=$1 and product_id=$2;"
)

var (
	ErrOrderNotFound          = errors.New("order not found")
	ErrOrdersNotFound         = errors.New("orders not found")
	ErrPoductsInOrderNotFound = errors.New("products in order not found")
	ErrPoductNotFound         = errors.New("product not found")
)

type OrderRepo struct {
	db pgxtype.Querier
}

func NewOrderRepo(db pgxtype.Querier) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

// TODO: Добавить добавление статуса в заказ
func (r *OrderRepo) CreateOrder(ctx context.Context, cart models.Cart, userID uuid.UUID) (models.Order, error) {
	orderID := uuid.New()
	_, err := r.db.Exec(ctx, createOrder, orderID, userID, time.Now().Add(24*time.Hour))
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return models.Order{}, err
	}
	var productsOrder []models.OrderProduct
	order := models.Order{Id: orderID, Products: productsOrder}
	for _, cartProduct := range cart.Products {
		err = r.db.QueryRow(ctx, getProductInfo, cartProduct.Id).Scan(
			&cartProduct.Name,
			&cartProduct.Description,
			&cartProduct.Price,
			&cartProduct.ImgSrc,
			&cartProduct.Rating,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return models.Order{}, ErrPoductNotFound
			}
			err = fmt.Errorf("error happened in row.Scan: %w", err)

			return models.Order{}, err
		}

		orderItemID := uuid.New()
		_, err = r.db.Exec(ctx, createOrderItem,
			orderItemID, orderID, cartProduct.Id, cartProduct.Quantity, cartProduct.Price)
		if err != nil {
			err = fmt.Errorf("error happened in db.Exec: %w", err)

			return models.Order{}, err
		}
		order.Products = append(order.Products, models.OrderProduct{Quantity: cartProduct.Quantity,
			Product: models.Product{Id: cartProduct.Id, Name: cartProduct.Name, Description: cartProduct.Description,
				Price: cartProduct.Price, ImgSrc: cartProduct.ImgSrc, Rating: cartProduct.Rating}})

	}

	return order, nil
}

func (r *OrderRepo) ReadOrderID(ctx context.Context, userID uuid.UUID, status string) (uuid.UUID, error) {
	var orderID uuid.UUID
	err := r.db.QueryRow(ctx, getCurrentOrderID, userID, status).Scan(&orderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.UUID{}, ErrOrderNotFound
		}
		err = fmt.Errorf("error happened in rows.Scan: %w", err)

		return uuid.UUID{}, err
	}

	return orderID, nil
}

func (r *OrderRepo) ReadOrder(ctx context.Context, orderID uuid.UUID) (models.Order, error) {
	rows, err := r.db.Query(ctx, getCurrentOrder, orderID)
	defer rows.Close()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Order{}, ErrPoductsInOrderNotFound
		}
		err = fmt.Errorf("error happened in db.Query: %w", err)

		return models.Order{}, err
	}

	var productsOrder []models.OrderProduct
	var productOrder models.OrderProduct
	order := models.Order{Id: orderID, Products: productsOrder}
	for rows.Next() {
		err = rows.Scan(
			&productOrder.Id,
			&productOrder.Name,
			&productOrder.Description,
			&productOrder.ImgSrc,
			&productOrder.Rating,
			&productOrder.Quantity,
			&productOrder.Price,
		)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return models.Order{}, err
		}
		order.Products = append(order.Products, productOrder)
	}

	return order, nil
}

func (r *OrderRepo) ReadOrdersID(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error) {
	rows, err := r.db.Query(ctx, getOrdersID, userID)
	defer rows.Close()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []uuid.UUID{}, ErrOrdersNotFound
		}
		err = fmt.Errorf("error happened in db.Query: %w", err)

		return []uuid.UUID{}, err
	}

	var orderID uuid.UUID
	var ordersID []uuid.UUID
	for rows.Next() {
		err = rows.Scan(&orderID)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return []uuid.UUID{}, err
		}
		ordersID = append(ordersID, orderID)
	}

	return ordersID, nil
}

//func (r *OrderRepo) ReadOrders(ctx context.Context, ordersID []uuid.UUID) (models.Order, error) {
//	var orders []models.Order
//	for _, orderID := range ordersID {
//		rows, err := r.db.Query(ctx, getCurrentOrder, orderID)
//
//		var productsOrder []models.OrderProduct
//		var productOrder models.OrderProduct
//		order := models.Order{Id: orderID, Products: productsOrder}
//
//		for rows.Next() {
//			err = rows.Scan(
//				&productOrder.Id,
//				&productOrder.Name,
//				&productOrder.Description,
//				&productOrder.ImgSrc,
//				&productOrder.Rating,
//				&productOrder.Quantity,
//				&productOrder.Price,
//			)
//			if err != nil {
//				err = fmt.Errorf("error happened in rows.Scan: %w", err)
//
//				return models.Order{}, err
//			}
//			order.Products = append(order.Products, productOrder)
//		}
//		rows.Close()
//	}
//	rows, err := r.db.Query(ctx, getCurrentOrder, orderID)
//	defer rows.Close()
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return models.Order{}, ErrPoductsInOrderNotFound
//		}
//		err = fmt.Errorf("error happened in db.Query: %w", err)
//
//		return models.Order{}, err
//	}
//
//	var productsOrder []models.OrderProduct
//	var productOrder models.OrderProduct
//	order := models.Order{Id: orderID, Products: productsOrder}
//	for rows.Next() {
//		err = rows.Scan(
//			&productOrder.Id,
//			&productOrder.Name,
//			&productOrder.Description,
//			&productOrder.ImgSrc,
//			&productOrder.Rating,
//			&productOrder.Quantity,
//			&productOrder.Price,
//		)
//		if err != nil {
//			err = fmt.Errorf("error happened in rows.Scan: %w", err)
//
//			return models.Order{}, err
//		}
//		order.Products = append(order.Products, productOrder)
//	}
//
//	return order, nil
//}

//func (r *CartRepo) CheckCart(ctx context.Context, userID uuid.UUID) (models.Cart, error) {
//	cart := models.Cart{}
//	err := r.db.QueryRow(ctx, getCart, userID).Scan(&cart.Id)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return models.Cart{}, ErrCartNotFound
//		}
//		err = fmt.Errorf("error happened in row.Scan: %w", err)
//
//		return models.Cart{}, err
//	}
//	return cart, nil
//}
//
//func (r *CartRepo) ReadCart(ctx context.Context, userID uuid.UUID) (models.Cart, error) {
//	cart, err := r.CheckCart(ctx, userID)
//	if err != nil {
//		return models.Cart{}, err
//	}
//
//	cart, err = r.ReadCartProducts(ctx, cart)
//
//	return cart, err
//}
//
