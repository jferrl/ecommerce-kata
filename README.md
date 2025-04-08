# Go Refactoring Workshop: Tell Don't Ask, Rich Domain Models & Deep Modules

This repository demonstrates the application of key software design principles through a practical refactoring exercise of a Go e-commerce application.

## 🎯 Learning Goals

This workshop focuses on three important software design principles:

1. **Tell, Don't Ask**  
   A principle that encourages telling objects what to do rather than asking them for data and making decisions outside. This reduces coupling and improves encapsulation.

2. **Rich Domain Models (vs. Anemic Domain Models)**  
   Moving beyond data-only objects toward entities that encapsulate both data and behavior, enforcing their own business rules and invariants.

3. **Deep Modules**  
   As described in "A Philosophy of Software Design" by John Ousterhout, deep modules provide powerful functionality while hiding implementation complexity behind simple interfaces.

## 🔍 The Example: E-commerce Order System

The repository contains a simplified e-commerce order processing system with:

- Products with inventory management
- Users with payment information
- Order creation and processing
- Payment handling

## 💡 Key Refactorings Demonstrated

### Tell, Don't Ask

**Before:**
```go
// Checking conditions outside the object
if product.StockQuantity < quantity {
    return errors.New("insufficient stock")
}
product.StockQuantity -= quantity
```

**After:**
```go
// Telling the object what to do
err := product.DecreaseStock(quantity)
if err != nil {
    return err
}
```

### Rich Domain Model

**Before:**
```go
// Anemic domain model - just data
type Order struct {
    ID            string
    UserID        string
    Items         []OrderItem
    TotalAmount   float64
    Status        string
    PaymentStatus string
    CreatedAt     time.Time
}

// Business logic in services
func (os *OrderService) ProcessOrder(orderID string) error {
    // ... logic for processing order
}
```

**After:**
```go
// Rich domain model with behavior
type Order struct {
    // Same fields as before
}

// Business logic in domain objects
func (o *Order) Process() error {
    if o.Status != "PENDING" {
        return errors.New("order already processed")
    }
    if o.PaymentStatus != "PAID" {
        return errors.New("cannot process unpaid order")
    }
    o.Status = "PROCESSED"
    return nil
}
```

### Deep Modules

**Before:**
- Shallow services with minimal functionality
- High coupling between services
- Implementation details exposed

**After:**
- Domain repositories that abstract storage concerns
- Meaningful interfaces that hide complexity
- Clear separation of domain and infrastructure concerns

## 🚀 How to Use This Repository

1. Start by examining the code in `before/main.go` to identify issues
2. Study the refactored implementation in the `after/` directory
3. Complete the exercises in the `exercises/` directory
4. Refer to the documentation for deeper understanding

## 📚 Further Reading

- "A Philosophy of Software Design" by John Ousterhout
- "Domain-Driven Design" by Eric Evans
- "Clean Architecture" by Robert C. Martin

## 🔗 Additional Resources

- [Tell Don't Ask Principle](https://martinfowler.com/bliki/TellDontAsk.html)
- [Anemic Domain Model](https://martinfowler.com/bliki/AnemicDomainModel.html)
- [John Ousterhout's lectures on software design](https://milkov.tech/assets/psd.pdf)

## 📋 Contributing

Feel free to submit pull requests with additional examples, exercises, or documentation improvements!
