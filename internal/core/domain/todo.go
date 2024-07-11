package domain

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AuthId    uuid.UUID `json:"auth_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}
