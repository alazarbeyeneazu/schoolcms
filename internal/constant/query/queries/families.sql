-- name: CreateFamilies :one 
INSERT INTO families(user_id,family_type,status)
VALUES ($1,$2,$3)
RETURNING *;

-- name: AssignFamilyToStudent :one
INSERT INTO family_to_students(student_id,family_id,family_type,status)
VALUES ($1,$2,$3,$4)
RETURNING *;