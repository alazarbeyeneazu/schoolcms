package family

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

type family struct {
	db  persistencedb.PersistenceDB
	log logger.Logger
}

func Init(db persistencedb.PersistenceDB, log logger.Logger) storage.Family {
	return &family{
		db:  db,
		log: log,
	}
}

func (f *family) CreateFamily(ctx context.Context, fam dto.Family) (dto.Family, error) {
	familyRet, err := f.db.Queries.CreateFamilies(ctx, db.CreateFamiliesParams{
		UserID:     fam.UserID,
		FamilyType: fam.FamilyType,
		Status:     fam.Status,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "error while creating family")
		f.log.Error(ctx, "error while creating family", zap.Error(err), zap.Any("family", fam))
		return dto.Family{}, err
	}
	return dto.Family{
		ID:         familyRet.ID,
		UserID:     familyRet.UserID,
		Status:     familyRet.Status,
		FamilyType: familyRet.FamilyType,
		CreatedAt:  fam.CreatedAt,
		UpdatedAt:  fam.UpdatedAt,
		DeletedAt:  fam.DeletedAt,
	}, err
}
