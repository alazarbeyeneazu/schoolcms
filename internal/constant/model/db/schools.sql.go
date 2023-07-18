// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: schools.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const assignStudent = `-- name: AssignStudent :one
INSERT INTO school_students(student_id,school_id,grade_id,status)
VALUES ($1,$2,$3,$4)
RETURNING id, student_id, school_id, grade_id, status, created_at, updated_at, deleted_at
`

type AssignStudentParams struct {
	StudentID uuid.UUID
	SchoolID  uuid.UUID
	GradeID   uuid.UUID
	Status    Status
}

func (q *Queries) AssignStudent(ctx context.Context, arg AssignStudentParams) (SchoolStudent, error) {
	row := q.db.QueryRow(ctx, assignStudent,
		arg.StudentID,
		arg.SchoolID,
		arg.GradeID,
		arg.Status,
	)
	var i SchoolStudent
	err := row.Scan(
		&i.ID,
		&i.StudentID,
		&i.SchoolID,
		&i.GradeID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const createSchool = `-- name: CreateSchool :one
INSERT INTO schools (name,logo,phone) VALUES ($1,$2,$3)
RETURNING id, name, logo, phone, status, created_at, updated_at, deleted_at
`

type CreateSchoolParams struct {
	Name  string
	Logo  sql.NullString
	Phone string
}

func (q *Queries) CreateSchool(ctx context.Context, arg CreateSchoolParams) (School, error) {
	row := q.db.QueryRow(ctx, createSchool, arg.Name, arg.Logo, arg.Phone)
	var i School
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Logo,
		&i.Phone,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getAllSchools = `-- name: GetAllSchools :many

select id, name, logo, phone, status, created_at, updated_at, deleted_at from schools where deleted_at is null order by created_at ASC  limit $1 offset $2
`

type GetAllSchoolsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetAllSchools(ctx context.Context, arg GetAllSchoolsParams) ([]School, error) {
	rows, err := q.db.Query(ctx, getAllSchools, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []School
	for rows.Next() {
		var i School
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Logo,
			&i.Phone,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSchoolById = `-- name: GetSchoolById :one
SELECT id, name, logo, phone, status, created_at, updated_at, deleted_at FROM schools where id = $1
`

func (q *Queries) GetSchoolById(ctx context.Context, id uuid.UUID) (School, error) {
	row := q.db.QueryRow(ctx, getSchoolById, id)
	var i School
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Logo,
		&i.Phone,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
