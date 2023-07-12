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
	teacherpersistent storage.Teacher
}

func Init(log logger.Logger, teacherpersistent storage.Teacher) module.Teacher {
	return &teacher{
		log:               log,
		teacherpersistent: teacherpersistent,
	}
}

func (t *teacher) CreateTeacher(ctx context.Context, tc dto.Teacher) (dto.Teacher, error) {
	if err := tc.ValidteTeachers(); err != nil {
		err = errors.ErrValidationError.Wrap(err, "validation error")
		t.log.Error(ctx, "validation error ", zap.Error(err), zap.Any("teacher ", tc))
		return dto.Teacher{}, err
	}

	return t.teacherpersistent.CreateTeacher(ctx, tc)
}

func (t *teacher) AssignTeachersToSchool(ctx context.Context, tc dto.TeacherToSchool) (dto.TeacherToSchool, error) {
	if err := tc.VlidateAssignTeacher(); err != nil {
		err = errors.ErrValidationError.Wrap(err, "validation error")
		t.log.Error(ctx, "error while validating teacher to dto.TeacherToSchool ", zap.Error(err), zap.Any("teacher ", tc))
		return dto.TeacherToSchool{}, err
	}

	return t.teacherpersistent.AssignTeacherToSchool(ctx, tc)
}
