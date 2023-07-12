// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: families.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createFamilies = `-- name: CreateFamilies :one
INSERT INTO families(user_id,family_type,status)
VALUES ($1,$2,$3)
RETURNING id, user_id, family_type, status, created_at, updated_at, deleted_at
`

type CreateFamiliesParams struct {
	UserID     uuid.UUID
	FamilyType string
	Status     Status
}

func (q *Queries) CreateFamilies(ctx context.Context, arg CreateFamiliesParams) (Family, error) {
	row := q.db.QueryRow(ctx, createFamilies, arg.UserID, arg.FamilyType, arg.Status)
	var i Family
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FamilyType,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
