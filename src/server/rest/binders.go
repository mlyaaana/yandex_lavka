package rest

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	timeutil "yandex-team.ru/bstask/util/time"
)

func QueryParamInt(c echo.Context, name string, value int) (int, error) {
	param := c.QueryParam(name)
	if param == "" {
		return value, nil
	}
	return strconv.Atoi(param)
}

func QueryParamDate(c echo.Context, name string, value time.Time) (time.Time, error) {
	param := c.QueryParam(name)
	if param == "" {
		return value, nil
	}
	return time.Parse(timeutil.DateLayout, param)
}

func QueryParamDateRequired(c echo.Context, name string) (time.Time, error) {
	param := c.QueryParam(name)
	if param == "" {
		return time.Time{}, fmt.Errorf("empty param %s", name)
	}
	return time.Parse(timeutil.DateLayout, param)
}

func PathParamInt64Required(c echo.Context, name string) (int64, error) {
	param := c.Param(name)
	return strconv.ParseInt(param, 10, 64)
}
