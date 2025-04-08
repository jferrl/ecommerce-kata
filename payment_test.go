package ecommercekata

import (
	"testing"
)

func TestPaymentInfo_Validate(t *testing.T) {
	type fields struct {
		CardNumber string
		ExpiryDate string
		CVV        string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "validate payment info successfully",
			fields: fields{
				CardNumber: "1234-5678-9012-3456",
				ExpiryDate: "12/25",
				CVV:        "123",
			},
			wantErr: false,
		},
		{
			name:    "validate payment info failed because of invalid card number",
			fields:  fields{},
			wantErr: true,
		},
		{
			name:    "validate payment info failed because of invalid expiry date",
			fields:  fields{CardNumber: "1234-5678-9012-3456"},
			wantErr: true,
		},
		{
			name:    "validate payment info failed because of invalid CVV",
			fields:  fields{CardNumber: "1234-5678-9012-3456", ExpiryDate: "12/25"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pi := &PaymentInfo{
				cardNumber: tt.fields.CardNumber,
				expiryDate: tt.fields.ExpiryDate,
				cvv:        tt.fields.CVV,
			}
			if err := pi.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("PaymentInfo.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPaymentInfo_Last4Digits(t *testing.T) {
	type fields struct {
		CardNumber string
		ExpiryDate string
		CVV        string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get last 4 digits successfully",
			fields: fields{
				CardNumber: "1234-5678-9012-3456",
				ExpiryDate: "12/25",
				CVV:        "123",
			},
			want: "3456",
		},
		{
			name: "get last 4 digits failed because of invalid card number",
			fields: fields{
				CardNumber: "111",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pi := &PaymentInfo{
				cardNumber: tt.fields.CardNumber,
				expiryDate: tt.fields.ExpiryDate,
				cvv:        tt.fields.CVV,
			}
			if got := pi.Last4Digits(); got != tt.want {
				t.Errorf("PaymentInfo.Last4Digits() = %v, want %v", got, tt.want)
			}
		})
	}
}
