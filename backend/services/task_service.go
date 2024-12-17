package services

import (
	"backend/models"
	"backend/repositories"
)

type TaskService struct {
	taskRepository repositories.TaskRepositoryInterface
}

func NewTaskService(taskRepository repositories.TaskRepositoryInterface) TaskServiceInterface {
	return &TaskService{taskRepository: taskRepository}
}

type TaskServiceInterface interface {
	Create(task *models.Task) (taskResponse *models.Task, err error)
}

func (s *TaskService) Create(task *models.Task) (taskResponse *models.Task, err error) {
	return s.taskRepository.Create(task)
}
