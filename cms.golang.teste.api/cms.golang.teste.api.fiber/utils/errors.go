package util

import "errors"

var (
	ErrInvalidEmail       = errors.New("invalid email")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrEmptyPassword      = errors.New("password can't be empty")
	ErrInvalidAuthToken   = errors.New("invalid auth-token")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("Unauthorized")
)

// import (
// 	"net/http"
// )

// type AppError struct {
// 	Code    int
// 	Message string
// }

// func (e AppError) Error() string {
// 	return e.Message
// }

// func NewNotFoundError(message string) error {
// 	return AppError{
// 		Code:    http.StatusNotFound,
// 		Message: message,
// 	}
// }

// func NewUnexpectedError() error {
// 	return AppError{
// 		Code:    http.StatusInternalServerError,
// 		Message: "unexpected error",
// 	}
// }

// func NewValidationError(message string) error {
// 	return AppError{
// 		Code:    http.StatusUnprocessableEntity,
// 		Message: message,
// 	}
// }

// func handleError(w http.ResponseWriter, err error) {
// 	switch e := err.(type) {
// 	case errs.AppError:
// 		w.WriteHeader(e.Code)
// 		fmt.Fprintln(w, e)
// 	case error:
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintln(w, e)
// 	}
// }
