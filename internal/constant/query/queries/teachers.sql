-- name: CreateTechers :one 
INSERT INTO teachers(user_id,title,status) 
VALUES ($1,$2,$3)
RETURNING *;