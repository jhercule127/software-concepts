package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"myapp/internal/adapters/email"
	"myapp/internal/adapters/httphandler"
	"myapp/internal/adapters/postgres"
	"myapp/internal/core"
)

func main() {
	// ── Infrastructure ────────────────────────────────────────────────────────
	// cmd/main.go is the ONLY file that knows about ALL adapters.
	// This is the "composition root" — the single place where the hexagon is wired.
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/orders?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("cannot reach postgres: %v", err)
	}

	// ── Wire adapters → core → handler ────────────────────────────────────────
	// To swap Postgres for SQLite: replace postgres.NewOrderRepo with sqlite.NewOrderRepo.
	// The core.OrderService and httphandler.OrderHandler don't change at all.
	repo     := postgres.NewOrderRepo(db)
	notifier := email.NewNotifier("smtp.example.com")
	svc      := core.NewOrderService(repo, notifier)
	handler  := httphandler.NewOrderHandler(svc)

	// ── HTTP server ───────────────────────────────────────────────────────────
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
