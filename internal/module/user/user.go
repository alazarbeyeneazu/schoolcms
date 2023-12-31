package user

import (
	"context"

	"schoolcms/internal/constant/dto"
	"schoolcms/internal/constant/errors"
	"schoolcms/internal/constant/model/db"
	"schoolcms/internal/module"
	"schoolcms/internal/storage"
	"schoolcms/platform/logger"

	"go.uber.org/zap"
)

type user struct {
	log            logger.Logger
	userpersistent storage.User
}

func Init(log logger.Logger, usrpersistent storage.User) module.User {
	return &user{
		log:            log,
		userpersistent: usrpersistent,
	}
}
func (u *user) CreateUser(ctx context.Context, ur dto.User) (dto.User, error) {
	if err := ur.ValidateUser(); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while validating the user")
		u.log.Error(ctx, "user validation failed ", zap.Error(err), zap.Any("user", ur))
		return dto.User{}, err
	}
	ur.Status = db.StatusACTIVE
	return u.userpersistent.CreateUser(ctx, ur)
}
