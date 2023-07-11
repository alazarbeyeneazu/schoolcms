package dto

import (
	"schoolcms/internal/constant/model/db"
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ID       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"user_id"`
	Status   db.Status `json:"status"`
	CreatdAt time.Time `json:"created_at"`
}
