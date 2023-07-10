-- name: CreateGrade :one 
INSERT INTO grades (name,school_id,status) 
VALUES($1,$2,$3) 
RETURNING *;