-- name: CreateSchool :one 
INSERT INTO schools (name,logo,phone) VALUES ($1,$2,$3)
RETURNING *;

-- name: AssignStudent :one 
INSERT INTO school_students(student_id,school_id,grade_id,status)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: GetAllSchools :many

select * from schools where deleted_at is null order by $1 ASC  limit $2 offset $3  ;  