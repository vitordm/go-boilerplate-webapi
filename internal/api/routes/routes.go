package routes

import (
	"cmp"
	"log/slog"
	"os"
	"slices"
	"strings"

	"github.com/labstack/echo/v4"
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

func OutputRoutes(e *Router) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Method", "Path"})

	routes := e.Routes()

	routesFiltered := funk.Filter(routes, func(x *echo.Route) bool {
		return !strings.HasSuffix(x.Path, "/*") && x.Method != echo.RouteNotFound
	}).([]*echo.Route)

	slices.SortFunc(routesFiltered, func(a, b *echo.Route) int {
		return cmp.Compare(a.Path, b.Path)
	})

	for i, route := range routesFiltered {

		t.AppendRows([]table.Row{
			{i + 1, route.Method, route.Path},
		})
		t.AppendSeparator()
	}

	t.Render()
}
