package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"task-manager-api/config"
	"task-manager-api/handlers"
	"task-manager-api/models"
)

func TestRegister(t *testing.T) {
	client, err := config.FirestoreClient()
	if err != nil {
		t.Fatalf("Error conectando a Firestore: %v", err)
	}

	app := fiber.New()
	userHandler := handlers.NewUserHandler(client)
	app.Post("/api/register", userHandler.Register)

	user := models.User{
		Nombre:           "Juan",
		Apellidos:        "Pérez",
		Email:            "juan@example.com",
		Contrasena:       "password123",
		PreguntaSecreta:  "Color favorito",
		RespuestaSecreta: "Azul",
	}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error en la solicitud: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado status 200, recibido %d", resp.StatusCode)
	}
}