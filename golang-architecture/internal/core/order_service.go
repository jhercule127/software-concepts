package core

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"myapp/internal/domain"
	"myapp/internal/ports"
)

// OrderService implements ports.OrderService (the driving port).
// It depends on ports.OrderRepository and ports.Notifier (driven ports) —
// NEVER on their concrete implementations. This is the heart of the hexagon.
type OrderService struct {
	repo     ports.OrderRepository // injected — could be Postgres, SQLite, in-memory
	notifier ports.Notifier        // injected — could be email, SMS, or a no-op
}

// NewOrderService constructs the service with its dependencies (manual DI).
// In larger projects you'd use wire or fx for this.
func NewOrderService(repo ports.OrderRepository, notifier ports.Notifier) *OrderService {
	return &OrderService{repo: repo, notifier: notifier}
}

// PlaceOrder is a use case: validate → create entity → persist → notify.
func (s *OrderService) PlaceOrder(ctx context.Context, customerID string, items []domain.Item) (domain.Order, error) {
	if len(items) == 0 {
		return domain.Order{}, domain.ErrEmptyOrder
	}

	order := domain.Order{
		ID:         uuid.NewString(),
		CustomerID: customerID,
		Items:      items,
		Status:     domain.StatusPending,
		CreatedAt:  time.Now(),
	}

	if err := s.repo.Save(ctx, order); err != nil {
		return domain.Order{}, fmt.Errorf("saving order: %w", err)
	}

	// Notify — if notifier fails, we log but don't roll back the order.
	// This is a deliberate business decision expressed here in the core.
	_ = s.notifier.NotifyOrderConfirmed(ctx, order)

	return order, nil
}

func (s *OrderService) GetOrder(ctx context.Context, id string) (domain.Order, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *OrderService) CancelOrder(ctx context.Context, id string) error {
	order, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Cancel() enforces the business rule — the service just orchestrates.
	if err := order.Cancel(); err != nil {
		return err // domain.ErrAlreadyCancelled
	}

	return s.repo.Save(ctx, order)
}
