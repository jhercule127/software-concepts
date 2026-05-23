package httphandler

import (
	"encoding/json"
	"errors"
	"net/http"

	"myapp/internal/domain"
	"myapp/internal/ports"
)

// OrderHandler is the PRIMARY (driving) adapter.
// It translates HTTP → use-case calls, and use-case results → HTTP responses.
// It depends on ports.OrderService — not on the concrete core.OrderService.
type OrderHandler struct {
	svc ports.OrderService
}

func NewOrderHandler(svc ports.OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

func (h *OrderHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /orders", h.PlaceOrder)
	mux.HandleFunc("GET /orders/{id}", h.GetOrder)
	mux.HandleFunc("DELETE /orders/{id}", h.CancelOrder)
}

// PlaceOrder handles POST /orders
func (h *OrderHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CustomerID string        `json:"customer_id"`
		Items      []domain.Item `json:"items"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	order, err := h.svc.PlaceOrder(r.Context(), req.CustomerID, req.Items)
	if err != nil {
		if errors.Is(err, domain.ErrEmptyOrder) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

// GetOrder handles GET /orders/{id}
func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	order, err := h.svc.GetOrder(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrOrderNotFound) {
			http.Error(w, "order not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// CancelOrder handles DELETE /orders/{id}
func (h *OrderHandler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.svc.CancelOrder(r.Context(), id); err != nil {
		switch {
		case errors.Is(err, domain.ErrOrderNotFound):
			http.Error(w, "order not found", http.StatusNotFound)
		case errors.Is(err, domain.ErrAlreadyCancelled):
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
