package services

import (
	"backend/models"
	"backend/repositories"
	"errors"
)

type TaskService struct {
	taskRepository repositories.TaskRepositoryInterface
}

func NewTaskService(taskRepository repositories.TaskRepositoryInterface) TaskServiceInterface {
	return &TaskService{taskRepository: taskRepository}
}

type TaskServiceInterface interface {
	Create(task *models.Task) (taskResponse *models.Task, err error)
	Update(task *models.Task) (taskResponse *models.Task, err error)
	Get(id int64) (task *models.Task, err error)
	Delete(id int64) (err error)
	GetTasksByAssignerID(assignerID int64) (tasks []models.Task, err error)
}

func (s *TaskService) Create(task *models.Task) (taskResponse *models.Task, err error) {
	return s.taskRepository.Create(task)
}

func (s *TaskService) Update(task *models.Task) (taskResponse *models.Task, err error) {
	// Check if the task exists
	existingTask, err := s.taskRepository.Get(task.ID)
	if err != nil {
		return nil, err
	}

	if existingTask == nil {
		return nil, errors.New("task not found")
	}

	return s.taskRepository.Update(task)
}

func (s *TaskService) Get(id int64) (task *models.Task, err error) {
	return s.taskRepository.Get(id)
}

func (s *TaskService) Delete(id int64) (err error) {
	return s.taskRepository.Delete(id)
}

func (s *TaskService) GetTasksByAssignerID(assignerID int64) (tasks []models.Task, err error) {
	return s.taskRepository.GetTasksByAssignerID(assignerID)
}
