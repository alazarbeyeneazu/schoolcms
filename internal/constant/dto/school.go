package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type School struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Log  string    `json:"logo"`
}

func (s School) ValidateSchool() error {

	return validation.ValidateStruct(&s,
		validation.Field(&s.Name, validation.Required.Error("school name required")),
	)
}
