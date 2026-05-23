# Hexagonal Architecture in Go (Ports & Adapters)

A clean, idiomatic Go example of Hexagonal Architecture using an **Order** domain.

---

## What Is Hexagonal Architecture?

This is widely considered the most idiomatic architecture for Go projects. It focuses on isolating the core business logic from external systems (databases, APIs, UI) by defining ports (interfaces) and adapters (implementations)

Before the code: the hexagon has three distinct zones:

| Zone | What lives here | Go Construct |
| ---- | --------------- | -----------  |
| Core domain |	Entities, business rules |	Plain struct + methods |
| Ports | Contracts between zones | interface |
| Adapters | Implementations of ports | Concrete types (struct) |

The rule: the core never imports adapters — dependency arrows always point inward.


The core idea: **your business logic should never know about your infrastructure.**

```
┌─────────────────────────────────────────────────────────┐
│                      OUTSIDE WORLD                      │
│                                                         │
│   [HTTP Client]   [CLI]   [gRPC]  ← PRIMARY ADAPTERS   │
│         │           │       │                           │
│   ══════╪═══════════╪═══════╪══ DRIVING PORTS ═══════  │
│         ▼           ▼       ▼                           │
│   ┌─────────────────────────────────────────────────┐   │
│   │              CORE (the hexagon)                 │   │
│   │                                                 │   │
│   │   domain/     — entities, business rules        │   │
│   │   core/       — use cases (application service) │   │
│   │   ports/      — interfaces (contracts)          │   │
│   └─────────────────────────────────────────────────┘   │
│         │           │       │                           │
│   ══════╪═══════════╪═══════╪══ DRIVEN PORTS ════════  │
│         ▼           ▼       ▼                           │
│  [Postgres]  [Email]  [Redis]  ← SECONDARY ADAPTERS    │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

**The golden rule:** dependency arrows always point **inward**.
Adapters import ports. The core never imports adapters.

---

## Directory Structure

```
hexagonal-go/
├── cmd/
│   └── main.go                         ← Composition root — wires EVERYTHING
│
├── internal/
│   ├── domain/
│   │   └── order.go                    ← Entities & business rules (NO imports)
│   │
│   ├── ports/
│   │   ├── driven.go                   ← What the app NEEDS  (e.g. DB, email)
│   │   └── driving.go                  ← What DRIVES the app (e.g. HTTP, CLI)
│   │
│   ├── core/
│   │   ├── order_service.go            ← Use cases — imports domain + ports ONLY
│   │   └── order_service_test.go       ← Tests with zero infrastructure
│   │
│   └── adapters/
│       ├── postgres/
│       │   ├── repository.go           ← Secondary (driven) adapter
│       │   └── schema.sql              ← DB schema for reference
│       ├── httphandler/
│       │   └── handler.go              ← Primary (driving) adapter
│       ├── inmemory/
│       │   └── repository.go           ← In-memory adapter (for tests)
│       └── email/
│           └── notifier.go             ← Email adapter + NoopNotifier
│
└── go.mod
```

---

## The Three Zones Explained

### 1. `internal/domain/` — The Core Entities
Pure Go structs and methods. **Zero imports from your own codebase.**
Business rules live here: `order.Cancel()` enforces that you can't cancel twice.
Domain errors (`ErrOrderNotFound`, etc.) are also defined here.

```go
// Business rule — lives in domain, not in HTTP handler or service
func (o *Order) Cancel() error {
    if o.Status == StatusCancelled {
        return ErrAlreadyCancelled  // enforced HERE, not scattered around
    }
    o.Status = StatusCancelled
    return nil
}
```

### 2. `internal/ports/` — The Contracts (Interfaces)
Two files, two directions:

| File | Direction | Who defines it | Who implements it |
|---|---|---|---|
| `driven.go` | Inward ← | Core | Adapters (Postgres, email…) |
| `driving.go` | Outward → | Core | Core itself (service) |

The port is the **interface**. The adapter is the **struct that satisfies it**.

### 3. `internal/adapters/` — The Implementations
Each adapter satisfies a port interface. They know about infrastructure (SQL,
HTTP, SMTP), but the core never knows about them.

**Swapping adapters is trivial** — change one line in `cmd/main.go`:
```go
// Production
repo := postgres.NewOrderRepo(db)

// Tests / local dev
repo := inmemory.NewOrderRepo()

// The rest of the code doesn't change at all.
```

---

## Running the Tests

Tests run with **zero infrastructure** — no Docker, no database, no network:

```bash
go test ./internal/core/...
```

The `inmemory` adapter replaces Postgres; `NoopNotifier` replaces the email
service. The tests exercise real business logic against real domain rules.

---

## Running the Server

Requires a Postgres instance. Update the connection string in `cmd/main.go`:

```bash
# Start Postgres (Docker example)
docker run -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=orders -p 5432:5432 postgres:16

# Apply schema
psql postgres://user:pass@localhost/orders -f internal/adapters/postgres/schema.sql

# Run
go run ./cmd/main.go
```

### API

| Method | Path | Description |
|---|---|---|
| `POST` | `/orders` | Place a new order |
| `GET` | `/orders/{id}` | Fetch an order by ID |
| `DELETE` | `/orders/{id}` | Cancel an order |

```bash
# Place an order
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"customer_id":"cust-1","items":[{"product_id":"p1","quantity":2,"unit_price":9.99}]}'

# Get the order (replace ID with the one returned above)
curl http://localhost:8080/orders/<id>

# Cancel it
curl -X DELETE http://localhost:8080/orders/<id>
```

---

## Key Patterns to Remember

| Pattern | Where | Why |
|---|---|---|
| Interfaces as ports | `internal/ports/` | Decouples core from infrastructure |
| Errors in domain | `internal/domain/order.go` | Single source of truth for business errors |
| `errors.Is()` at adapter boundary | `adapters/httphandler/handler.go` | Translates domain errors → HTTP codes |
| Composition root | `cmd/main.go` | Only place that imports ALL packages |
| In-memory adapters | `adapters/inmemory/` | Fast, zero-infra tests |

---

## Extending This Project

**Add a new driven adapter** (e.g. Redis cache):
1. Create `internal/adapters/redis/repository.go`
2. Implement `ports.OrderRepository`
3. Wire it in `cmd/main.go` — done

**Add a new driving adapter** (e.g. gRPC):
1. Create `internal/adapters/grpc/server.go`
2. Inject `ports.OrderService` into it
3. Wire it in `cmd/main.go` — done

The core (`domain/`, `ports/`, `core/`) **never changes** for either of these.
