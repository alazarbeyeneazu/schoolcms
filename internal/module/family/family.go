package family

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

type family struct {
	log              logger.Logger
	familyPersistant storage.Family
}

func Init(familyPersistant storage.Family, log logger.Logger) module.Family {
	return &family{
		log:              log,
		familyPersistant: familyPersistant,
	}
}
func (f *family) CreateFamily(ctx context.Context, fam dto.Family) (dto.Family, error) {
	if err := fam.ValidateFamily(); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while validating family")
		f.log.Error(ctx, "error while validating family", zap.Error(err), zap.Any("family", fam))
		return dto.Family{}, err
	}
	fam.Status = db.StatusACTIVE
	return f.familyPersistant.CreateFamily(ctx, fam)

}
