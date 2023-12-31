// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: teachers.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const assignTeachersToSchool = `-- name: AssignTeachersToSchool :one

INSERT INTO school_teachers  (school_id,teacher_id,subject,status)
VALUES ($1,$2,$3,$4)
RETURNING id, school_id, teacher_id, subject, status, created_at, updated_at, deleted_at
`

type AssignTeachersToSchoolParams struct {
	SchoolID  uuid.UUID
	TeacherID uuid.UUID
	Subject   string
	Status    Status
}

func (q *Queries) AssignTeachersToSchool(ctx context.Context, arg AssignTeachersToSchoolParams) (SchoolTeacher, error) {
	row := q.db.QueryRow(ctx, assignTeachersToSchool,
		arg.SchoolID,
		arg.TeacherID,
		arg.Subject,
		arg.Status,
	)
	var i SchoolTeacher
	err := row.Scan(
		&i.ID,
		&i.SchoolID,
		&i.TeacherID,
		&i.Subject,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const createTechers = `-- name: CreateTechers :one
INSERT INTO teachers(user_id,title,status) 
VALUES ($1,$2,$3)
RETURNING id, user_id, title, status, created_at, updated_at, deleted_at
`

type CreateTechersParams struct {
	UserID uuid.UUID
	Title  string
	Status Status
}

func (q *Queries) CreateTechers(ctx context.Context, arg CreateTechersParams) (Teacher, error) {
	row := q.db.QueryRow(ctx, createTechers, arg.UserID, arg.Title, arg.Status)
	var i Teacher
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
