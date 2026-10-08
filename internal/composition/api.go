package composition

import (
	"log/slog"

	"log/slog"
	"os"
	"path/filepath"

	"github.com/vitordm/go-boilerplate-webapi/internal/app/data"
	"github.com/vitordm/go-boilerplate-webapi/internal/app/services"
	"github.com/vitordm/go-boilerplate-webapi/internal/infrastructure/di"

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

func registerDependencies(container *di.ContainerDI, logger *slog.Logger) {

	//repositories
	container.Provide(func() (data.ExampleRepository, error) {
		return data.NewExampleRepository(), nil
	})

	//services
	container.Provide(func(repository data.ExampleRepository) (services.ExampleService, error) {
		return services.NewExampleService(repository), nil
	})

}

type webApi struct {
	e    *echo.Echo
	port string
}

func (webApi *webApi) Run() error {
	return webApi.e.Start(webApi.port)
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

func NewApi() *webApi {
	e := echo.New()

	configureRootPath()

	e.Validator = server.NewRequestValidator()

	container = ioc.NewContainerDI()
	cache = helpers.BuildCache()
	logger = helpers.BuildLogger()

	registerDependencies(container, logger)

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

	port := utils.GetEnvOrDefault(constants.APPLICATION_PORT_KEY,
		constants.APPLICATION_PORT_DEFAULT)

	return &webApi{e, port}

}
