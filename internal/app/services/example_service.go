package services

import "github.com/vitordm/go-boilerplate-webapi/internal/app/data"

type ExampleService interface {
	ExampleMethodFromService() string
}

type exampleService struct {
	repository data.ExampleRepository
}

func NewExampleService(repository data.ExampleRepository) ExampleService {
	return &exampleService{
		repository: repository,
	}
}

func (service *exampleService) ExampleMethodFromService() string {
	return service.repository.ExampleMethodFromRepositoy()
}
