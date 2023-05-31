package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/multierr"

	"yandex-team.ru/bstask/service"
)

func errorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil && !c.Response().Committed {
			var errorCode ErrorCode
			switch err := err.(type) {
			case *Error:
				errorCode = err.Code
			case *service.Error:
				errorCode = serviceErrorCodeMap[err.Code]
			default:
				errorCode = ErrorCodeBadRequest
			}

			code := errorCodeHttpStatusMap[errorCode]
			response := errorCodeResponseMap[errorCode]
			return multierr.Append(err, c.JSON(code, response))
		}
		return nil
	}
}

func rateLimitMiddleware() echo.MiddlewareFunc {
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store:   middleware.NewRateLimiterMemoryStore(10),
		DenyHandler: func(c echo.Context, _ string, err error) error {
			return c.JSON(http.StatusTooManyRequests, ErrorBadRequest)
		},
	}
	return middleware.RateLimiterWithConfig(config)
}
