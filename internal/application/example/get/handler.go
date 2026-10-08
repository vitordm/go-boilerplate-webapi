package get

import "github.com/vitordm/go-boilerplate-webapi/internal/contracts/persistence"

type Handler struct{ repository persistence.ExampleRepository }

func NewHandler(repository persistence.ExampleRepository) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) Get() string { return h.repository.ExampleMethodFromRepository() }
