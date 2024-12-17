package handlers

import (
	"backend/models"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskService services.TaskServiceInterface
}

func NewTaskHandler(taskService services.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

type TaskHandlerInterface interface {
	CreateTask(c *fiber.Ctx) error
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with the provided details
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Task object"
// @Success 201 {object} models.Task
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/task [post]
func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	createdTask, err := h.taskService.Create(&task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdTask)
}
