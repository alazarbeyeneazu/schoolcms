package dto

import (
	"schoolcms/internal/constant/model/db"
	"schoolcms/platform/utils"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type Family struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	Status     db.Status `json:"status"`
	FamilyType string    `json:"family_type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

func (f Family) ValidateFamily() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.UserID, validation.By(utils.CheckForNullUUID("user_id required"))),
		validation.Field(&f.FamilyType, validation.Required.Error("family_type required")),
	)
}

type FamilyToStudent struct {
	ID         uuid.UUID `json:"id"`
	StudentID  uuid.UUID `json:"student_id"`
	FamilyID   uuid.UUID `json:"family_id"`
	Status     db.Status `json:"status"`
	FamilyType string    `json:"family_type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

func (f FamilyToStudent) ValidateFamilyToStudent() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.StudentID, validation.By(utils.CheckForNullUUID("student_id required"))),
		validation.Field(&f.FamilyID, validation.By(utils.CheckForNullUUID("family_id required"))),
		validation.Field(&f.FamilyType, validation.Required.Error("family_type required")),
	)
}
