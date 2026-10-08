package todo

type CreateTodoRequest struct {
	Name string `json:"name"`
}

type CreateTodoResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type TodoResponse struct {
	UUID  string   `json:"uuid"`
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

type CreateTodoItemRequest struct {
	Item string `json:"item"`
}
