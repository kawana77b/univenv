package common

import (
	"fmt"
)

type UnknownError struct {
	reason string
}

func NewUnknownError(reason string) *UnknownError {
	return &UnknownError{
		reason: reason,
	}
}

func (e *UnknownError) Error() string {
	return fmt.Sprintf("unknown error: %s", e.reason)
}

type InvalidError struct {
	reason string
}

func NewInvalidError(reason string) *InvalidError {
	return &InvalidError{
		reason: reason,
	}
}

func (e *InvalidError) Error() string {
	return fmt.Sprintf("invalid error: %s", e.reason)
}
