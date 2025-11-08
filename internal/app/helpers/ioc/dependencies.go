package ioc

import (
	"log/slog"

	"github.com/vitordm/go-boilerplate-webapi/internal/app/data"
	"github.com/vitordm/go-boilerplate-webapi/internal/app/services"
	"github.com/vitordm/go-boilerplate-webapi/internal/core/ioc"
)

func RegisterDependencies(container *ioc.ContainerDI, logger *slog.Logger) {

	//repositories
	container.Provide(func() (data.ExampleRepository, error) {
		return data.NewExampleRepository(), nil
	})

	//services
	container.Provide(func(repository data.ExampleRepository) (services.ExampleService, error) {
		return services.NewExampleService(repository), nil
	})

}
