package dto

import (
	"schoolcms/internal/constant/model/db"

	"github.com/google/uuid"
)

type Teacher struct {
	ID         uuid.UUID `json:"id"`
	UserId     uuid.UUID `json:"user_id"`
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name"`
	Title      string    `json:"title"`
	Status     db.Status `json:"status"`
}
