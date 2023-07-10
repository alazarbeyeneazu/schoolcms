package teacher

import (
	"context"
	"schoolcms/internal/constant/dto"
	"schoolcms/internal/constant/errors"
	"schoolcms/internal/module"
	"schoolcms/internal/storage"
	"schoolcms/platform/logger"

	"go.uber.org/zap"
)

type teacher struct {
	log               logger.Logger
	teacherPersistant storage.Teacher
}

func Init(log logger.Logger, teacherPersistant storage.Teacher) module.Teacher {
	return &teacher{
		log:               log,
		teacherPersistant: teacherPersistant,
	}
}

func (t *teacher) CreateTeacher(ctx context.Context, tc dto.Teacher) (dto.Teacher, error) {
	if err := tc.ValidteTeachers(); err != nil {
		err = errors.ErrValidationError.Wrap(err, "validation error")
		t.log.Error(ctx, "validation error ", zap.Error(err), zap.Any("teacher ", tc))
		return dto.Teacher{}, err
	}

	return t.teacherPersistant.CreateTeacher(ctx, tc)
}
