package core_test

import (
	"context"
	"errors"
	"testing"

	"myapp/internal/adapters/email"
	"myapp/internal/adapters/inmemory"
	"myapp/internal/core"
	"myapp/internal/domain"
)

// helper to build a wired service with zero infrastructure
func newTestService() *core.OrderService {
	return core.NewOrderService(
		inmemory.NewOrderRepo(),
		&email.NoopNotifier{},
	)
}

func TestPlaceOrder_Success(t *testing.T) {
	svc := newTestService()

	order, err := svc.PlaceOrder(context.Background(), "cust-1", []domain.Item{
		{ProductID: "p1", Quantity: 2, UnitPrice: 9.99},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if order.ID == "" {
		t.Fatal("expected a non-empty order ID")
	}
	if order.Status != domain.StatusPending {
		t.Fatalf("expected status pending, got %s", order.Status)
	}
}

func TestPlaceOrder_EmptyItems(t *testing.T) {
	svc := newTestService()

	_, err := svc.PlaceOrder(context.Background(), "cust-1", nil)
	if !errors.Is(err, domain.ErrEmptyOrder) {
		t.Fatalf("expected ErrEmptyOrder, got %v", err)
	}
}

func TestCancelOrder_Success(t *testing.T) {
	svc := newTestService()
	ctx := context.Background()

	order, _ := svc.PlaceOrder(ctx, "cust-1", []domain.Item{
		{ProductID: "p1", Quantity: 1, UnitPrice: 5.00},
	})

	if err := svc.CancelOrder(ctx, order.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cancelled, _ := svc.GetOrder(ctx, order.ID)
	if cancelled.Status != domain.StatusCancelled {
		t.Fatalf("expected cancelled status, got %s", cancelled.Status)
	}
}

func TestCancelOrder_AlreadyCancelled(t *testing.T) {
	svc := newTestService()
	ctx := context.Background()

	order, _ := svc.PlaceOrder(ctx, "cust-1", []domain.Item{
		{ProductID: "p1", Quantity: 1, UnitPrice: 5.00},
	})

	svc.CancelOrder(ctx, order.ID) // first cancel — OK

	err := svc.CancelOrder(ctx, order.ID) // second — should fail
	if !errors.Is(err, domain.ErrAlreadyCancelled) {
		t.Fatalf("expected ErrAlreadyCancelled, got %v", err)
	}
}

func TestGetOrder_NotFound(t *testing.T) {
	svc := newTestService()

	_, err := svc.GetOrder(context.Background(), "nonexistent-id")
	if !errors.Is(err, domain.ErrOrderNotFound) {
		t.Fatalf("expected ErrOrderNotFound, got %v", err)
	}
}

func TestTotalPrice(t *testing.T) {
	order := domain.Order{
		Items: []domain.Item{
			{ProductID: "p1", Quantity: 2, UnitPrice: 10.00},
			{ProductID: "p2", Quantity: 1, UnitPrice: 5.50},
		},
	}

	expected := 25.50
	if got := order.TotalPrice(); got != expected {
		t.Fatalf("expected %.2f, got %.2f", expected, got)
	}
}
