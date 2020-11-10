package domain

import "errors"

var (
	ErrInvalidID        = errors.New("Invalid ID")
	ErrCustomerRequired = errors.New("Customer required")
	ErrInvalidValue     = errors.New("Invalid Value")
)
