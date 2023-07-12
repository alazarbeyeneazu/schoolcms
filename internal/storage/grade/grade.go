package grade

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

type grade struct {
	db  persistencedb.PersistenceDB
	log logger.Logger
}

func Init(db persistencedb.PersistenceDB, log logger.Logger) storage.Grade {
	return &grade{
		log: log,
		db:  db,
	}
}

func (g *grade) CreateGrade(ctx context.Context, grd dto.Grade) (dto.Grade, error) {
	gr, err := g.db.Queries.CreateGrade(ctx, db.CreateGradeParams{
		Name:     grd.Name,
		SchoolID: grd.SchoolID,
		Status:   grd.Status,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "error while creating grade")
		g.log.Error(ctx, "error while creating grade", zap.Error(err), zap.Any("grade", grd))
		return dto.Grade{}, err
	}
	return dto.Grade{
		ID:        gr.ID,
		SchoolID:  gr.SchoolID,
		Status:    gr.Status,
		Name:      gr.Name,
		CreatedAt: gr.CreatedAt,
	}, nil
}
