package grade

import (
	"context"
	"schoolcms/internal/constant/dto"
	"schoolcms/internal/constant/errors"
	"schoolcms/internal/constant/model/db"
	"schoolcms/internal/module"
	"schoolcms/internal/storage"
	"schoolcms/platform/logger"

	"go.uber.org/zap"
)

type grade struct {
	gradePersistant storage.Grade
	log             logger.Logger
}

func Init(ctx context.Context, gradePersistant storage.Grade, log logger.Logger) module.Grade {
	return &grade{
		gradePersistant: gradePersistant,
		log:             log,
	}
}

func (g *grade) CreateGrade(ctx context.Context, grd dto.Grade) (dto.Grade, error) {
	if err := grd.ValidateGrade(); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while validating grade")
		g.log.Error(ctx, "error while validating grade", zap.Error(err), zap.Any("grade", grd))
		return dto.Grade{}, err
	}
	grd.Status = db.StatusACTIVE
	return g.gradePersistant.CreateGrade(ctx, grd)
}
