package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type Token struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	ActivatedAt time.Time `json:"activated_at"`
	ExpiredAt   time.Time `json:"expired_at"`
	Type        string    `json:"type"`
	Jwt         string    `json:"jwt"`
}
