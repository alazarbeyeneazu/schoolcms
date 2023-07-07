package school

import (
	"context"
	"database/sql"
	"schoolcms/internal/constant/dto"
	"schoolcms/internal/constant/model/db"
	persistencedb "schoolcms/internal/constant/persistenceDB"
	"schoolcms/internal/storage"
	"schoolcms/platform/logger"
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
func (s *school) CreateSchool(ctx context.Context, ur dto.School) (dto.School, error) {
	sc, err := s.db.CreateSchool(ctx, db.CreateSchoolParams{
		Name: ur.Name,
		Logo: sql.NullString{String: ur.Log, Valid: true},
	})

	if err != nil {

	}
	return dto.School{
		ID:   sc.ID,
		Name: sc.Name,
		Log:  sc.Logo.String,
	}, nil
}
