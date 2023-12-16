package quote;

import (
	"time"
	"github.com/google/uuid"
)

type Quote struct {
	ID         uuid.UUID `json:"id,omitempty" db:"id"`
	Book       string    `json:"book,omitempty" db:"book"`
	Quote      string    `json:"quote,omitempty" db:"quote"`
	CreatedAt  time.Time `json:"insertedAt,omitempty" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty" db:"updated_at"`
}