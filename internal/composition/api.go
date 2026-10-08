package composition

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	apiMiddleware "github.com/vitordm/go-boilerplate-webapi/internal/api/middleware"
	"github.com/vitordm/go-boilerplate-webapi/internal/api/routes"
	"github.com/vitordm/go-boilerplate-webapi/internal/api/server"
	"github.com/vitordm/go-boilerplate-webapi/internal/application/example/get"
	"github.com/vitordm/go-boilerplate-webapi/internal/contracts/persistence"
	infraCache "github.com/vitordm/go-boilerplate-webapi/internal/infrastructure/cache"
	"github.com/vitordm/go-boilerplate-webapi/internal/infrastructure/di"
	"github.com/vitordm/go-boilerplate-webapi/internal/infrastructure/persistence/repositories"
	"github.com/vitordm/go-boilerplate-webapi/internal/shared"
	"github.com/vitordm/go-boilerplate-webapi/internal/shared/constants"
	"github.com/vitordm/go-boilerplate-webapi/internal/shared/utils"
)

var container *di.ContainerDI
var cache *infraCache.Cache
var logger *slog.Logger

func registerDependencies(container *di.ContainerDI, logger *slog.Logger) {

	//repositories
	container.Provide(func() (persistence.ExampleRepository, error) {
		return repositories.NewExampleRepository(), nil
	})

	//services
	container.Provide(func() get.GetExampleHandler {
		return get.NewExampleService()
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

	container = di.NewContainerDI()
	cache = shared.BuildCache()
	logger = shared.BuildLogger()

	registerDependencies(container, logger)

	// Middleware
	e.Use(apiMiddleware.EasterEggMiddleware())
	e.Use(apiMiddleware.CorrelationId())
	e.Use(apiMiddleware.Logger(logger))
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())
	routes.DefineAllRoutes(e, container, cache, logger)

	if utils.IsDev() {
		e.Debug = true
		routes.OutputRoutes(e)
	}

	port := utils.GetEnvOrDefault(constants.APPLICATION_PORT_KEY,
		constants.APPLICATION_PORT_DEFAULT)

	return &webApi{e, port}

}

func NewAPI() *webApi {
	return NewApi()
}
