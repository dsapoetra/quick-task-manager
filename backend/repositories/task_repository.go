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
	Update(task *models.Task) (taskResponse *models.Task, err error)
	Get(id int64) (task *models.Task, err error)
	Delete(id int64) (err error)
	GetTasksByAssignerID(assignerID int64) ([]models.Task, error)
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

func (r *TaskRepository) Update(task *models.Task) (taskResponse *models.Task, err error) {
	taskResponse = &models.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		AssigneeID:  task.AssigneeID,
		AssignerID:  task.AssignerID,
		Priority:    task.Priority,
	}

	err = r.db.QueryRow(`
		UPDATE tasks SET title = $1, description = $2, status = $3, assignee_id = $4, assigner_id = $5, priority = $6, updated_at = NOW()
		WHERE id = $7
		RETURNING id
	`,
		task.Title,
		task.Description,
		task.Status,
		task.AssigneeID,
		task.AssignerID,
		task.Priority,
		task.ID,
	).Scan(&taskResponse.ID)

	if err != nil {
		return nil, err
	}

	return taskResponse, nil
}

func (r *TaskRepository) Get(id int64) (task *models.Task, err error) {
	// Select the task from the database but not using *
	task = &models.Task{}
	err = r.db.QueryRow(`
		SELECT id, title, description, status, assignee_id, assigner_id, priority, created_at, updated_at
		FROM tasks
		WHERE id = $1
	`, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.AssigneeID, &task.AssignerID, &task.Priority, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) Delete(id int64) (err error) {
	_, err = r.db.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	return err
}

func (r *TaskRepository) GetTasksByAssignerID(assignerID int64) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Select(&tasks, "SELECT id, title, description, status, assignee_id, assigner_id, priority, created_at, updated_at FROM tasks WHERE assigner_id = $1", assignerID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
