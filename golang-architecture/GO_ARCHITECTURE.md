# Go Design Principles & Code Review Guide

## Package-Based / Domain-Driven Design

Rather than organizing code by technical layers (`/controllers`, `/services`, `/repositories`), Go encourages organizing by **feature domain**:

- `/internal/user` — holds user models, service logic, and repository all together
- `/internal/order` — encapsulates all order-related code

## Flat Architecture

For small Go programs that do one thing, a flat structure (all files in a single package) is perfectly acceptable and common in the ecosystem.

***

## SOLID Principles in Go

| Principle | Go Implementation |
|---|---|
| **Single Responsibility** | Each struct, function, or package does one thing |
| **Open/Closed** | Extend behavior via interfaces, not by modifying existing structs. Software entities should be open for extension, closed for modification |
| **Liskov Substitution** | Use interface composition instead of inheritance. A subtype should not remove behavior from its parent or throw unexpected exceptions that the base type wouldn't |
| **Interface Segregation** | Keep interfaces small — Go strongly favors tiny, focused interfaces |
| **Dependency Inversion** | Depend on interfaces, not concrete types. High-level modules should not depend on low-level modules; both should depend on abstractions |

### LSP — Classic Bird/Flying Violation

A common LSP violation is forcing a non-flying bird to implement a `Fly()` method:

```go
// ❌ BAD: Violates LSP — Penguin can't fly, but is forced to implement Fly()
type Bird interface {
    Eat()
    Fly()
}

type Sparrow struct{}
func (s Sparrow) Eat() { fmt.Println("Sparrow eating") }
func (s Sparrow) Fly() { fmt.Println("Sparrow flying") }

type Penguin struct{}
func (p Penguin) Eat() { fmt.Println("Penguin eating") }
func (p Penguin) Fly() { panic("Penguins can't fly!") } // 💥 Breaks LSP
```

### Dependency Inversion — Direct Dependency Violation

```go
// ❌ BAD: High-level service depends directly on low-level concrete type
package main

import "fmt"

type MySQLDatabase struct{}

func (db MySQLDatabase) Save(order string) {
    fmt.Println("Saving to MySQL:", order)
}

type OrderService struct {
    db MySQLDatabase // ← concrete dependency, hard to swap or test
}

func (s OrderService) PlaceOrder(order string) {
    s.db.Save(order)
}
```

```go
// ✅ GOOD: Both layers depend on the abstraction (interface)
package main

import "fmt"

// Abstraction — defined by the high-level module
type Database interface {
    Save(order string)
}

// Low-level module: MySQL implementation
type MySQLDatabase struct{}

func (db MySQLDatabase) Save(order string) {
    fmt.Println("Saving to MySQL:", order)
}

// Low-level module: Mock for testing
type MockDatabase struct {
    SavedOrders []string
}

func (db *MockDatabase) Save(order string) {
    db.SavedOrders = append(db.SavedOrders, order)
    fmt.Println("Mock saving:", order)
}

// High-level module: depends only on the interface, not a concrete type
type OrderService struct {
    db Database // ← abstraction
}

func NewOrderService(db Database) *OrderService {
    return &OrderService{db: db}
}

func (s *OrderService) PlaceOrder(order string) {
    s.db.Save(order)
}
```

> The `Database` interface is owned by the high-level module (`OrderService`), not by the low-level implementations — this is the "inversion" part.

***

## How to Review Go Code

> Most languages reward cleverness. Go rewards boredom.

```go
// Go rewards clarity over cleverness.
// If the standard library does it a certain way, follow it.
// The linter is almost always right.

// BAD: clever one-liner that obscures intent
result := func(x int) int { if x > 0 { return x }; return -x }(val)

// GOOD: boring, obvious, readable
result := val
if val < 0 {
    result = -val
}
```

### Idiomatic Go

- Clear, not clever
- If the standard library does it a certain way, go with it
- The linter is almost always right

***

## Naming Conventions

Packages should describe one clear purpose:

