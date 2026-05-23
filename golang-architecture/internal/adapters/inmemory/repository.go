package inmemory

import (
	"context"
	"sync"

	"myapp/internal/domain"
)

// OrderRepo is an in-memory adapter that satisfies ports.OrderRepository.
// Use this in tests — no database, no docker, instant feedback.
type OrderRepo struct {
	mu     sync.RWMutex
	orders map[string]domain.Order
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{orders: make(map[string]domain.Order)}
}

func (r *OrderRepo) Save(_ context.Context, o domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[o.ID] = o
	return nil
}

func (r *OrderRepo) FindByID(_ context.Context, id string) (domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	o, ok := r.orders[id]
	if !ok {
		return domain.Order{}, domain.ErrOrderNotFound
	}
	return o, nil
}

func (r *OrderRepo) FindByCustomer(_ context.Context, customerID string) ([]domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []domain.Order
	for _, o := range r.orders {
		if o.CustomerID == customerID {
			result = append(result, o)
		}
	}
	return result, nil
}
