package entities

import (
	"time"

	"github.com/google/uuid"
)

type TodoModel struct {
	ID        uuid.UUID `gorm:"primarykey" db:"id" json:"id" validate:"required,uuid"`
	UserID    uuid.UUID `db:"user_id" json:"userId" validate:"required,uuid"`
	Completed bool      `db:"completed" json:"completed" validate:"required,bool"`
	Message   string    `db:"message" json:"message" validate:"required,string"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
