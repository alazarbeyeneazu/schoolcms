package dto

import (
	"fmt"
	"schoolcms/internal/constant/model/db"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type Teacher struct {
	ID     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`
	Title  string    `json:"title"`
	Status db.Status `json:"status"`
}

func (ts Teacher) ValidteTeachers() error {
	return validation.ValidateStruct(&ts,
		validation.Field(&ts.UserId, validation.By(func(value interface{}) error {
			if ts.UserId == uuid.Nil {
				return fmt.Errorf("invalid user id, user id required")
			}
			return nil
		})),
		validation.Field(&ts.Title, validation.Required.Error("title required")),
	)
}

type TeacherToSchool struct {
	ID        uuid.UUID `json:"id"`
	TeacherId uuid.UUID `json:"teacherId"`
	SchoolId  uuid.UUID `json:"schoolId"`
	Subject   string    `json:"subject"`
	Status    db.Status `json:"status"`
}

func (t TeacherToSchool) VlidateAssignTeacher() error {
	return validation.ValidateStruct(&t, validation.Field(&t.TeacherId, validation.By(func(value interface{}) error {
		if t.TeacherId == uuid.Nil {
			return fmt.Errorf("teacher id required")
		}
		return nil
	})),
		validation.Field(&t.SchoolId, validation.By(func(value interface{}) error {
			if t.SchoolId == uuid.Nil {
				return fmt.Errorf("school id required")
			}
			return nil
		})), validation.Field(&t.Subject, validation.Required))
}
