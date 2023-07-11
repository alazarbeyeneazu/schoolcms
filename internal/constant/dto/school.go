package dto

import (
	"fmt"
	"schoolcms/internal/constant/model/db"
	"schoolcms/platform/utils"
	"time"

	"github.com/dongri/phonenumber"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type School struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Status db.Status `json:"status"`
	Phone  string    `json:"phone"`
	Log    string    `json:"logo"`
}

func (s School) ValidateSchool() error {

	return validation.ValidateStruct(&s,
		validation.Field(&s.Name, validation.Required.Error("school name required")),
		validation.Field(&s.Phone, validation.By(func(value interface{}) error {
			phone := phonenumber.Parse(s.Phone, "ET")
			if phone == "" {
				return fmt.Errorf("invalid phone number")
			}
			return nil
		})),
	)
}

type StudentToSchool struct {
	ID        uuid.UUID `json:"id"`
	StudentId uuid.UUID `json:"student_id"`
	SchoolId  uuid.UUID `json:"school_id"`
	GradeId   uuid.UUID `json:"grade_id"`
	Status    db.Status `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (s StudentToSchool) ValidateStudentToSchool() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.StudentId, validation.By(utils.CheckForNullUUID("student id required"))),
		validation.Field(&s.SchoolId, validation.By(utils.CheckForNullUUID("school id required"))),
		validation.Field(&s.GradeId, validation.By(utils.CheckForNullUUID("grade id required"))),
	)
}
