-- name: CreateTechers :one 
INSERT INTO teachers(user_id,title,status) 
VALUES ($1,$2,$3)
RETURNING *;

-- name: AssignTeachersToSchool :one 

INSERT INTO school_teachers  (school_id,teacher_id,subject,status)
VALUES ($1,$2,$3,$4)
RETURNING *;

