package http

import (
	"fmt"
	"time"

	"github.com/vitordm/go-boilerplate-webapi/internal/app/services"
	"github.com/vitordm/go-boilerplate-webapi/internal/core/ioc"
	"github.com/vitordm/go-boilerplate-webapi/internal/core/server"
	"github.com/vitordm/go-boilerplate-webapi/pkg/models/requests"
	"github.com/vitordm/go-boilerplate-webapi/pkg/models/responses"
)

func GetExample(container *ioc.ContainerDI, c server.Context) error {
	return container.Invoke(func(service services.ExampleService) error {
		response := service.ExampleMethodFromService()
		return Ok(c, response)
	})
}

func PostExample(container *ioc.ContainerDI, c server.Context) error {
	return container.Invoke(func(service services.ExampleService) error {
		request := new(requests.ExampleRequest)
		if err := c.Bind(request); err != nil {
			return BadRequest(c, err)
		}

		message := service.ExampleMethodFromService()

		response := new(responses.ExampleResponse)
		response.Message = fmt.Sprintf("%s - %s", request.ExampleField, message)
		response.Date = time.Now()
		return Ok(c, response)
	})
}
