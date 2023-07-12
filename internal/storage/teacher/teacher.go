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
		UserID: tc.UserID,
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
		UserID: teacher.UserID,
		Title:  teacher.Title,
		Status: teacher.Status,
	}, nil
}

func (t *teacher) AssignTeacherToSchool(ctx context.Context, tToS dto.TeacherToSchool) (dto.TeacherToSchool, error) {
	assignedTeacher, err := t.db.Queries.AssignTeachersToSchool(ctx, db.AssignTeachersToSchoolParams{
		SchoolID:  tToS.SchoolID,
		TeacherID: tToS.TeacherID,
		Subject:   tToS.Subject,
		Status:    tToS.Status,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "error while assign teacher ")

		t.log.Error(ctx, "error while assign teacher", zap.Error(err), zap.Any("teacher assign", tToS))
		return dto.TeacherToSchool{}, err

	}
	return dto.TeacherToSchool{
		ID:        assignedTeacher.ID,
		TeacherID: assignedTeacher.TeacherID,
		SchoolID:  assignedTeacher.TeacherID,
		Subject:   assignedTeacher.Subject,
		Status:    assignedTeacher.Status,
	}, nil
}
