package domain

import (
	"errors"
	"time"
)

// Order is a pure domain entity — no DB tags, no HTTP concerns.
type Order struct {
	ID         string      `json:"id"`
	CustomerID string      `json:"customer_id"`
	Items      []Item      `json:"items"`
	Status     OrderStatus `json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
}

type Item struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusConfirmed OrderStatus = "confirmed"
	StatusCancelled OrderStatus = "cancelled"
)

// Domain errors — live here, owned by the core.
var (
	ErrOrderNotFound    = errors.New("order not found")
	ErrEmptyOrder       = errors.New("order must have at least one item")
	ErrAlreadyCancelled = errors.New("order is already cancelled")
)

// TotalPrice is pure business logic — no infrastructure needed.
func (o *Order) TotalPrice() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += float64(item.Quantity) * item.UnitPrice
	}
	return total
}

// Cancel encapsulates the business rule for cancellation.
func (o *Order) Cancel() error {
	if o.Status == StatusCancelled {
		return ErrAlreadyCancelled
	}
	o.Status = StatusCancelled
	return nil
}
