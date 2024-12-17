package main

import (
	"log"

	"backend/config"
	"backend/db"
	"backend/handlers"
	"backend/repositories"
	"backend/routes"
	"backend/services"

	_ "backend/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Task Management API
// @version 1.0.0
// @description API for managing tasks
// @host localhost:8080
// @BasePath /
func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	database, err := db.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(*database)
	taskRepo := repositories.NewTaskRepository(*database)

	// Initialize services
	userService := services.NewUserService(userRepo)
	taskService := services.NewTaskService(taskRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app, userHandler, taskHandler)

	// Setup Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
