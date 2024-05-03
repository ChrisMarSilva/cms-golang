package utils

import (
	"errors"
	"fmt"
)

var (
		ErrInvalidEmail       = errors.New("invalid email")
		ErrEmailAlreadyExists = errors.New("email already exists")
		ErrEmptyPassword      = errors.New("password can't be empty")
		ErrInvalidAuthToken   = errors.New("invalid auth-token")
		ErrInvalidCredentials = errors.New("invalid credentials")
		ErrUnauthorized       = errors.New("Unauthorized")

	ErrNoCapacity   = errors.New("capacity is full")
	ErrTripNotFound = errors.New("this trip does not exist")

	ErrExceedAllowedTicketToPurchase = func(limit int) error {
		return fmt.Errorf("exceed number of tickets allowed to be purchased(%d)", limit)
	}

	ErrExceedMaleTicketNumber = errors.New("exceed number of male ticket allowed to be purchased")
)
