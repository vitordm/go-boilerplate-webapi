package routes

import (
	apiExample "github.com/vitordm/go-boilerplate-webapi/internal/api/example"
	"github.com/vitordm/go-boilerplate-webapi/internal/api/server"
)

func DefineAllRoutes(router *server.Router, handler *apiExample.Handler) {
	router.GET("/ping", func(c server.Context) error { return c.String(200, "pong") })
	router.GET("/example", handler.GetExample)
	router.POST("/example", handler.PostExample)
}
