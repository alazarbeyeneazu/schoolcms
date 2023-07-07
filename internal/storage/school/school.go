package school

import (
	"context"
	"database/sql"
	"schoolcms/internal/constant/dto"
	"schoolcms/internal/constant/errors"
	"schoolcms/internal/constant/model/db"
	persistencedb "schoolcms/internal/constant/persistenceDB"
	"schoolcms/internal/storage"
	"schoolcms/platform/logger"

	"go.uber.org/zap"
)

type school struct {
	db  persistencedb.PersistenceDB
	log logger.Logger
}

func Init(db persistencedb.PersistenceDB, log logger.Logger) storage.School {
	return &school{
		db:  db,
		log: log,
	}
}
func (s *school) CreateSchool(ctx context.Context, sc dto.School) (dto.School, error) {
	scl, err := s.db.CreateSchool(ctx, db.CreateSchoolParams{
		Name: sc.Name,
		Logo: sql.NullString{String: sc.Log, Valid: true},
	})

	if err != nil {
		err := errors.ErrWriteError.Wrap(err, "unable to register school")
		s.log.Error(ctx, "unable to create school ", zap.Error(err), zap.Any("school", sc))
		return dto.School{}, err
	}
	return dto.School{
		ID:   scl.ID,
		Name: scl.Name,
		Log:  scl.Logo.String,
	}, nil
}
