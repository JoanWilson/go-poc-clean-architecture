package domain

import "github.com/google/uuid"

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool  `json:"is_completed"`
}

func NewTask(title string, description string) *Task {
	return &Task{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		IsCompleted: false,
	}
}

type TaskRepository interface {
	Create(t *Task) error
	FindAll() ([]*Task, error)
}
