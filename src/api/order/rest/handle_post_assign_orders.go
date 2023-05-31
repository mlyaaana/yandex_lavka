package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"yandex-team.ru/bstask/api/order"
	"yandex-team.ru/bstask/server/rest"
	timeutil "yandex-team.ru/bstask/util/time"
)

func (c *Controller) assignOrders(e echo.Context) error {
	ctx := e.Request().Context()

	date, err := rest.QueryParamDate(e, "date", timeutil.NowDate())
	if err != nil {
		return &rest.Error{
			Code:    rest.ErrorCodeBadRequest,
			Message: fmt.Sprintf("failed to parse date: %v", err),
		}
	}

	out, err := c.service.AssignOrders(ctx, &order.AssignOrdersIn{
		Date: date,
	})
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, unpackAssignOrdersResponse(out))
}
