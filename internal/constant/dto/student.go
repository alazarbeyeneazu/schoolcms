package dto

import (
	"fmt"
	"schoolcms/internal/constant/model/db"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type Student struct {
	ID       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"user_id"`
	Status   db.Status `json:"status"`
	CreatdAt time.Time `json:"created_at"`
}

func (s *Student) ValidateStudent() error {
	return validation.ValidateStruct(&s, validation.Field(&s.UserId, validation.By(func(value interface{}) error {
		if s.UserId == uuid.Nil {
			return fmt.Errorf("user id required")
		}
		return nil
	})))
}
