package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"yandex-team.ru/bstask/api/courier"
	"yandex-team.ru/bstask/domain"
	"yandex-team.ru/bstask/server/rest"
)

func (c *Controller) getCourierMetaInfo(e echo.Context) error {
	ctx := e.Request().Context()

	courierId, err := rest.PathParamInt64Required(e, "courier_id")
	if err != nil {
		return &rest.Error{
			Code:    rest.ErrorCodeBadRequest,
			Message: fmt.Sprintf("failed to parse courier id: %v", err),
		}
	}
	startDate, err := rest.QueryParamDateRequired(e, "startDate")
	if err != nil {
		return &rest.Error{
			Code:    rest.ErrorCodeBadRequest,
			Message: fmt.Sprintf("failed to parse start date: %v", err),
		}
	}
	endDate, err := rest.QueryParamDateRequired(e, "endDate")
	if err != nil {
		return &rest.Error{
			Code:    rest.ErrorCodeBadRequest,
			Message: fmt.Sprintf("failed to parse end date: %v", err),
		}
	}

	out, err := c.service.GetCourierMetaInfo(ctx, &courier.GetCourierMetaInfoIn{
		CourierId: domain.CourierId(courierId),
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, unpackCourierMetaInfo(out))
}
