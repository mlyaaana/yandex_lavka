package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"yandex-team.ru/bstask/api/order"
	"yandex-team.ru/bstask/gen"
	"yandex-team.ru/bstask/server/rest"
)

func (c *Controller) completeOrders(e echo.Context) error {
	ctx := e.Request().Context()

	req := &gen.CompleteOrderRequestDto{}
	if err := e.Bind(req); err != nil {
		return &rest.Error{
			Code:    rest.ErrorCodeBadRequest,
			Message: fmt.Sprintf("failed to parse request: %v", err),
		}
	}

	out, err := c.service.CompleteOrders(ctx, &order.CompleteOrdersIn{
		CompleteOrders: packCompleteOrders(req.CompleteInfo),
	})
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, unpackOrders(out.Orders))
}
