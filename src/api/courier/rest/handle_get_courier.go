package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"yandex-team.ru/bstask/api/courier"
	"yandex-team.ru/bstask/domain"
	"yandex-team.ru/bstask/server/rest"
)

func (c *Controller) getCourier(e echo.Context) error {
	ctx := e.Request().Context()

	courierId, err := rest.PathParamInt64Required(e, "courier_id")
	if err != nil {
		return &rest.Error{
			Code:    rest.ErrorCodeBadRequest,
			Message: fmt.Sprintf("failed to parse courier id: %v", err),
		}
	}

	out, err := c.service.GetCourier(ctx, &courier.GetCourierIn{
		CourierId: domain.CourierId(courierId),
	})
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, unpackCourier(out.Courier))
}
