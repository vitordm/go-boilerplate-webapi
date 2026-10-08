package routes

import (
	"cmp"
	"log/slog"
	"os"
	"slices"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/labstack/echo/v4"
	"github.com/thoas/go-funk"
	api "github.com/vitordm/go-boilerplate-webapi/internal/api/example"
	"github.com/vitordm/go-boilerplate-webapi/internal/api/server"
	"github.com/vitordm/go-boilerplate-webapi/internal/infrastructure/cache"
	"github.com/vitordm/go-boilerplate-webapi/internal/infrastructure/di"
)

func DefineAllRoutes(router *server.Router, container *di.ContainerDI, cache *cache.Cache, logger *slog.Logger) {

	router.GET("/ping", func(c server.Context) error {
		return c.String(200, "pong")
	})

	router.GET("/example", func(c server.Context) error {
		return api.GetExample(container, c)
	})

	router.POST("/example", func(c server.Context) error {
		return api.PostExample(container, c)
	})
}

func OutputRoutes(e *server.Router) {
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
