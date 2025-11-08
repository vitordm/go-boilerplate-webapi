package routes

import (
	"log/slog"

	"github.com/vitordm/go-boilerplate-webapi/internal/app/http"
	coreCache "github.com/vitordm/go-boilerplate-webapi/internal/core/cache"
	"github.com/vitordm/go-boilerplate-webapi/internal/core/ioc"
	"github.com/vitordm/go-boilerplate-webapi/internal/core/server"
)

func DefineAllRoutes(router *server.Router, container *ioc.ContainerDI, cache *coreCache.Cache, logger *slog.Logger) {

	router.GET("/ping", func(c server.Context) error {
		return c.String(200, "pong")
	})

	router.GET("/example", func(c server.Context) error {
		return http.GetExample(container, c)
	})

	router.POST("/example", func(c server.Context) error {
		return http.PostExample(container, c)
	})
}