| Idiomatic | Anti-pattern |
|---|---|
| `user` | `utils` |
| `httputil` | `helpers` |
| `tokens` | `misc` |

### Initialisms

Initialisms must be all-caps or all-lowercase — never mixed case:

```go
// GOOD
var userID int
var requestURL string
type HTTPClient struct{}

// BAD
var userId int        // should be userID
var requestUrl string // should be requestURL
type HttpClient struct{} // should be HTTPClient
```

Flag `snake_case` or `SCREAMING_CASE` as violations:

```go
// GOOD
var userName string
const maxRetries = 3
```

### Getters

In Go, a no-argument method returning a field is simply named after the field itself — no `Get` prefix:

```go
// GOOD: named after the field itself
func (u User) Name() string {
    return u.name
}

// BAD: Java-style getter
func (u User) GetName() string {
    return u.name
}
```

***

## Interfaces

Define interfaces **where they are consumed**, not next to the implementation.

### Why Not Return an Interface from a Constructor?

Returning an interface from a constructor limits the caller and narrows abstraction if they want to extend behavior later:

```go
// GOOD: accept an interface, return a concrete *Server.
// Returning an interface narrows the caller's options unnecessarily.
type Store interface {
    Get(key string) (string, error)
}

func NewServer(store Store) *Server {
    return &Server{store: store}
}

// BAD: returning an interface from a constructor
func NewServer(store Store) Store {
    return &Server{store: store} // limits the caller
}
```

### `I` Prefixes and `Impl` Suffixes

Naming an interface `IUserStore` or an implementation `UserStoreImpl` are anti-patterns imported from other languages:

```go
// BAD: anti-patterns imported from Java/C#
type IUserStore interface{ ... }
type UserStoreImpl struct{ ... }

// GOOD
type UserStore interface{ ... }
type PostgresUserStore struct{ ... }  // concrete name describes the impl
```

### Pointers to Interfaces

Seeing something like `*io.Reader` is almost always wrong — an interface already contains a pointer internally.

**Rule of thumb: Accept interfaces, return concrete types.**

***

## Errors

### Wrapping Errors

Flattening errors using `fmt.Errorf("...:%s", err)` or string concatenation destroys the error chain. Use `%w` to wrap errors and preserve the chain for `errors.Is` / `errors.As`:

```go
// GOOD: %w preserves the chain for errors.Is / errors.As
return fmt.Errorf("findUser: %w", err)
```

### Matching Errors

Always use `errors.Is()` — never match errors with `==`.

### Error Types

| Error Type | Use Case | Example |
|---|---|---|
| Sentinel | Known condition | `io.EOF` |
| Typed | Structured data | `*os.PathError` |

```go
// Sentinel: signals a known, fixed condition
var ErrNotFound = errors.New("not found")

// Typed: carries structured data, match with errors.As
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

var valErr *ValidationError
if errors.As(err, &valErr) {
    fmt.Println("bad field:", valErr.Field)
}
```

***

## Receivers

- **Naming**: Receiver names should be 1–3 letters matching the type name (e.g., `s` for `Server`) — never `this` or `self`
- **Consistency**: Do not mix pointer and value receivers across methods of the same type — pick one and use it for all methods

***

## Concurrency Bugs

### Fire-and-Forget Goroutines

`go func()` calls that fire and forget are dangerous — every goroutine needs a stop condition.

***

## Contexts

- `context.Context` should always be the **first parameter** of a function — never stored in a struct field
- A function that accepts a context implies it respects cancellation

### Cancellable Polling with `time.Sleep`

`time.Sleep` cannot be cancelled. If you see it inside a context-aware function, that's a cancellation bug. Use a `select` statement waiting on `ctx.Done()` alongside a timer instead:

```go
// BAD: time.Sleep cannot be cancelled
func poll(ctx context.Context) {
    for {
        time.Sleep(5 * time.Second) // ignores cancellation entirely
        doWork()
    }
}

// GOOD: select on ctx.Done() alongside a timer
func poll(ctx context.Context, interval time.Duration) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        case <-time.After(interval):
            doWork()
        }
    }
}
```