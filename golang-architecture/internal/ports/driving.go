package ports

import (
	"context"

	"myapp/internal/domain"
)

// OrderService is a DRIVING port — the primary adapter (HTTP handler, CLI, gRPC server)
// depends on this interface, not on the concrete service struct.
// This allows you to swap or mock the entire application layer in tests.
type OrderService interface {
	PlaceOrder(ctx context.Context, customerID string, items []domain.Item) (domain.Order, error)
	GetOrder(ctx context.Context, id string) (domain.Order, error)
	CancelOrder(ctx context.Context, id string) error
}
