-- name: CreateSchool :one 
INSERT INTO schools (name,logo,phone) VALUES ($1,$2,$3)
RETURNING *;

-- name: AssignStudent :one 
INSERT INTO school_students(student_id,school_id,grade_id,status)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: GetAllSchools :many

select * from schools where deleted_at is null order by created_at ASC  limit $1 offset $2  ;  

-- name: GetSchoolById :one 
SELECT * FROM schools where id = $1 and deleted_at is null;

-- name: GetSchoolByPhone :one 
SELECT * FROM schools where phone = $1 and deleted_at is null;

-- name: UpdateSchoolStatus :exec 
update schools set status = $1 where id = $2 and deleted_at is null; 

-- name: DeleteSchool :exec 
update schools set deleted_at = now() where id = $1; 

-- name: UpdateSchoolInformations :one

UPDATE schools set name = $1 , logo = $2 , phone = $3 where phone = $4
RETURNING *;