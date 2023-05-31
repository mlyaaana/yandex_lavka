package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"yandex-team.ru/bstask/api/order"
	"yandex-team.ru/bstask/domain"
	"yandex-team.ru/bstask/server/rest"
	timeutil "yandex-team.ru/bstask/util/time"
)

func (c *Controller) getAssignedOrders(e echo.Context) error {
	ctx := e.Request().Context()

	date, err := rest.QueryParamDate(e, "date", timeutil.NowDate())
	if err != nil {
		return &rest.Error{
			Code:    rest.ErrorCodeBadRequest,
			Message: fmt.Sprintf("failed to parse date: %v", err),
		}
	}
	courierId, err := rest.QueryParamInt(e, "courier_id", 0)
	if err != nil {
		return &rest.Error{
			Code:    rest.ErrorCodeBadRequest,
			Message: fmt.Sprintf("failed to parse courier id: %v", err),
		}
	}

	out, err := c.service.GetAssignedOrders(ctx, &order.GetAssignedOrdersIn{
		Date:      date,
		CourierId: domain.CourierId(courierId),
	})
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, unpackGetAssignedOrdersResponse(out))
}
