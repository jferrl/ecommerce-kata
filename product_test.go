package ecommercekata

import (
	"testing"
)

func TestProduct_UpdateStock(t *testing.T) {
	type fields struct {
		ID            string
		Name          string
		Price         float64
		StockQuantity int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "update stock successfully",
			fields: fields{
				ID:            "P001",
				Name:          "Laptop",
				Price:         1200.00,
				StockQuantity: 10,
			},
			args: args{
				quantity: 5,
			},
			wantErr: false,
		},
		{
			name: "update stock failed",
			fields: fields{
				ID:            "P001",
				Name:          "Laptop",
				Price:         1200.00,
				StockQuantity: 10,
			},
			args: args{
				quantity: 15,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				ID:            tt.fields.ID,
				Name:          tt.fields.Name,
				Price:         tt.fields.Price,
				StockQuantity: tt.fields.StockQuantity,
			}
			if err := p.UpdateStock(tt.args.quantity); (err != nil) != tt.wantErr {
				t.Errorf("Product.UpdateStock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProduct_IsInStock(t *testing.T) {
	type fields struct {
		ID            string
		Name          string
		Price         float64
		StockQuantity int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "product is in stock",
			fields: fields{
				ID:            "P001",
				Name:          "Laptop",
				Price:         1200.00,
				StockQuantity: 10,
			},
			args: args{
				quantity: 5,
			},
			want: true,
		},
		{
			name: "product is not in stock",
			fields: fields{
				ID:            "P001",
				Name:          "Laptop",
				Price:         1200.00,
				StockQuantity: 10,
			},
			args: args{
				quantity: 15,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				ID:            tt.fields.ID,
				Name:          tt.fields.Name,
				Price:         tt.fields.Price,
				StockQuantity: tt.fields.StockQuantity,
			}
			if got := p.IsInStock(tt.args.quantity); got != tt.want {
				t.Errorf("Product.IsInStock() = %v, want %v", got, tt.want)
			}
		})
	}
}
