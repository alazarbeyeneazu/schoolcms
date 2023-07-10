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