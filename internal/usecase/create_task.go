package usecase

import "todoapp/internal/domain"

type CreateTaskInputDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTaskOutputDto struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

type CreateTaskUseCase struct {
	repository domain.TaskRepository
}

func NewCreateTaskUseCase(r domain.TaskRepository) CreateTaskUseCase {
	return CreateTaskUseCase{repository: r}
}

func (u *CreateTaskUseCase) Execute(input CreateTaskInputDto) (*CreateTaskOutputDto, error) {
	task := domain.NewTask(input.Title, input.Description)
	err := u.repository.Create(task)
	if err != nil {
		return nil, err
	}

	return &CreateTaskOutputDto{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
	}, nil
}
