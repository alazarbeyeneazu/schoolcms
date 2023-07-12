package school

import (
	"context"
	"database/sql"
	"schoolcms/internal/constant/dto"
	"schoolcms/internal/constant/errors"
	"schoolcms/internal/constant/model/db"
	persistencedb "schoolcms/internal/constant/persistenceDB"
	"schoolcms/internal/storage"
	"schoolcms/platform/logger"

	"go.uber.org/zap"
)

type school struct {
	db  persistencedb.PersistenceDB
	log logger.Logger
}

func Init(db persistencedb.PersistenceDB, log logger.Logger) storage.School {
	return &school{
		db:  db,
		log: log,
	}
}
func (s *school) CreateSchool(ctx context.Context, sc dto.School) (dto.School, error) {
	scl, err := s.db.CreateSchool(ctx, db.CreateSchoolParams{
		Name:  sc.Name,
		Phone: sc.Phone,
		Logo:  sql.NullString{String: sc.Logo, Valid: true},
	})

	if err != nil {
		err := errors.ErrWriteError.Wrap(err, "unable to register school")
		s.log.Error(ctx, "unable to create school ", zap.Error(err), zap.Any("school", sc))
		return dto.School{}, err
	}
	return dto.School{
		ID:    scl.ID,
		Name:  scl.Name,
		Phone: scl.Phone,
		Logo:  scl.Logo.String,
	}, nil
}

func (s *school) AssignStudentToSchool(ctx context.Context, std dto.StudentToSchool) (dto.StudentToSchool, error) {
	assignedStd, err := s.db.Queries.AssignStudent(ctx, db.AssignStudentParams{
		StudentID: std.StudentID,
		SchoolID:  std.SchoolID,
		GradeID:   std.GradeID,
		Status:    std.Status,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "error while assign student to school")
		s.log.Error(ctx, "error while writing to database", zap.Error(err), zap.Any("student", std))
		return dto.StudentToSchool{}, nil
	}
	return dto.StudentToSchool{
		ID:        assignedStd.ID,
		StudentID: assignedStd.StudentID,
		SchoolID:  assignedStd.SchoolID,
		GradeID:   assignedStd.GradeID,
		Status:    assignedStd.Status,
		CreatedAt: assignedStd.CreatedAt,
	}, nil
}
