package service

import (
	"fmt"
)

type ErrorCode int

const (
	ErrorCodeInvalidArgument ErrorCode = iota
	ErrorCodeNotFound
)

type Error struct {
	Code    ErrorCode
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("error: %s (code: %d)", e.Message, e.Code)
}
