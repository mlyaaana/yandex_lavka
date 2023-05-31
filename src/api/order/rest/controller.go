package rest

import (
	"net/http"

	"yandex-team.ru/bstask/api/order"
	"yandex-team.ru/bstask/server/rest"
)

type Controller struct {
	service order.Service
}

type ControllerParams struct {
	Service order.Service
}

func NewController(p *ControllerParams) *Controller {
	return &Controller{
		service: p.Service,
	}
}

func (c *Controller) Register(s *rest.Server) {
	s.AddRoute(&rest.Route{
		Method:  http.MethodGet,
		Path:    "/orders",
		Handler: c.getOrders,
	})
	s.AddRoute(&rest.Route{
		Method:  http.MethodPost,
		Path:    "/orders",
		Handler: c.postOrders,
	})
	s.AddRoute(&rest.Route{
		Method:  http.MethodGet,
		Path:    "/orders/:order_id",
		Handler: c.getOrder,
	})
	s.AddRoute(&rest.Route{
		Method:  http.MethodPost,
		Path:    "/orders/complete",
		Handler: c.completeOrders,
	})
	s.AddRoute(&rest.Route{
		Method:  http.MethodPost,
		Path:    "/orders/assign",
		Handler: c.assignOrders,
	})
	s.AddRoute(&rest.Route{
		Method:  http.MethodGet,
		Path:    "/couriers/assignments",
		Handler: c.getAssignedOrders,
	})
}
