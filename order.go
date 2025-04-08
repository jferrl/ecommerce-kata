package ecommercekata

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type OrderItem struct {
	ProductID string
	Product   Product
	Quantity  int
	Price     float64
}

type Order struct {
	ID            string
	UserID        string
	Items         []OrderItem
	Status        string
	PaymentStatus string
	CreatedAt     time.Time
}

func (o *Order) CalculateTotalAmount() float64 {
	totalAmount := 0.0
	for _, item := range o.Items {
		totalAmount += item.Price * float64(item.Quantity)
	}
	return totalAmount
}

func (o *Order) AddItem(i OrderItem, p Product) error {
	if !p.IsInStock(i.Quantity) {
		return errors.New("insufficient stock for product: " + p.Name)
	}

	o.Items = append(o.Items, i)
	i.Price = p.Price

	return nil
}

func NewOrder(userID string) Order {
	return Order{
		ID:            uuid.NewString(),
		UserID:        userID,
		Status:        "PENDING",
		PaymentStatus: "UNPAID",
		CreatedAt:     time.Now(),
	}
}

type OrderService struct {
	PaymentService PaymentService
	ProductService ProductService
	UserService    UserService
}

func (os *OrderService) CreateOrder(userID string, items []OrderItem) (Order, error) {
	_, err := os.UserService.GetUser(userID)
	if err != nil {
		return Order{}, err
	}

	order := NewOrder(userID)

	for _, item := range items {
		product, err := os.ProductService.GetProduct(item.ProductID)
		if err != nil {
			return Order{}, err
		}

		if err := order.AddItem(item, product); err != nil {
			return Order{}, err
		}
	}

	orders[order.ID] = order

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

	err := os.PaymentService.ProcessPayment(order)
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		err := os.ProductService.UpdateProductStock(item.ProductID, item.Quantity)
		if err != nil {
			return err
		}
	}

	order.Status = "PROCESSED"
	order.PaymentStatus = "PAID"
	orders[orderID] = order

	return nil
}
