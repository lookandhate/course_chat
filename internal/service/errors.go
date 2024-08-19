package service

import "errors"

var (
	// ErrChatDoesNotExist - when trying to access not existing chat
	ErrChatDoesNotExist = errors.New("chat does not exist")

	// ErrInvalidID - when passed an invalid ID (<= 0 for integer IDs for example).
	ErrInvalidID = errors.New("invalid id")
)
