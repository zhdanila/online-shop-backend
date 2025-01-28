package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"online-shop-backend/internal/app"
	"strings"
)

type EmptyResponse = struct{}

type Server struct {
	*echo.Echo
	port string
}

func (s *Server) GetPort() string {
	port := s.port
	if port == "" {
		return ":8080"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	return port
}

func NewServer(cnf *app.Config) *Server {
	middleware.DefaultRecoverConfig.LogErrorFunc = func(c echo.Context, err error, stack []byte) error {
		zap.L().Error(fmt.Sprintf("Server servers panic for host: %s; recovered: %v\n%s", c.Path(), err, string(stack)))
		return nil
	}

	server := echo.New()
	server.Use(middleware.RequestID())

	return &Server{
		server,
		cnf.GetPort(),
	}
}
