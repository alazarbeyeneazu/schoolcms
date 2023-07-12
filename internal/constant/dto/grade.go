package dto

import (
	"fmt"
	"schoolcms/internal/constant/model/db"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type Grade struct {
	ID        uuid.UUID `json:"id"`
	SchoolID  uuid.UUID `json:"school_id"`
	Status    db.Status `json:"status"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (g Grade) ValidateGrade() error {
	return validation.ValidateStruct(&g, validation.Field(&g.Name, validation.Required.Error("grade name required")),
		validation.Field(&g.SchoolID, validation.By(func(value interface{}) error {
			if g.SchoolID == uuid.Nil {
				return fmt.Errorf("school id should not be empty")
			}
			return nil
		},
		), validation.Required.Error("school_id required")),
	)
}
