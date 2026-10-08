package server

import (
	"cmp"
	"os"
	"slices"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/labstack/echo/v4"
	"github.com/thoas/go-funk"
)

type Router = echo.Echo
type RequestContext = echo.Context
type Context = echo.Context
type HandlerFunc = echo.HandlerFunc

func OutputRoutes(e *Router) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Method", "Path"})
	routes := funk.Filter(e.Routes(), func(x *echo.Route) bool {
		return !strings.HasSuffix(x.Path, "/*") && x.Method != echo.RouteNotFound
	}).([]*echo.Route)
	slices.SortFunc(routes, func(a, b *echo.Route) int { return cmp.Compare(a.Path, b.Path) })
	for i, route := range routes {
		t.AppendRow(table.Row{i + 1, route.Method, route.Path})
	}
	t.Render()
}
