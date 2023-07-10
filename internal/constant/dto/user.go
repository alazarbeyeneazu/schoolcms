package dto

import (
	"fmt"
	"schoolcms/internal/constant/model/db"

	"github.com/dongri/phonenumber"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"firstName"`
	MiddleName string    `json:"middleName"`
	LastName   string    `json:"lastName"`
	Phone      string    `json:"phone"`
	Gender     string    `json:"gender"`
	Profile    string    `json:"profile"`
	Status     db.Status `json:"status"`
}

func (u User) ValidateUser() error {
	return validation.ValidateStruct(&u,

		validation.Field(&u.FirstName, validation.Required.Error("first name  is required")),
		validation.Field(&u.MiddleName, validation.Required.Error("middle name  is required")),
		validation.Field(&u.FirstName, validation.Length(1, 64).Error("first name can not be more than 64 characters")),
		validation.Field(&u.MiddleName, validation.Length(1, 64).Error("first name can not be more than 64 characters")),
		validation.Field(&u.LastName, validation.Length(1, 64).Error("first name can not be more than 64 characters")),
		validation.Field(&u.Phone, validation.By(func(value interface{}) error {
			phone := phonenumber.Parse(u.Phone, "ET")
			if phone == "" {
				return fmt.Errorf("invalid phone number")
			}
			return nil
		})),
		validation.Field(&u.Gender, validation.Required.Error("gender required")),
	)

}
