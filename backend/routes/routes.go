package routes

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer &lt;token&gt;**
func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler, taskHandler *handlers.TaskHandler) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: true,
	}))

	api := app.Group("/api")
	auth := api.Group("/auth")

	// Public routes
	auth.Post("/register", userHandler.Register)
	auth.Post("/login", userHandler.Login)

	// Protected routes
	protected := auth.Group("/", middleware.AuthMiddleware())
	protected.Get("/profile", userHandler.GetProfile) // Fixed this line

	// Task routes (already protected correctly)
	task := api.Group("/task", middleware.AuthMiddleware())
	task.Get("/assigner", taskHandler.GetTasksByAssignerID)
	task.Post("/", taskHandler.CreateTask)
	task.Put("/:id", taskHandler.UpdateTask)
	task.Get("/:id", taskHandler.GetTask)
	task.Delete("/:id", taskHandler.DeleteTask)
}
