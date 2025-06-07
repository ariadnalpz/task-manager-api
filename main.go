package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"task-manager-api/config"
	"task-manager-api/routes"
)

func main() {
	// Inicializar Firestore
	client, err := config.FirestoreClient()
	if err != nil {
		log.Fatalf("Error conectando a Firestore: %v", err)
	}
	defer client.Close()

	// Inicializar Fiber
	app := fiber.New()

	// Configurar rutas
	routes.SetupRoutes(app, client)

	// Iniciar el servidor
	log.Fatal(app.Listen(":3000"))
}