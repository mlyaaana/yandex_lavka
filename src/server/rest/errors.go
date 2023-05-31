package rest

import (
	"encoding/json"
	"fmt"

	"yandex-team.ru/bstask/gen"
	"yandex-team.ru/bstask/service"
)

type ErrorCode int

const (
	ErrorCodeBadRequest ErrorCode = iota
	ErrorCodeNotFound
	ErrorCodeTooManyRequests
	ErrorCodeUnknown
)

func (c ErrorCode) MarshalJSON() ([]byte, error) {
	s, ok := errorCodeStringMap[c]
	if !ok {
		s = errorCodeStringMap[ErrorCodeUnknown]
	}
	return json.Marshal(s)
}

type Error struct {
	Code    ErrorCode `json:"error_code"`
	Message string    `json:"error_message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("error: %s (code: %d)", e.Message, e.Code)
}

var (
	ErrorBadRequest = gen.BadRequestResponse{}
	ErrorNotFound   = gen.NotFoundResponse{}
)

var errorCodeStringMap = map[ErrorCode]string{
	ErrorCodeBadRequest:      "BAD_REQUEST",
	ErrorCodeNotFound:        "NOT_FOUND",
	ErrorCodeTooManyRequests: "TOO_MANY_REQUESTS",
	ErrorCodeUnknown:         "UNKNOWN",
}

var errorCodeResponseMap = map[ErrorCode]interface{}{
	ErrorCodeBadRequest:      ErrorBadRequest,
	ErrorCodeNotFound:        ErrorNotFound,
	ErrorCodeTooManyRequests: ErrorBadRequest,
	ErrorCodeUnknown:         ErrorBadRequest,
}

var errorCodeHttpStatusMap = map[ErrorCode]int{
	ErrorCodeBadRequest:      400,
	ErrorCodeNotFound:        404,
	ErrorCodeTooManyRequests: 429,
	ErrorCodeUnknown:         500,
}

var serviceErrorCodeMap = map[service.ErrorCode]ErrorCode{
	service.ErrorCodeInvalidArgument: ErrorCodeBadRequest,
	service.ErrorCodeNotFound:        ErrorCodeNotFound,
}
