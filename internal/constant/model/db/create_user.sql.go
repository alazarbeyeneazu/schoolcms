// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: create_user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (first_name,middle_name,last_name,phone,profile,status) 
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING id, first_name, middle_name, last_name, gender, phone, profile, status, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	FirstName  sql.NullString
	MiddleName sql.NullString
	LastName   sql.NullString
	Phone      sql.NullString
	Profile    sql.NullString
	Status     Status
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.FirstName,
		arg.MiddleName,
		arg.LastName,
		arg.Phone,
		arg.Profile,
		arg.Status,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.Gender,
		&i.Phone,
		&i.Profile,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
