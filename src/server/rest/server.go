package rest

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo    *echo.Echo
	address string
}

type ServerParams struct {
	Address string
}

func NewServer(p *ServerParams) *Server {
	s := &Server{
		address: p.Address,
	}
	s.initEcho()
	return s
}

func (s *Server) Start() error {
	return s.echo.Start(s.address)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

func (s *Server) AddRoute(r *Route) {
	s.echo.Add(r.Method, r.Path, r.Handler, rateLimitMiddleware())
}

func (s *Server) initEcho() {
	s.echo = echo.New()
	s.echo.HidePort = true
	s.echo.HideBanner = true
	s.echo.Use(
		middleware.Logger(),
		middleware.CORS(),
		errorMiddleware,
	)
}
