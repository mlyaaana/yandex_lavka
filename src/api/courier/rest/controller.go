package rest

import (
	"net/http"

	"yandex-team.ru/bstask/api/courier"
	"yandex-team.ru/bstask/server/rest"
)

type Controller struct {
	service courier.Service
}

type ControllerParams struct {
	Service courier.Service
}

func NewController(p *ControllerParams) *Controller {
	return &Controller{
		service: p.Service,
	}
}

func (c *Controller) Register(s *rest.Server) {
	s.AddRoute(&rest.Route{
		Method:  http.MethodGet,
		Path:    "/couriers",
		Handler: c.getCouriers,
	})
	s.AddRoute(&rest.Route{
		Method:  http.MethodPost,
		Path:    "/couriers",
		Handler: c.postCouriers,
	})
	s.AddRoute(&rest.Route{
		Method:  http.MethodGet,
		Path:    "/couriers/:courier_id",
		Handler: c.getCourier,
	})
	s.AddRoute(&rest.Route{
		Method:  http.MethodGet,
		Path:    "/couriers/meta-info/:courier_id",
		Handler: c.getCourierMetaInfo,
	})
}
