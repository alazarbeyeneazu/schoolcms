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

	"github.com/google/uuid"
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
		Name:  sc.Name,
		Phone: sc.Phone,
		Logo:  sql.NullString{String: sc.Logo, Valid: true},
	})

	if err != nil {
		err := errors.ErrWriteError.Wrap(err, "unable to register school")
		s.log.Error(ctx, "unable to create school ", zap.Error(err), zap.Any("school", sc))
		return dto.School{}, err
	}
	return dto.School{
		ID:    scl.ID,
		Name:  scl.Name,
		Phone: scl.Phone,
		Logo:  scl.Logo.String,
	}, nil
}

func (s *school) AssignStudentToSchool(ctx context.Context, std dto.StudentToSchool) (dto.StudentToSchool, error) {
	assignedStd, err := s.db.Queries.AssignStudent(ctx, db.AssignStudentParams{
		StudentID: std.StudentID,
		SchoolID:  std.SchoolID,
		GradeID:   std.GradeID,
		Status:    std.Status,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "error while assign student to school")
		s.log.Error(ctx, "error while writing to database", zap.Error(err), zap.Any("student", std))
		return dto.StudentToSchool{}, nil
	}
	return dto.StudentToSchool{
		ID:        assignedStd.ID,
		StudentID: assignedStd.StudentID,
		SchoolID:  assignedStd.SchoolID,
		GradeID:   assignedStd.GradeID,
		Status:    assignedStd.Status,
		CreatedAt: assignedStd.CreatedAt,
	}, nil
}

func (s *school) GetAllSchools(ctx context.Context, filter dto.GetSchoolsFilter) ([]dto.School, error) {
	schools := []dto.School{}
	retschools, err := s.db.Queries.GetAllSchools(ctx, db.GetAllSchoolsParams{
		Limit:  filter.PerPage,
		Offset: filter.Page,
	})
	if err != nil {
		err = errors.ErrReadError.Wrap(err, "error while reading schools")
		s.log.Error(ctx, "error while reading schools", zap.Error(err), zap.Any("filter", filter))
		return []dto.School{}, err
	}
	for _, sc := range retschools {
		schools = append(schools, dto.School{
			ID:        sc.ID,
			Name:      sc.Name,
			Status:    sc.Status.Status,
			Phone:     sc.Phone,
			Logo:      sc.Logo.String,
			UpdatedAt: sc.CreatedAt.Time,
			CreatedAt: sc.CreatedAt.Time,
		})
	}
	return schools, err
}
func (s *school) GetSchoolByID(ctx context.Context, id uuid.UUID) (dto.School, error) {

	retSchool, err := s.db.Queries.GetSchoolById(ctx, id)
	if err != nil {
		err = errors.ErrReadError.Wrap(err, "error while reading schools")
		s.log.Error(ctx, "error while reading school", zap.Error(err), zap.Any("school id ", id))
		return dto.School{}, err
	}
	return dto.School{
		ID:        retSchool.ID,
		Name:      retSchool.Name,
		Status:    retSchool.Status.Status,
		Phone:     retSchool.Phone,
		Logo:      retSchool.Logo.String,
		CreatedAt: retSchool.CreatedAt.Time,
		UpdatedAt: retSchool.UpdatedAt.Time,
		DeletedAt: retSchool.DeletedAt.Time,
	}, err
}

func (s *school) GetSchoolByPhone(ctx context.Context, phone string) (dto.School, error) {

	retSchool, err := s.db.Queries.GetSchoolByPhone(ctx, phone)
	if err != nil {
		err = errors.ErrReadError.Wrap(err, "error while reading schools")
		s.log.Error(ctx, "error while reading school", zap.Error(err), zap.Any("school phone ", phone))
		return dto.School{}, err
	}
	return dto.School{
		ID:        retSchool.ID,
		Name:      retSchool.Name,
		Status:    retSchool.Status.Status,
		Phone:     retSchool.Phone,
		Logo:      retSchool.Logo.String,
		CreatedAt: retSchool.CreatedAt.Time,
		UpdatedAt: retSchool.UpdatedAt.Time,
		DeletedAt: retSchool.DeletedAt.Time,
	}, err
}

func (s *school) UpdateSchoolStatus(ctx context.Context, stat dto.SchoolStatus) error {
	return s.db.Queries.UpdateSchoolStatus(ctx, db.UpdateSchoolStatusParams{
		Status: db.NullStatus{Status: stat.Status, Valid: true},
		ID:     stat.SchoolID})
}
