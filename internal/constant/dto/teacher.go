package dto

import (
	"fmt"
	"schoolcms/internal/constant/model/db"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type Teacher struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Title     string    `json:"title"`
	Status    db.Status `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (ts Teacher) ValidteTeachers() error {
	return validation.ValidateStruct(&ts,
		validation.Field(&ts.UserID, validation.By(func(value interface{}) error {
			if ts.UserID == uuid.Nil {
				return fmt.Errorf("invalid user id, user id required")
			}
			return nil
		})),
		validation.Field(&ts.Title, validation.Required.Error("title required")),
	)
}

type TeacherToSchool struct {
	ID        uuid.UUID `json:"id"`
	TeacherID uuid.UUID `json:"teacher_id"`
	SchoolID  uuid.UUID `json:"school_id"`
	Subject   string    `json:"subject"`
	Status    db.Status `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (t TeacherToSchool) VlidateAssignTeacher() error {
	return validation.ValidateStruct(&t, validation.Field(&t.TeacherID, validation.By(func(value interface{}) error {
		if t.TeacherID == uuid.Nil {
			return fmt.Errorf("teacher id required")
		}
		return nil
	})),
		validation.Field(&t.SchoolID, validation.By(func(value interface{}) error {
			if t.SchoolID == uuid.Nil {
				return fmt.Errorf("school id required")
			}
			return nil
		})), validation.Field(&t.Subject, validation.Required))
}
