package example

import "time"

type ExampleRequest struct {
	ExampleField string `json:"example_field"`
}

type ExampleResponse struct {
	Date    time.Time `json:"date"`
	Message string    `json:"message"`
}
