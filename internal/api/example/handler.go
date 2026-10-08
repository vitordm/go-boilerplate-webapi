package http

import (
	"fmt"
	"time"

	. "github.com/vitordm/go-boilerplate-webapi/internal/api/responses"
	"github.com/vitordm/go-boilerplate-webapi/internal/api/server"
	"github.com/vitordm/go-boilerplate-webapi/internal/application/example/get"
	examplecontract "github.com/vitordm/go-boilerplate-webapi/pkg/contracts/example"
)

type Handler struct{ service *get.Handler }

func NewHandler(service *get.Handler) *Handler { return &Handler{service: service} }

func (h *Handler) GetExample(c server.Context) error {
	return Ok(c, h.service.Get())
}

func (h *Handler) PostExample(c server.Context) error {
	request := new(examplecontract.ExampleRequest)
	if err := c.Bind(request); err != nil {
		return BadRequest(c, err)
	}

	response := new(examplecontract.ExampleResponse)
	response.Message = fmt.Sprintf("%s - %s", request.ExampleField, h.service.Get())
	response.Date = time.Now()
	return Ok(c, response)
}
