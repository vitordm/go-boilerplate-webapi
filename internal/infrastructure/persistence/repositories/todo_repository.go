package repositories

type todoRepository struct {
}

func NewTodoRepository() TodoRepository {
	return &exampleRepository{}
}
