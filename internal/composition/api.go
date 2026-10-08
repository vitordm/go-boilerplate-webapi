package composition

import (
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	apiExample "github.com/vitordm/go-boilerplate-webapi/internal/api/example"
	apiMiddleware "github.com/vitordm/go-boilerplate-webapi/internal/api/middleware"
	"github.com/vitordm/go-boilerplate-webapi/internal/api/routes"
	"github.com/vitordm/go-boilerplate-webapi/internal/api/server"
	"github.com/vitordm/go-boilerplate-webapi/internal/application/example/get"
	"github.com/vitordm/go-boilerplate-webapi/internal/contracts/persistence"
	"github.com/vitordm/go-boilerplate-webapi/internal/infrastructure/logging"
	"github.com/vitordm/go-boilerplate-webapi/internal/infrastructure/persistence/repositories"
	"github.com/vitordm/go-boilerplate-webapi/internal/shared/constants"
	"github.com/vitordm/go-boilerplate-webapi/internal/shared/utils"
)

type webAPI struct {
	e    *echo.Echo
	port string
}

func (app *webAPI) Run() error { return app.e.Start(app.port) }

func configureRootPath() {
	exePath, err := os.Executable()
	if err != nil {
		_ = os.Setenv(constants.ROOT_PATH_KEY, "/app")
		return
	}
	_ = os.Setenv(constants.ROOT_PATH_KEY, filepath.Dir(exePath))
}

func buildExampleHandler() *apiExample.Handler {
	var repository persistence.ExampleRepository = repositories.NewExampleRepository()
	return apiExample.NewHandler(get.NewHandler(repository))
}

func NewAPI() *webAPI {
	configureRootPath()
	e := echo.New()
	e.Validator = server.NewRequestValidator()
	logger := logging.BuildLogger()
	e.Use(apiMiddleware.EasterEggMiddleware())
	e.Use(apiMiddleware.CorrelationId())
	e.Use(apiMiddleware.Logger(logger))
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())
	routes.DefineAllRoutes(e, buildExampleHandler())
	if utils.IsDev() {
		e.Debug = true
		server.OutputRoutes(e)
	}
	port := utils.GetEnvOrDefault(constants.APPLICATION_PORT_KEY, constants.APPLICATION_PORT_DEFAULT)
	return &webAPI{e: e, port: port}
}

func NewApi() *webAPI { return NewAPI() }
