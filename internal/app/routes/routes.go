package routes

import (
	"github.concur.com/I573758/example-golang-webapi/internal/app/http"
	coreCache "github.concur.com/I573758/example-golang-webapi/internal/core/cache"
	"github.concur.com/I573758/example-golang-webapi/internal/core/ioc"
	"github.concur.com/I573758/example-golang-webapi/internal/core/server"
	"log/slog"
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
