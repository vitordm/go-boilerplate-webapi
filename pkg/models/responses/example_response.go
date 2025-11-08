package responses

import "time"

type ExampleResponse struct {
	Date    time.Time `json:"date"`
	Message string    `json:"message"`
}
