package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID `gorm:"primarykey" db:"id" json:"id" validate:"required,uuid"`
	Username  string    `db:"username" json:"username" validate:"required,string"`
	Password  string    `db:"password" json:"password" validate:"required,string"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
}
