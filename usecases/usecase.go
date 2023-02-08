package usecases

type todoUseCase struct {
	todoRepo domains.todoRepository
}

func NewToDoUseCase(repo domains.ToDoRepository) domains.ToDoUseCase {
	return &todoUseCase{
		todoRepo: repo,
	}
}
