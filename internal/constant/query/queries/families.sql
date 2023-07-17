-- name: CreateFamilies :one 
INSERT INTO families(user_id,family_type,status)
VALUES ($1,$2,$3)
RETURNING *;
