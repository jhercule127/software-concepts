package ports

import (
	"context"

	"myapp/internal/domain"
)

// OrderRepository is a DRIVEN port — the core depends on this interface,
// but an adapter (Postgres, in-memory, etc.) provides the actual implementation.
// Dependency arrows always point INWARD: adapters → ports ← core.
type OrderRepository interface {
	Save(ctx context.Context, order domain.Order) error
	FindByID(ctx context.Context, id string) (domain.Order, error)
	FindByCustomer(ctx context.Context, customerID string) ([]domain.Order, error)
}

// Notifier is another driven port — the core never knows if it's email, SMS, Slack, or a no-op.
type Notifier interface {
	NotifyOrderConfirmed(ctx context.Context, order domain.Order) error
}
