package user

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

type user struct {
	db  persistencedb.PersistenceDB
	log logger.Logger
}

func Init(db persistencedb.PersistenceDB, log logger.Logger) storage.User {
	return &user{
		db:  db,
		log: log,
	}
}
func (u *user) CreateUser(ctx context.Context, ur dto.User) (dto.User, error) {
	usr, err := u.db.Queries.CreateUser(ctx, db.CreateUserParams{
		FirstName:  sql.NullString{String: ur.FirstName, Valid: true},
		MiddleName: sql.NullString{String: ur.MiddleName, Valid: true},
		LastName:   sql.NullString{String: ur.LastName, Valid: true},
		Phone:      sql.NullString{String: ur.Phone, Valid: true},
		Profile:    sql.NullString{String: ur.Profile, Valid: true},
		Status:     ur.Status,
	})
	if err != nil {
		err := errors.ErrWriteError.Wrap(err, "unable to register user")
		u.log.Error(ctx, "unable to create user ", zap.Error(err), zap.Any("user", ur))
		return dto.User{}, err
	}
	return dto.User{
		ID:         usr.ID,
		FirstName:  usr.FirstName.String,
		MiddleName: usr.MiddleName.String,
		LastName:   usr.LastName.String,
		Phone:      usr.Phone.String,
		Profile:    usr.Profile.String,
		Status:     usr.Status,
	}, nil
}
