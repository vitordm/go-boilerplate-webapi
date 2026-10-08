package repositories

import "github.com/vitordm/go-boilerplate-webapi/internal/contracts/persistence"

type todoRepository struct{}

func NewTodoRepository() persistence.TodoRepository {
	return &todoRepository{}
}
