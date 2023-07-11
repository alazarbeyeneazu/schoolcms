-- name: CreateStudent :one 
INSERT INTO students (user_id,status)
VALUES ($1,$2) RETURNING *;

-- name: AssignStudent :one 
INSERT INTO school_students(user_id,school_id,grade_id,status)
VALUES ($1,$2,$3,$4)
RETURNING *;