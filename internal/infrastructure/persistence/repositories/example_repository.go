package repositories

import (
	"fmt"

	"github.com/vitordm/go-boilerplate-webapi/internal/infrastructure/common"
)

type exampleRepository struct {
}

func NewExampleRepository() ExampleRepository {
	return &exampleRepository{}
}

func (repository *exampleRepository) ExampleMethodFromRepositoy() string {
	return fmt.Sprintf("Hello from ExampleRepository %s", common.RandomString(10))
}
