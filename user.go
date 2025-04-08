package ecommercekata

import "errors"

type User struct {
	ID          string
	Name        string
	Email       string
	Address     string
	PaymentInfo PaymentInfo
}

type UserService struct{}

func (us *UserService) GetUser(userID string) (User, error) {
	user, exists := users[userID]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}
