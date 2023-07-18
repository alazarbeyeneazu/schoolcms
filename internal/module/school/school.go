package school

import (
	"context"
	"fmt"
	"schoolcms/internal/constant/dto"
	"schoolcms/internal/constant/errors"
	"schoolcms/internal/module"
	"schoolcms/internal/storage"
	"schoolcms/platform/logger"
	"schoolcms/platform/utils"

	"github.com/dongri/phonenumber"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type School struct {
	schoolPersistance storage.School
	log               logger.Logger
}

func Init(schoolPersistance storage.School, log logger.Logger) module.School {
	return &School{
		schoolPersistance: schoolPersistance,
		log:               log,
	}
}

func (s *School) CreateSchool(ctx context.Context, sc dto.School) (dto.School, error) {
	if err := sc.ValidateSchool(); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while validating the school")
		s.log.Error(ctx, "school validation failed ", zap.Error(err), zap.Any("school", sc))
		return dto.School{}, err
	}
	if len(sc.Phone) == 10 {
		sc.Phone = string("+251" + sc.Phone[1:])
	}
	return s.schoolPersistance.CreateSchool(ctx, sc)
}

func (s *School) AssignStudentToSchool(ctx context.Context, sc dto.StudentToSchool) (dto.StudentToSchool, error) {
	if err := sc.ValidateStudentToSchool(); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while validating student ")
		s.log.Error(ctx, "error while validating user ", zap.Error(err), zap.Any("student", sc))
		return dto.StudentToSchool{}, err
	}

	return s.schoolPersistance.AssignStudentToSchool(ctx, sc)
}
func (s *School) GetAllSchools(ctx context.Context, filter dto.GetSchoolsFilter) ([]dto.School, error) {
	if err := filter.Validate(); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while validation user input")
		s.log.Error(ctx, "error while validating user ", zap.Error(err), zap.Any("filter", filter))
		return []dto.School{}, err
	}

	if filter.Page != 0 {
		filter.Page = (filter.Page - 1) * filter.PerPage
	}
	return s.schoolPersistance.GetAllSchools(ctx, filter)
}

func (s *School) GetSchoolByID(ctx context.Context, id uuid.UUID) (dto.School, error) {
	if err := validation.Validate(id, validation.By(utils.CheckForNullUUID("school id required"))); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while validating school id")
		s.log.Error(ctx, "error while validating school id ", zap.Error(err), zap.Any("school id ", id))
		return dto.School{}, err
	}

	return s.schoolPersistance.GetSchoolByID(ctx, id)
}

func (s *School) GetSchoolByPhone(ctx context.Context, phone string) (dto.School, error) {
	if err := validation.Validate(phone, validation.Required.Error("phone required"),
		validation.By(func(value interface{}) error {
			phone := phonenumber.Parse(phone, "ET")
			if phone == "" {
				return fmt.Errorf("invalid phone number")
			}
			return nil
		})); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while validating school phone")
		s.log.Error(ctx, "error while validating school phone ", zap.Error(err), zap.Any("school phone ", phone))
		return dto.School{}, err
	}
	if len(phone) == 10 {
		phone = string("+251" + phone[1:])
	}
	return s.schoolPersistance.GetSchoolByPhone(ctx, phone)
}

func (s *School) UpdateSchoolStatus(ctx context.Context, stat dto.SchoolStatus) error {

	if err := stat.Validate(); err != nil {

		err = errors.ErrValidationError.Wrap(err, "error while validating school")

		s.log.Error(ctx, "error while validating school", zap.Error(err), zap.Any("school", stat))
		return err
	}
	return s.schoolPersistance.UpdateSchoolStatus(ctx, stat)
}

func (s *School) DeleteSchool(ctx context.Context, sc uuid.UUID) error {

	return s.schoolPersistance.DeleteSchool(ctx, sc)
}
