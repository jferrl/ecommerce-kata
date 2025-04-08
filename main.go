package ecommercekata

import (
	"fmt"
)

func main() {
	orderService := OrderService{
		PaymentService: PaymentService{},
		ProductService: ProductService{},
		UserService:    UserService{},
	}

	orderItems := []OrderItem{
		{ProductID: "P001", Quantity: 1},
		{ProductID: "P003", Quantity: 2},
	}

	order, err := orderService.CreateOrder("U001", orderItems)
	if err != nil {
		fmt.Println("Error creating order:", err)
		return
	}

	fmt.Printf("Order created with ID: %s, Total Amount: $%.2f\n", order.ID, order.CalculateTotalAmount())

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
