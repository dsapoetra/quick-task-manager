package models

import "time"

// Task represents a task in the system
// @Description Task object
type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	AssigneeID  *int64    `json:"assignee_id,omitempty"` // Make pointer for nullable
	AssignerID  *int64    `json:"assigner_id,omitempty"` // Make pointer for nullable
	Priority    Priority  `json:"priority"`              // Make sure this tag is correct
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Priority int

const (
	PriorityLow    Priority = 1
	PriorityMedium Priority = 2
	PriorityHigh   Priority = 3
)
