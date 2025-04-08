package main

import (
	"errors"
	"fmt"
	"time"
)

var products = map[string]Product{
	"P001": {ID: "P001", Name: "Laptop", Price: 1200.00, StockQuantity: 10},
	"P002": {ID: "P002", Name: "Smartphone", Price: 800.00, StockQuantity: 15},
	"P003": {ID: "P003", Name: "Headphones", Price: 150.00, StockQuantity: 20},
}

var orders = make(map[string]Order)
var users = map[string]User{
	"U001": {ID: "U001", Name: "John Doe", Email: "john@example.com", Address: "123 Main St", PaymentInfo: PaymentInfo{CardNumber: "1234-5678-9012-3456", ExpiryDate: "12/25", CVV: "123"}},
}

type Product struct {
	ID            string
	Name          string
	Price         float64
	StockQuantity int
}

type OrderItem struct {
	ProductID string
	Quantity  int
	Price     float64
}

type Order struct {
	ID            string
	UserID        string
	Items         []OrderItem
	TotalAmount   float64
	Status        string
	PaymentStatus string
	CreatedAt     time.Time
}

type PaymentInfo struct {
	CardNumber string
	ExpiryDate string
	CVV        string
}

type User struct {
	ID          string
	Name        string
	Email       string
	Address     string
	PaymentInfo PaymentInfo
}

type ProductService struct{}
type OrderService struct{}
type UserService struct{}
type PaymentService struct{}

func (ps *ProductService) GetProduct(productID string) (Product, error) {
	product, exists := products[productID]
	if !exists {
		return Product{}, errors.New("product not found")
	}
	return product, nil
}

func (ps *ProductService) UpdateProductStock(productID string, quantity int) error {
	product, exists := products[productID]
	if !exists {
		return errors.New("product not found")
	}
	if product.StockQuantity < quantity {
		return errors.New("insufficient stock")
	}
	product.StockQuantity -= quantity
	products[productID] = product
	return nil
}

func (us *UserService) GetUser(userID string) (User, error) {
	user, exists := users[userID]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func (os *OrderService) CreateOrder(userID string, items []OrderItem) (Order, error) {
	userService := UserService{}
	_, err := userService.GetUser(userID)
	if err != nil {
		return Order{}, err
	}

	productService := ProductService{}
	totalAmount := 0.0
	for i, item := range items {
		product, err := productService.GetProduct(item.ProductID)
		if err != nil {
			return Order{}, err
		}

		if product.StockQuantity < item.Quantity {
			return Order{}, errors.New("insufficient stock for product: " + product.Name)
		}

		items[i].Price = product.Price
		totalAmount += product.Price * float64(item.Quantity)
	}

	orderID := fmt.Sprintf("O%03d", len(orders)+1)
	order := Order{
		ID:            orderID,
		UserID:        userID,
		Items:         items,
		TotalAmount:   totalAmount,
		Status:        "PENDING",
		PaymentStatus: "UNPAID",
		CreatedAt:     time.Now(),
	}
	orders[orderID] = order

	return order, nil
}

func (os *OrderService) ProcessOrder(orderID string) error {
	order, exists := orders[orderID]
	if !exists {
		return errors.New("order not found")
	}

	if order.Status != "PENDING" {
		return errors.New("order is already processed")
	}

	paymentService := PaymentService{}
	err := paymentService.ProcessPayment(order)
	if err != nil {
		return err
	}

	productService := ProductService{}
	for _, item := range order.Items {
		err := productService.UpdateProductStock(item.ProductID, item.Quantity)
		if err != nil {
			return err
		}
	}

	order.Status = "PROCESSED"
	order.PaymentStatus = "PAID"
	orders[orderID] = order

	return nil
}

func (ps *PaymentService) ProcessPayment(order Order) error {
	userService := UserService{}
	user, err := userService.GetUser(order.UserID)
	if err != nil {
		return err
	}

	if user.PaymentInfo.CardNumber == "" {
		return errors.New("payment information not available")
	}

	fmt.Printf("Processing payment of $%.2f for order %s with card ending in %s\n",
		order.TotalAmount,
		order.ID,
		user.PaymentInfo.CardNumber[len(user.PaymentInfo.CardNumber)-4:])

	return nil
}

func main() {
	orderService := OrderService{}

	orderItems := []OrderItem{
		{ProductID: "P001", Quantity: 1},
		{ProductID: "P003", Quantity: 2},
	}

	order, err := orderService.CreateOrder("U001", orderItems)
	if err != nil {
		fmt.Println("Error creating order:", err)
		return
	}

	fmt.Printf("Order created with ID: %s, Total Amount: $%.2f\n", order.ID, order.TotalAmount)

	err = orderService.ProcessOrder(order.ID)
	if err != nil {
		fmt.Println("Error processing order:", err)
		return
	}

	processedOrder := orders[order.ID]
	fmt.Printf("Order %s status: %s, Payment status: %s\n",
		processedOrder.ID,
		processedOrder.Status,
		processedOrder.PaymentStatus)

	for _, product := range products {
		fmt.Printf("Product: %s, Stock remaining: %d\n", product.Name, product.StockQuantity)
	}
}
