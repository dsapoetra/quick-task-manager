package routes

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gofiber/fiber/v2"
)

// @title Task Manager API
// @version 1.0
// @description This is a task management server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler, taskHandler *handlers.TaskHandler) {
	api := app.Group("/api")

	// Public routes
	auth := api.Group("/auth")
	auth.Post("/register", userHandler.Register)
	auth.Post("/login", userHandler.Login)

	task := api.Group("/task", middleware.AuthMiddleware())
	task.Post("/", taskHandler.CreateTask, middleware.AuthMiddleware())
	task.Put("/:id", taskHandler.UpdateTask, middleware.AuthMiddleware())
	task.Get("/:id", taskHandler.GetTask, middleware.AuthMiddleware())
}
