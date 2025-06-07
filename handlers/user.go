package handlers

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"task-manager-api/models"
	"task-manager-api/utils"
)

// UserHandler maneja las operaciones de usuarios
type UserHandler struct {
	client *firestore.Client
}

// NewUserHandler crea un nuevo manejador de usuarios
func NewUserHandler(client *firestore.Client) *UserHandler {
	return &UserHandler{client: client}
}

// Register registra un nuevo usuario
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No se pudo parsear el cuerpo"})
	}

	// Encriptar contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Contrasena), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al encriptar la contraseña"})
	}
	user.Contrasena = string(hashedPassword)
	user.ID = h.client.Collection("users").NewDoc().ID

	// Guardar en Firestore
	_, err = h.client.Collection("users").Doc(user.ID).Set(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al guardar el usuario"})
	}

	return c.JSON(user)
}

// Login autentica a un usuario y genera un token JWT
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var loginData struct {
		Email      string `json:"email"`
		Contrasena string `json:"contrasena"`
	}
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No se pudo parsear el cuerpo"})
	}

	// Buscar usuario en Firestore
	query := h.client.Collection("users").Where("email", "==", loginData.Email).Limit(1)
	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil || len(docs) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Credenciales inválidas"})
	}

	var user models.User
	docs[0].DataTo(&user)
	user.ID = docs[0].Ref.ID

	// Verificar contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(user.Contrasena), []byte(loginData.Contrasena)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Credenciales inválidas"})
	}

	// Generar token JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al generar el token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

// GetUser obtiene un usuario por ID
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	doc, err := h.client.Collection("users").Doc(userID).Get(context.Background())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Usuario no encontrado"})
	}

	var user models.User
	doc.DataTo(&user)
	user.ID = doc.Ref.ID
	return c.JSON(user)
}

// UpdateUser actualiza un usuario
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No se pudo parsear el cuerpo"})
	}

	// Actualizar en Firestore
	_, err := h.client.Collection("users").Doc(userID).Set(context.Background(), user, firestore.MergeAll)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al actualizar el usuario"})
	}

	return c.JSON(user)
}

// DeleteUser elimina un usuario
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	_, err := h.client.Collection("users").Doc(userID).Delete(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al eliminar el usuario"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// RecoverPassword recupera la contraseña usando pregunta secreta
func (h *UserHandler) RecoverPassword(c *fiber.Ctx) error {
	var recoverData struct {
		Email            string `json:"email"`
		PreguntaSecreta  string `json:"pregunta_secreta"`
		RespuestaSecreta string `json:"respuesta_secreta"`
		NuevaContrasena  string `json:"nueva_contrasena"`
	}
	if err := c.BodyParser(&recoverData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No se pudo parsear el cuerpo"})
	}

	// Buscar usuario
	query := h.client.Collection("users").Where("email", "==", recoverData.Email).Limit(1)
	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil || len(docs) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Usuario no encontrado"})
	}

	var user models.User
	docs[0].DataTo(&user)
	user.ID = docs[0].Ref.ID

	// Verificar pregunta y respuesta secreta
	if user.PreguntaSecreta != recoverData.PreguntaSecreta || user.RespuestaSecreta != recoverData.RespuestaSecreta {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Pregunta o respuesta secreta incorrecta"})
	}

	// Encriptar nueva contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(recoverData.NuevaContrasena), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al encriptar la contraseña"})
	}

	// Actualizar contraseña
	_, err = h.client.Collection("users").Doc(user.ID).Update(context.Background(), []firestore.Update{
		{Path: "contrasena", Value: string(hashedPassword)},
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al actualizar la contraseña"})
	}

	return c.JSON(fiber.Map{"message": "Contraseña actualizada correctamente"})
}