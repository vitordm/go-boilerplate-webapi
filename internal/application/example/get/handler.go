package get

type GetExampleHandler interface {
	Get() string
}

type exampleHandler struct {
}

func NewExampleService() GetExampleHandler {
	return &exampleHandler{}
}

func (service *exampleHandler) Get() string {
	//return service.repository.ExampleMethodFromRepositoy()
	return ""
}
