package entities

import (
	"time"

	"github.com/google/uuid"
)

type TodoModel struct {
	ID        uuid.UUID `gorm:"primarykey" db:"id" json:"id" validate:"required,uuid"`
	UserID    uuid.UUID `db:"userId" json:"userId" validate:"required,uuid"`
	Completed bool      `db:"completed" json:"completed" validate:"required,bool"`
	Message   string    `db:"message" json:"message" validate:"required,string"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
}
