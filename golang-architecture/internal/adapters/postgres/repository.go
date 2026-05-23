package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"myapp/internal/domain"
)

// OrderRepo is the Postgres adapter — it satisfies ports.OrderRepository.
// The core never imports this package; the wiring happens in cmd/main.go only.
type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Save(ctx context.Context, order domain.Order) error {
	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		return fmt.Errorf("marshalling items: %w", err)
	}

	_, err = r.db.ExecContext(ctx,
		`INSERT INTO orders (id, customer_id, items, status, created_at)
		 VALUES ($1, $2, $3, $4, $5)
		 ON CONFLICT (id) DO UPDATE
		   SET status = EXCLUDED.status`,
		order.ID, order.CustomerID, itemsJSON, order.Status, order.CreatedAt,
	)
	return err
}

func (r *OrderRepo) FindByID(ctx context.Context, id string) (domain.Order, error) {
	row := r.db.QueryRowContext(ctx,
		`SELECT id, customer_id, items, status, created_at FROM orders WHERE id = $1`, id,
	)

	var o domain.Order
	var itemsJSON []byte

	if err := row.Scan(&o.ID, &o.CustomerID, &itemsJSON, &o.Status, &o.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return domain.Order{}, domain.ErrOrderNotFound
		}
		return domain.Order{}, err
	}

	return o, json.Unmarshal(itemsJSON, &o.Items)
}

func (r *OrderRepo) FindByCustomer(ctx context.Context, customerID string) ([]domain.Order, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, customer_id, items, status, created_at FROM orders WHERE customer_id = $1`, customerID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var o domain.Order
		var itemsJSON []byte
		if err := rows.Scan(&o.ID, &o.CustomerID, &itemsJSON, &o.Status, &o.CreatedAt); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(itemsJSON, &o.Items); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, rows.Err()
}
