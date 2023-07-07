-- name: CreateSchool :one 
INSERT INTO schools (name,logo,phone) VALUES ($1,$2,$3)
RETURNING *;

