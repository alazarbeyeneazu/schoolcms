-- name: CreateSchool :one 
INSERT INTO schools (name,logo) VALUES ($1,$2)
RETURNING *;

