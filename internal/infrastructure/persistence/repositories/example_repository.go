package repositories

import (
	"fmt"

	"github.com/vitordm/go-boilerplate-webapi/internal/contracts/persistence"
	"github.com/vitordm/go-boilerplate-webapi/internal/shared/utils"
)

type exampleRepository struct {
}

func NewExampleRepository() persistence.ExampleRepository {
	return &exampleRepository{}
}

func (repository *exampleRepository) ExampleMethodFromRepository() string {
	return fmt.Sprintf("Hello from ExampleRepository %s", utils.RandomString(10))
}
