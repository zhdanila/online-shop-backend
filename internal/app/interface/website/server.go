package website

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"net/http"
	"online-shop-backend/internal/app"
	"online-shop-backend/internal/app/interface/http/middleware"
	"online-shop-backend/internal/app/interface/website/handler"
	server2 "online-shop-backend/pkg/server"
)

func NewServer(cnf *app.Config) *server2.Server {
	server := server2.NewServer(cnf)
	server.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"service": "image-optimization-api",
			"message": "Root page do nothing. Please go somewhere else",
		})
	})

	server.Use(middleware.SetCORS())
	server.Use(middleware.Recover())

	nameGroup := server.Group("/api")
	handler.NewImage(
		do.MustInvoke[*image.Service](inj),
	).Register(nameGroup)

	server.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, `{"status":"OK"}`)
	})

	return server
}
