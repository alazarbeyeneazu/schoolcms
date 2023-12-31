package storage

import (
	"context"
	"schoolcms/internal/constant/dto"

	"github.com/google/uuid"
)

type User interface {
	CreateUser(ctx context.Context, ur dto.User) (dto.User, error)
}

type School interface {
	CreateSchool(ctx context.Context, sc dto.School) (dto.School, error)
	AssignStudentToSchool(ctx context.Context, std dto.StudentToSchool) (dto.StudentToSchool, error)
	GetAllSchools(ctx context.Context, filter dto.GetSchoolsFilter) ([]dto.School, error)
	GetSchoolByID(ctx context.Context, id uuid.UUID) (dto.School, error)
	GetSchoolByPhone(ctx context.Context, phone string) (dto.School, error)
	UpdateSchoolStatus(ctx context.Context, stat dto.SchoolStatus) error
	DeleteSchool(ctx context.Context, stat uuid.UUID) error
	UpdateSchoolInformation(ctx context.Context, sc dto.School, oldPhone string) (dto.School, error)
}

type Teacher interface {
	CreateTeacher(ctx context.Context, tc dto.Teacher) (dto.Teacher, error)
	AssignTeacherToSchool(ctx context.Context, tToS dto.TeacherToSchool) (dto.TeacherToSchool, error)
}

type Grade interface {
	CreateGrade(ctx context.Context, grd dto.Grade) (dto.Grade, error)
}

type Student interface {
	CreateStudent(ctx context.Context, std dto.Student) (dto.Student, error)
}

type Family interface {
	CreateFamily(ctx context.Context, fam dto.Family) (dto.Family, error)
	AssignFamilyToStudent(ctx context.Context, fam dto.FamilyToStudent) (dto.FamilyToStudent, error)
}
