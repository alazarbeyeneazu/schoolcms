package student

import (
	"context"
	"schoolcms/internal/constant/dto"
	"schoolcms/internal/constant/errors"
	"schoolcms/internal/constant/model/db"
	persistencedb "schoolcms/internal/constant/persistenceDB"
	"schoolcms/internal/storage"
	"schoolcms/platform/logger"

	"go.uber.org/zap"
)

type student struct {
	db  persistencedb.PersistenceDB
	log logger.Logger
}

func Init(db persistencedb.PersistenceDB, log logger.Logger) storage.Student {
	return &student{
		log: log,
		db:  db,
	}
}
func (s *student) CreateStudent(ctx context.Context, std dto.Student) (dto.Student, error) {
	student, err := s.db.Queries.CreateStudent(ctx, db.CreateStudentParams{
		UserID: std.UserID,
		Status: std.Status,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "error while creating student")
		s.log.Error(ctx, "error while creating student ", zap.Error(err), zap.Any("student", std))
		return dto.Student{}, err
	}
	return dto.Student{
		ID:        student.ID,
		UserID:    student.UserID,
		Status:    student.Status,
		CreatedAt: student.CreatedAt,
	}, nil
}
