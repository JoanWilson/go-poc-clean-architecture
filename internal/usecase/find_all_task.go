package usecase

import "todoapp/internal/domain"

type FindAllOutputDto struct {
	ID          string `json:"id" example:"1" format:"string"`
	Title       string `json:"title" example:"title" format:"string"`
	Description string `json:"description" example:"description" format:"string"`
	IsCompleted bool   `json:"is_completed" example:"false" format:"bool"`
}

type FindAllTaskUseCase struct {
	repository domain.TaskRepository
}

func NewFindAllUseCase(r domain.TaskRepository) *FindAllTaskUseCase {
	return &FindAllTaskUseCase{repository: r}
}

func (u *FindAllTaskUseCase) Execute() ([]*FindAllOutputDto, error) {
	tasks, err := u.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var tasksOutput []*FindAllOutputDto
	for _, task := range tasks {
		tasksOutput = append(tasksOutput, &FindAllOutputDto{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			IsCompleted: task.IsCompleted,
		})
	}
	return tasksOutput, nil
}
