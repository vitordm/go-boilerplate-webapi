package app

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4/middleware"
	"github.concur.com/I573758/example-golang-webapi/internal/app/helpers"
	"github.concur.com/I573758/example-golang-webapi/internal/app/helpers/constants"
	di "github.concur.com/I573758/example-golang-webapi/internal/app/helpers/ioc"
	"github.concur.com/I573758/example-golang-webapi/internal/app/middlewares"
	"github.concur.com/I573758/example-golang-webapi/internal/app/routes"
	coreCache "github.concur.com/I573758/example-golang-webapi/internal/core/cache"
	"github.concur.com/I573758/example-golang-webapi/internal/core/ioc"
	"github.concur.com/I573758/example-golang-webapi/internal/core/server"
	"github.concur.com/I573758/example-golang-webapi/internal/core/utils"
)

var container *ioc.ContainerDI
var cache *coreCache.Cache
var logger *slog.Logger

func Configure(e *echo.Echo) {

	configureRootPath()

	e.Validator = server.NewRequestValidator()

	container = ioc.NewContainerDI()
	cache = helpers.BuildCache()
	logger = helpers.NewSLogJsonCommandLine()
	//logger = helpers.NewSLogJsonCommandLineAndFile()

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
