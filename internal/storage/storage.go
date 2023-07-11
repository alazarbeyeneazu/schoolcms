package storage

import (
	"context"
	"schoolcms/internal/constant/dto"
)

type User interface {
	CreateUser(ctx context.Context, ur dto.User) (dto.User, error)
}

type School interface {
	CreateSchool(ctx context.Context, sc dto.School) (dto.School, error)
}

type Teacher interface {
	CreateTeacher(ctx context.Context, tc dto.Teacher) (dto.Teacher, error)
	AssignTeacherToSchool(ctx context.Context, tToS dto.TeacherToSchool) (dto.TeacherToSchool, error)
}

type Grade interface {
	CreateGrade(ctx context.Context, grd dto.Grade) (dto.Grade, error)
}
