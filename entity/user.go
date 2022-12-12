package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
