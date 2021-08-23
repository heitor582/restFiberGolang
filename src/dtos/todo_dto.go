package dtos

import "github.com/google/uuid"

type TodoDto struct {
	Completed bool      `json:"completed"`
	Message   string    `json:"message"`
	UserID    uuid.UUID `json:"user_id"`
}
