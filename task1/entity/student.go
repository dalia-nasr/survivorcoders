package entity

import "time"

type Student struct {
	Id        int       `json:"id"`
	FirstName string    `json:"FirstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
