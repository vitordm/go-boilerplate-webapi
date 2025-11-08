package http

import (
	"fmt"
	"github.concur.com/I573758/example-golang-webapi/internal/app/services"
	"github.concur.com/I573758/example-golang-webapi/internal/core/ioc"
	"github.concur.com/I573758/example-golang-webapi/internal/core/server"
	"github.concur.com/I573758/example-golang-webapi/pkg/models/requests"
	"github.concur.com/I573758/example-golang-webapi/pkg/models/responses"
	"time"
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
