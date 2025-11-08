package server

import (
	"cmp"
	"github.com/labstack/echo/v4"
	"github.com/thoas/go-funk"
	"os"
	"slices"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Router = echo.Echo
type RequestContext = echo.Context
type Context = echo.Context
type HandlerFunc = echo.HandlerFunc

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
