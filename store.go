package ecommercekata

var (
	products = map[string]Product{
		"P001": {ID: "P001", Name: "Laptop", Price: 1200.00, StockQuantity: 10},
		"P002": {ID: "P002", Name: "Smartphone", Price: 800.00, StockQuantity: 15},
		"P003": {ID: "P003", Name: "Headphones", Price: 150.00, StockQuantity: 20},
	}
	orders = make(map[string]Order)
	users  = map[string]User{
		"U001": {ID: "U001", Name: "John Doe", Email: "john@example.com", Address: "123 Main St", PaymentInfo: PaymentInfo{cardNumber: "1234-5678-9012-3456", expiryDate: "12/25", cvv: "123"}},
	}
)
