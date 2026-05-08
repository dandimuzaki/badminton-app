package utils

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserInactive      = errors.New("user account is inactive")
	ErrInvalidUserID     = errors.New("invalid user ID")
	ErrPasswordMismatch  = errors.New("password mismatch")
	ErrWeakPassword      = errors.New("password is too weak")
	ErrEmailAlreadyUsed  = errors.New("email already registered")
	ErrCourtAlreadyBooked = errors.New("court is already booked")
	ErrInvalidReservation = errors.New("only pending reservations can be canceled")
	ErrUnauthorized = errors.New("unauthorized")
)