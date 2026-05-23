package email

import (
	"context"
	"fmt"
	"log"

	"myapp/internal/domain"
)

// Notifier is a DRIVEN adapter that satisfies ports.Notifier.
// In production this would call an SMTP server or a service like SendGrid.
type Notifier struct {
	smtpHost string
}

func NewNotifier(smtpHost string) *Notifier {
	return &Notifier{smtpHost: smtpHost}
}

func (n *Notifier) NotifyOrderConfirmed(_ context.Context, order domain.Order) error {
	// Placeholder — swap for real SMTP/SendGrid logic.
	log.Printf("[email] order %s confirmed for customer %s (total: $%.2f) via %s\n",
		order.ID, order.CustomerID, order.TotalPrice(), n.smtpHost)
	return nil
}

// NoopNotifier satisfies ports.Notifier but does nothing.
// Useful for tests where you don't care about notifications.
type NoopNotifier struct{}

func (n *NoopNotifier) NotifyOrderConfirmed(_ context.Context, order domain.Order) error {
	fmt.Printf("[noop] would notify for order %s\n", order.ID)
	return nil
}
