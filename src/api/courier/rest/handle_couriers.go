package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"yandex-team.ru/bstask/api/courier"
	"yandex-team.ru/bstask/gen"
	"yandex-team.ru/bstask/server/rest"
)

func (c *Controller) getCouriers(e echo.Context) error {
	ctx := e.Request().Context()

	limit, err := rest.QueryParamInt(e, "limit", 1)
	if err != nil {
		return &rest.Error{
			Code:    rest.ErrorCodeBadRequest,
			Message: fmt.Sprintf("failed to parse limit: %v", err),
		}
	}
	offset, err := rest.QueryParamInt(e, "offset", 0)
	if err != nil {
		return &rest.Error{
			Code:    rest.ErrorCodeBadRequest,
			Message: fmt.Sprintf("failed to parse offset: %v", err),
		}
	}

	out, err := c.service.ListCouriers(ctx, &courier.ListCouriersIn{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, unpackCouriers(out.Couriers))
}

func (c *Controller) postCouriers(e echo.Context) error {
	ctx := e.Request().Context()

	req := &gen.CreateCourierRequest{}
	if err := e.Bind(req); err != nil {
		return &rest.Error{
			Code:    rest.ErrorCodeBadRequest,
			Message: fmt.Sprintf("failed to parse request: %v", err),
		}
	}

	createCouriers, err := packCreateCouriers(req)
	if err != nil {
		return err
	}
	out, err := c.service.CreateCouriers(ctx, createCouriers)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, unpackCouriers(out.Couriers))
}
