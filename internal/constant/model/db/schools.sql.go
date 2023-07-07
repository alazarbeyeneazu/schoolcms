// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: schools.sql

package db

import (
	"context"
	"database/sql"
)

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
