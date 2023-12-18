package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	uuid "github.com/satori/go.uuid"
)

const (
	createOrder = `
	INSERT INTO order_info (id, profile_id, status_id, address_id, delivery_at_time, delivery_at_date, promocode_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7);
 	`
	createOrderWithoutPromo = `
	 INSERT INTO order_info (id, profile_id, status_id, address_id, delivery_at_time, delivery_at_date)
	 VALUES ($1, $2, $3, $4, $5, $6);
	  `

	createOrderItem = "INSERT INTO order_item (id, order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4, $5);"

	getProductInfo = `
	SELECT p.name, p.description, p.price, p.imgsrc, p.rating, 
    c.id AS category_id, c.name AS category_name
	FROM product p
	JOIN category c ON p.category_id = c.id
	WHERE p.id = $1;
	`

	getCurrentOrder = `
	SELECT p.id AS product_id, p.name AS product_name, p.description AS product_description, p.price AS product_price, 
	p.imgsrc AS product_imgsrc, p.rating AS product_rating, oi.quantity AS product_quantity, c.id AS category_id, 
	c.name AS category_name, s.name, pm.name, a.id AS address_id, a.city AS address_city, 
	a.street AS address_street, a.house AS address_house, a.flat AS address_flat, a.is_current as is_current,
	o.creation_at, o.delivery_at_time, o.delivery_at_date
	FROM order_item oi
	JOIN product p ON oi.product_id = p.id
	JOIN order_info o ON oi.order_id = o.id
	JOIN category c ON p.category_id = c.id
	JOIN address a ON o.address_id = a.id
	JOIN status s ON o.status_id = s.id
	JOIN promocode pm ON o.promocode_id = pm.id
	WHERE oi.order_id = $1;
	`

	getCurrentOrderID = "SELECT oi.id AS order_id " +
		"FROM order_info oi " +
		"WHERE oi.profile_id = $1 " +
		"ORDER BY oi.creation_at DESC;"

	getOrdersID = `
	SELECT id AS order_id
	FROM order_info
	WHERE profile_id = $1
	ORDER BY creation_at DESC;
	`

	getUpdates = `
	SELECT user_id, created, message_info
	FROM messages
	WHERE user_id = $1 AND created > $2
	ORDER BY created ASC;
	`

	SetCurrentAddressToOrder = `
	UPDATE order_info
	SET address_id = $1
	WHERE id = $2;
	`
)

//CREATE OR ALTER PROCEDURE create_order(a UUiD, b UUID, c TIMESTAMPTZ, e int, )
//LANGUAGE SQL
//AS $$
//INSERT INTO order_info (id, profile_id, delivery_at, status_id) VALUES (a, b, c, e);
//INSERT INTO order_item (id, order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4);
//$$;

// iota + map = status

var (
	ErrOrderNotFound          = errors.New("order not found")
	ErrOrdersNotFound         = errors.New("orders not found")
	ErrPoductsInOrderNotFound = errors.New("products in order not found")
	ErrMessageNotFound        = errors.New("message not found")
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

func (r *OrderRepo) GetUpdates(ctx context.Context, userID uuid.UUID, time time.Time) ([]models.Message, error) {
	rows, err := r.db.Query(ctx, getUpdates, userID, time)
	defer rows.Close()
	if err != nil {
		err = fmt.Errorf("error happened in db.Query: %w", err)

		return []models.Message{}, err
	}

	var message models.Message
	var messages []models.Message
	for rows.Next() {
		err = rows.Scan(
			&message.UserID,
			&message.Created,
			&message.MessageInfo,
		)
		if err != nil {
			err = fmt.Errorf("error happened in rows.Scan: %w", err)

			return []models.Message{}, err
		}
		messages = append(messages, message)
	}
	if len(messages) == 0 {
		return []models.Message{}, ErrMessageNotFound
	}

	return messages, nil
}

func (r *OrderRepo) CreateOrder(ctx context.Context, cart models.Cart, addressID uuid.UUID, userID uuid.UUID,
	statusID, promocodeID int64, deliveryTime, deliveryDate string) (models.Order, error) {
	orderID := uuid.NewV4()

	if promocodeID != -1 {
		_, err := r.db.Exec(ctx, createOrder, orderID, userID,
			statusID, addressID, deliveryTime, deliveryDate, promocodeID)
		if err != nil {
			err = fmt.Errorf("error happened in db.Exec: %w", err)

			return models.Order{}, err
		}
	} else {
		_, err := r.db.Exec(ctx, createOrderWithoutPromo, orderID, userID,
			statusID, addressID, deliveryTime, deliveryDate)
		if err != nil {
			err = fmt.Errorf("error happened in db.Exec: %w", err)

			return models.Order{}, err
		}
	}

	var productsOrder []models.OrderProduct
	order := models.Order{
		Id:       orderID,
		Products: productsOrder,
	}
	for _, cartProduct := range cart.Products {
		orderProduct := models.OrderProduct{Quantity: cartProduct.Quantity, Product: models.Product{Id: cartProduct.Id}}
		err := r.db.QueryRow(ctx, getProductInfo, cartProduct.Id).Scan(
			&orderProduct.Name,
			&orderProduct.Description,
			&orderProduct.Price,
			&orderProduct.ImgSrc,
			&orderProduct.Rating,
			&orderProduct.Category.Id,
			&orderProduct.Category.Name,
		)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return models.Order{}, ErrPoductNotFound
			}
			err = fmt.Errorf("error happened in row.Scan: %w", err)

			return models.Order{}, err
		}

		orderItemID := uuid.NewV4()
		_, err = r.db.Exec(ctx, createOrderItem,
			orderItemID, orderID, cartProduct.Id, cartProduct.Quantity, cartProduct.Price)
		if err != nil {
			err = fmt.Errorf("error happened in db.Exec: %w", err)

			return models.Order{}, err
		}
		order.Products = append(order.Products, orderProduct)
	}

	return order, nil
}

func (r *OrderRepo) ReadOrderID(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	var orderID uuid.UUID
	err := r.db.QueryRow(ctx, getCurrentOrderID, userID).Scan(&orderID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
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
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Order{}, ErrPoductsInOrderNotFound
		}
		err = fmt.Errorf("error happened in db.Query: %w", err)

		return models.Order{}, err
	}

	var productOrder models.OrderProduct
	var productsOrder []models.OrderProduct
	order := models.Order{Id: orderID, Products: productsOrder}
	for rows.Next() {
		err = rows.Scan(
			&productOrder.Id,
			&productOrder.Name,
			&productOrder.Description,
			&productOrder.Price,
			&productOrder.ImgSrc,
			&productOrder.Rating,
			&productOrder.Quantity,
			&productOrder.Category.Id,
			&productOrder.Category.Name,
			&order.Status,
			&order.PomocodeName,
			&order.Address.Id,
			&order.Address.City,
			&order.Address.Street,
			&order.Address.House,
			&order.Address.Flat,
			&order.Address.IsCurrent,
			&order.CreationAt,
			&order.DeliveryTime,
			&order.DeliveryDate,
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
		if errors.Is(err, pgx.ErrNoRows) {
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
	if len(ordersID) == 0 {
		return []uuid.UUID{}, ErrOrdersNotFound
	}

	return ordersID, nil
}
