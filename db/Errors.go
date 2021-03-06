package db

import "errors"

var (
	// ErrUserAlreadyVerified returns error
	ErrUserAlreadyVerified = errors.New("User already verified")
	// ErrUserNotFound returns error
	ErrUserNotFound = errors.New("Cannot Find User")
	// ErrItemNotFound returns error
	ErrItemNotFound = errors.New("Cannot Find Item")
)
