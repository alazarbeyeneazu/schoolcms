// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusACTIVE      Status = "ACTIVE"
	StatusPENDING     Status = "PENDING"
	StatusDIACTIVATED Status = "DIACTIVATED"
)

func (e *Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Status(s)
	case string:
		*e = Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Status: %T", src)
	}
	return nil
}

type NullStatus struct {
	Status Status
	Valid  bool // Valid is true if Status is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStatus) Scan(value interface{}) error {
	if value == nil {
		ns.Status, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Status.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Status), nil
}

type Grade struct {
	ID        uuid.UUID
	Name      string
	SchoolID  uuid.UUID
	Status    Status
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type School struct {
	ID        uuid.UUID
	Name      string
	Logo      sql.NullString
	Phone     string
	Status    NullStatus
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
	DeletedAt sql.NullString
}

type SchoolTeacher struct {
	ID        uuid.UUID
	SchoolID  uuid.UUID
	TeacherID uuid.UUID
	Subject   string
	Status    Status
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type Teacher struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Title     string
	Status    Status
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type User struct {
	ID         uuid.UUID
	FirstName  sql.NullString
	MiddleName sql.NullString
	LastName   sql.NullString
	Gender     string
	Phone      sql.NullString
	Profile    sql.NullString
	Status     Status
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
	DeletedAt  sql.NullTime
}
