-- name: CreateUser :one 
INSERT INTO users (first_name,middle_name,last_name,phone,profile,status) 
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;


