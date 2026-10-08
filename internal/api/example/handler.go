package http

import (
	"fmt"
	"time"

	. "github.com/vitordm/go-boilerplate-webapi/internal/api/responses"
	"github.com/vitordm/go-boilerplate-webapi/internal/api/server"
	"github.com/vitordm/go-boilerplate-webapi/internal/application/example/get"
	"github.com/vitordm/go-boilerplate-webapi/internal/infrastructure/di"
	examplecontract "github.com/vitordm/go-boilerplate-webapi/pkg/contracts/example"
)

func GetExample(container *di.ContainerDI, c server.Context) error {
	return container.Invoke(func(service get.GetExampleHandler) error {
		response := service.Get()
		return Ok(c, response)
	})
}

func PostExample(container *di.ContainerDI, c server.Context) error {
	return container.Invoke(func(service get.GetExampleHandler) error {
		request := new(examplecontract.ExampleRequest)
		if err := c.Bind(request); err != nil {
			return BadRequest(c, err)
		}

		message := service.Get()

		response := new(examplecontract.ExampleResponse)
		response.Message = fmt.Sprintf("%s - %s", request.ExampleField, message)
		response.Date = time.Now()
		return Ok(c, response)
	})
}
