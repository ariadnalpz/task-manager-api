package handlers

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"task-manager-api/models"
)

type TaskHandler struct {
	client *firestore.Client
}

func NewTaskHandler(client *firestore.Client) *TaskHandler {
	return &TaskHandler{client: client}
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No se pudo parsear el cuerpo"})
	}

	task.ID = h.client.Collection("tasks").NewDoc().ID
	task.UsuarioID = userID

	fmt.Printf("Creando tarea con ID %s y usuario_id %s\n", task.ID, task.UsuarioID)

	_, err := h.client.Collection("tasks").Doc(task.ID).Set(context.Background(), task)
	if err != nil {
		fmt.Printf("Error al crear tarea: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al crear la tarea"})
	}

	return c.JSON(task)
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	fmt.Printf("Buscando tareas para usuario %s\n", userID)
	docs, err := h.client.Collection("tasks").Where("usuario_id", "==", userID).Documents(context.Background()).GetAll()
	if err != nil {
		fmt.Printf("Error al obtener tareas: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al obtener las tareas"})
	}

	var tasks []models.Task
	for _, doc := range docs {
		var task models.Task
		if err := doc.DataTo(&task); err != nil {
			fmt.Printf("Error al deserializar tarea %s: %v\n", doc.Ref.ID, err)
			continue
		}
		task.ID = doc.Ref.ID
		fmt.Printf("Tarea encontrada: ID %s, usuario_id %s\n", task.ID, task.UsuarioID)
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	taskID := c.Params("id")
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No se pudo parsear el cuerpo"})
	}

	fmt.Printf("Intentando actualizar tarea con ID %s para usuario %s\n", taskID, userID)

	// Verificar que la tarea existe
	doc, err := h.client.Collection("tasks").Doc(taskID).Get(context.Background())
	if err != nil {
		fmt.Printf("Error al buscar tarea %s: %v\n", taskID, err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tarea no encontrada"})
	}

	var existingTask models.Task
	if err := doc.DataTo(&existingTask); err != nil {
		fmt.Printf("Error al deserializar tarea %s: %v\n", taskID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al leer los datos de la tarea"})
	}

	fmt.Printf("Tarea encontrada en Firestore: ID %s, usuario_id %s\n", doc.Ref.ID, existingTask.UsuarioID)

	// Verificar que la tarea pertenece al usuario
	if existingTask.UsuarioID != userID {
		fmt.Printf("Usuario no autorizado: tarea usuario_id %s, usuario actual %s\n", existingTask.UsuarioID, userID)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "No autorizado"})
	}

	// Convertir la estructura task a un mapa para usar MergeAll
	updateData := map[string]interface{}{
		"titulo":       task.Titulo,
		"descripcion":  task.Descripcion,
		"fecha_inicio": task.FechaInicio,
		"deadline":     task.Deadline,
	}

	// Actualizar en Firestore con MergeAll
	_, err = h.client.Collection("tasks").Doc(taskID).Set(context.Background(), updateData, firestore.MergeAll)
	if err != nil {
		fmt.Printf("Error al actualizar tarea %s: %v\n", taskID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al actualizar la tarea"})
	}

	return c.JSON(task)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	taskID := c.Params("id")

	doc, err := h.client.Collection("tasks").Doc(taskID).Get(context.Background())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tarea no encontrada"})
	}
	var existingTask models.Task
	doc.DataTo(&existingTask)
	if existingTask.UsuarioID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "No autorizado"})
	}

	_, err = h.client.Collection("tasks").Doc(taskID).Delete(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al eliminar la tarea"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}