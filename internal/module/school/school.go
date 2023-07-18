package school

import (
	"context"
	"schoolcms/internal/constant/dto"
	"schoolcms/internal/constant/errors"
	"schoolcms/internal/module"
	"schoolcms/internal/storage"
	"schoolcms/platform/logger"

	"go.uber.org/zap"
)

type School struct {
	schoolPersistance storage.School
	log               logger.Logger
}

func Init(schoolPersistance storage.School, log logger.Logger) module.School {
	return &School{
		schoolPersistance: schoolPersistance,
		log:               log,
	}
}

func (s *School) CreateSchool(ctx context.Context, sc dto.School) (dto.School, error) {
	if err := sc.ValidateSchool(); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while validating the school")
		s.log.Error(ctx, "school validation failed ", zap.Error(err), zap.Any("school", sc))
		return dto.School{}, err
	}

	return s.schoolPersistance.CreateSchool(ctx, sc)
}

func (s *School) AssignStudentToSchool(ctx context.Context, sc dto.StudentToSchool) (dto.StudentToSchool, error) {
	if err := sc.ValidateStudentToSchool(); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while validating student ")
		s.log.Error(ctx, "error while validating user ", zap.Error(err), zap.Any("student", sc))
		return dto.StudentToSchool{}, err
	}

	return s.schoolPersistance.AssignStudentToSchool(ctx, sc)
}
func (s *School) GetAllSchools(ctx context.Context, filter dto.GetSchoolsFilter) ([]dto.School, error) {
	if err := filter.Validate(); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while validation user input")
		s.log.Error(ctx, "error while validating user ", zap.Error(err), zap.Any("filter", filter))
		return []dto.School{}, err
	}

	if filter.Page != 0 {
		filter.Page = (filter.Page - 1) * filter.PerPage
	}
	return s.schoolPersistance.GetAllSchools(ctx, filter)
}
