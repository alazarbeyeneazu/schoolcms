package student

import (
	"context"
	"schoolcms/internal/constant/dto"
	"schoolcms/internal/constant/errors"
	"schoolcms/internal/constant/model/db"
	"schoolcms/internal/storage"
	"schoolcms/platform/logger"
)

type student struct {
	log               logger.Logger
	studentpersistent storage.Student
}

func Init(std storage.Student, log logger.Logger) storage.Student {
	return &student{
		log:               log,
		studentpersistent: std,
	}
}

func (s *student) CreateStudent(ctx context.Context, std dto.Student) (dto.Student, error) {
	if err := std.ValidateStudent(); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while validating student")
		s.log.Error(ctx, "error while validating user")
		return dto.Student{}, err
	}
	std.Status = db.StatusACTIVE

	return s.studentpersistent.CreateStudent(ctx, std)
}
