package data

import (
	"fmt"
	"github.concur.com/I573758/example-golang-webapi/internal/core/utils"
)

type ExampleRepository interface {
	ExampleMethodFromRepositoy() string
}

type exampleRepository struct {
}

func NewExampleRepository() ExampleRepository {
	return &exampleRepository{}
}

func (repository *exampleRepository) ExampleMethodFromRepositoy() string {
	return fmt.Sprintf("Hello from ExampleRepository %s", utils.RandomString(10))
}
