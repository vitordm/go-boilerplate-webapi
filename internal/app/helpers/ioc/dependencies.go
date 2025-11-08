package ioc

import (
	"github.concur.com/I573758/example-golang-webapi/internal/app/data"
	"github.concur.com/I573758/example-golang-webapi/internal/app/services"
	"github.concur.com/I573758/example-golang-webapi/internal/core/ioc"
	"log/slog"
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
