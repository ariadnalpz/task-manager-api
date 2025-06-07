package models

import "time"

// Task representa la estructura de una tarea
type Task struct {
	ID          string    `json:"id"`
	Titulo      string    `json:"titulo" firestore:"titulo"`
	Descripcion string    `json:"descripcion" firestore:"descripcion"`
	FechaInicio time.Time `json:"fecha_inicio" firestore:"fecha_inicio"`
	Deadline    time.Time `json:"deadline" firestore:"deadline"`
	UsuarioID   string    `json:"usuario_id" firestore:"usuario_id"`
}