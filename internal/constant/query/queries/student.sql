-- name: CreateStudent :one 
INSERT INTO students (user_id,status)
VALUES ($1,$2) RETURNING *;

