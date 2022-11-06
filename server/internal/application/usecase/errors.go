package usecase

import (
	"errors"
)

var (
	ErrInvalidRequestFormat = errors.New("invalid request format")
	ErrNotAuthenticated     = errors.New("not authenticated")
	ErrInternalServerError  = errors.New("internal server error")
	ErrLoginAlreadyUse      = errors.New("the login is already in use")
	ErrUserNotFound         = errors.New("user not found")
	ErrTypeNotFound         = errors.New("type not found")
	ErrEntryNotFound        = errors.New("entry not found")
	ErrSessionNotFound      = errors.New("session not found")
)
