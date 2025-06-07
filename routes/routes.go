package routes

import (
	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"task-manager-api/handlers"
	"task-manager-api/middleware"
)

// SetupRoutes configura las rutas de la API
func SetupRoutes(app *fiber.App, client *firestore.Client) {
	userHandler := handlers.NewUserHandler(client)
	taskHandler := handlers.NewTaskHandler(client)

	// Rutas de autenticación (sin middleware)
	app.Post("/api/register", userHandler.Register)
	app.Post("/api/login", userHandler.Login)
	app.Post("/api/recover-password", userHandler.RecoverPassword)

	// Rutas protegidas (con middleware)
	api := app.Group("/api", middleware.AuthMiddleware)
	api.Get("/users/me", userHandler.GetUser)
	api.Put("/users/me", userHandler.UpdateUser)
	api.Delete("/users/me", userHandler.DeleteUser)

	api.Post("/tasks", taskHandler.CreateTask)
	api.Get("/tasks", taskHandler.GetTasks)
	api.Put("/tasks/:id", taskHandler.UpdateTask)
	api.Delete("/tasks/:id", taskHandler.DeleteTask)
}