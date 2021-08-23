package dtos

import "github.com/google/uuid"

type TodoDto struct {
	Completed bool      `json:"completed" validate:"required,bool"`
	Message   string    `json:"message" validate:"required,string"`
	UserID    uuid.UUID `json:"user_id" validate:"required,uuid"`
}
