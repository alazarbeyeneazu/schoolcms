package user

import (
	"context"
	"net/http"
	"schoolcms/internal/constant/dto"
	"schoolcms/internal/constant/errors"
	"schoolcms/internal/constant/model/response"
	"schoolcms/internal/handler/rest"
	"schoolcms/internal/module"
	"schoolcms/platform/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type user struct {
	UserModule     module.User
	log            logger.Logger
	ContextTImeOut time.Duration
}

func Init(usermodule module.User, log logger.Logger, contextTimeOut time.Duration) rest.User {
	return &user{
		UserModule:     usermodule,
		log:            log,
		ContextTImeOut: contextTimeOut,
	}
}
func (u *user) CreatUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, u.ContextTImeOut)
	defer cancel()
	var usr dto.User
	if err := c.ShouldBind(&usr); err != nil {
		err = errors.ErrValidationError.Wrap(err, "invalid input")
		u.log.Error(ctx, "invalid  user input", zap.Error(err))
		_ = c.Error(err)

		return
	}
	userDetail, err := u.UserModule.CreateUser(ctx, usr)
	if err != nil {
		_ = c.Error(err)
		return
	}
	userDetail.ID = uuid.Nil
	response.SendSuccessResponse(c, http.StatusOK, userDetail, nil)
}
