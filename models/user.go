package models

import "time"

// User representa la estructura de un usuario
type User struct {
	ID                string    `json:"id"`
	Nombre            string    `json:"nombre" firestore:"nombre"`
	Apellidos         string    `json:"apellidos" firestore:"apellidos"`
	Email             string    `json:"email" firestore:"email"`
	Contrasena        string    `json:"contrasena" firestore:"contrasena"`
	FechaNacimiento   time.Time `json:"fecha_nacimiento" firestore:"fecha_nacimiento"`
	PreguntaSecreta   string    `json:"pregunta_secreta" firestore:"pregunta_secreta"`
	RespuestaSecreta  string    `json:"respuesta_secreta" firestore:"respuesta_secreta"`
}