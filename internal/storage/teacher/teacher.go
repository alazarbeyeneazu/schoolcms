package teacher

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

type teacher struct {
	db  persistencedb.PersistenceDB
	log logger.Logger
}

func Init(db persistencedb.PersistenceDB, log logger.Logger) storage.Teacher {
	return &teacher{
		log: log,
		db:  db,
	}
}

func (t *teacher) CreateTeacher(ctx context.Context, tc dto.Teacher) (dto.Teacher, error) {
	teacher, err := t.db.Queries.CreateTechers(ctx, db.CreateTechersParams{
		UserID: tc.UserId,
		Title:  tc.Title,
		Status: db.StatusACTIVE,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "unable to register teacher")
		t.log.Error(ctx, "unable to register teacher", zap.Error(err), zap.Any("teacher ", tc))
		return dto.Teacher{}, err
	}

	return dto.Teacher{
		ID:     teacher.ID,
		UserId: teacher.UserID,
		Title:  teacher.Title,
		Status: teacher.Status,
	}, nil
}
