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
