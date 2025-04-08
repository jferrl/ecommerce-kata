package ecommercekata

import (
	"errors"
	"fmt"
)

type PaymentInfo struct {
	cardNumber string
	expiryDate string
	cvv        string
}

func (pi *PaymentInfo) Validate() error {
	if pi.cardNumber == "" {
		return errors.New("card number is required")
	}
	if pi.expiryDate == "" {
		return errors.New("expiry date is required")
	}
	if pi.cvv == "" {
		return errors.New("CVV is required")
	}
	return nil
}

func (pi *PaymentInfo) Last4Digits() string {
	if len(pi.cardNumber) < 4 {
		return ""
	}
	return pi.cardNumber[len(pi.cardNumber)-4:]
}

type PaymentService struct {
	UserService UserService
}

func (ps *PaymentService) ProcessPayment(order Order) error {
	user, err := ps.UserService.GetUser(order.UserID)
	if err != nil {
		return err
	}

	if err := user.PaymentInfo.Validate(); err != nil {
		return err
	}

	fmt.Printf("Processing payment of $%.2f for order %s with card ending in %s\n",
		order.CalculateTotalAmount(),
		order.ID,
		user.PaymentInfo.Last4Digits(),
	)

	return nil
}
