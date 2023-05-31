package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"yandex-team.ru/bstask/api/order"
	"yandex-team.ru/bstask/domain"
	"yandex-team.ru/bstask/server/rest"
)

func (c *Controller) getOrder(e echo.Context) error {
	ctx := e.Request().Context()

	orderId, err := rest.PathParamInt64Required(e, "order_id")
	if err != nil {
		return &rest.Error{
			Code:    rest.ErrorCodeBadRequest,
			Message: fmt.Sprintf("failed to parse order id: %v", err),
		}
	}

	out, err := c.service.GetOrder(ctx, &order.GetOrderIn{
		OrderId: domain.OrderId(orderId),
	})
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, unpackOrder(out.Order))
}
