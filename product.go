package ecommercekata

import "errors"

type Product struct {
	ID            string
	Name          string
	Price         float64
	StockQuantity int
}

func (p *Product) IsInStock(quantity int) bool {
	return p.StockQuantity >= quantity
}

func (p *Product) UpdateStock(quantity int) error {
	if p.StockQuantity < quantity {
		return errors.New("insufficient stock")
	}
	p.StockQuantity -= quantity
	return nil
}

type ProductService struct{}

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
	if err := product.UpdateStock(quantity); err != nil {
		return err
	}
	products[productID] = product
	return nil
}
