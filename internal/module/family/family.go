package family

import (
	"context"
	"schoolcms/internal/constant/dto"
	"schoolcms/internal/constant/errors"
	"schoolcms/internal/module"
	"schoolcms/platform/logger"

	"go.uber.org/zap"
)

type family struct {
	log          logger.Logger
	familyModule module.Family
}

func Init(familyModule module.Family, log logger.Logger) module.Family {
	return &family{
		log:          log,
		familyModule: familyModule,
	}
}
func (f *family) CreateFamily(ctx context.Context, fam dto.Family) (dto.Family, error) {
	if err := fam.ValidateFamily(); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while validating family")
		f.log.Error(ctx, "error while validating family", zap.Error(err), zap.Any("family", fam))
		return dto.Family{}, err
	}
	return f.familyModule.CreateFamily(ctx, fam)

}
