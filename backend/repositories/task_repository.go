//go:generate mockgen -destination=mocks/mock_user_repository.go -package=mocks backend/repositories UserRepositoryInterface

package repositories

import (
	"backend/models"

	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db sqlx.DB
}

func NewTaskRepository(db sqlx.DB) TaskRepositoryInterface {
	return &TaskRepository{db: db}
}

// Interface
type TaskRepositoryInterface interface {
	Create(task *models.Task) (taskResponse *models.Task, err error)
}

func (r *TaskRepository) Create(task *models.Task) (taskResponse *models.Task, err error) {
	taskResponse = &models.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		AssigneeID:  task.AssigneeID,
		AssignerID:  task.AssignerID,
		Priority:    task.Priority,
	}

	err = r.db.QueryRow(`
        INSERT INTO tasks (title, description, status, assignee_id, assigner_id, priority, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
        RETURNING id
    `,
		task.Title,
		task.Description,
		task.Status,
		task.AssigneeID,
		task.AssignerID,
		task.Priority,
	).Scan(&taskResponse.ID)

	if err != nil {
		return nil, err
	}

	return taskResponse, nil
}
