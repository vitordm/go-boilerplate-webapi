package app

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
	"github.com/vitordm/go-boilerplate-webapi/internal/app/helpers"
	"github.com/vitordm/go-boilerplate-webapi/internal/app/helpers/constants"
	di "github.com/vitordm/go-boilerplate-webapi/internal/app/helpers/ioc"
	"github.com/vitordm/go-boilerplate-webapi/internal/app/middlewares"
	"github.com/vitordm/go-boilerplate-webapi/internal/app/routes"
	coreCache "github.com/vitordm/go-boilerplate-webapi/internal/core/cache"
	"github.com/vitordm/go-boilerplate-webapi/internal/core/ioc"
	"github.com/vitordm/go-boilerplate-webapi/internal/core/server"
	"github.com/vitordm/go-boilerplate-webapi/internal/core/utils"
)

var container *ioc.ContainerDI
var cache *coreCache.Cache
var logger *slog.Logger

func Configure(e *echo.Echo) {

	configureRootPath()

	e.Validator = server.NewRequestValidator()

	container = ioc.NewContainerDI()
	cache = helpers.BuildCache()
	logger = helpers.BuildLogger()

	di.RegisterDependencies(container, logger)

	// Middleware
	e.Use(middlewares.EasterEggMiddleware())
	e.Use(middlewares.CorrelationId())
	e.Use(middlewares.Logger(logger))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	routes.DefineAllRoutes(e, container, cache, logger)

	if utils.IsDev() {
		e.Debug = true
		server.OutputRoutes(e)
	}
}

func Start(e *echo.Echo) {
	//setting PORT
	port := utils.GetEnvOrDefault(constants.APPLICATION_PORT_KEY,
		constants.APPLICATION_PORT_DEFAULT)

	e.Logger.Fatal(e.Start(port))
}

func configureRootPath() {
	exePath, err := os.Executable()
	if err != nil {
		os.Setenv(constants.ROOT_PATH_KEY, "/app")
		return
	}
	rootPath := filepath.Dir(exePath)

	os.Setenv(constants.ROOT_PATH_KEY, rootPath)
}
