package models

import "time"

// Task represents a task in the system
// @Description Task object
type Task struct {
	ID          int64     `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Status      string    `db:"status" json:"status"`
	AssigneeID  *int64    `db:"assignee_id" json:"assignee_id"` // Add db tag
	AssignerID  *int64    `db:"assigner_id" json:"assigner_id"` // Add db tag
	Priority    int       `db:"priority" json:"priority"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type Priority int

const (
	PriorityLow    Priority = 1
	PriorityMedium Priority = 2
	PriorityHigh   Priority = 3
)
