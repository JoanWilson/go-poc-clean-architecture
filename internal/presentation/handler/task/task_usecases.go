package task

import "todoapp/internal/usecase"

type UseCases struct {
	Create  usecase.CreateTaskUseCase
	FindAll usecase.FindAllTaskUseCase
}

func NewUseCases(c usecase.CreateTaskUseCase, f usecase.FindAllTaskUseCase) UseCases {
	return UseCases{
		Create:  c,
		FindAll: f,
	}
}
