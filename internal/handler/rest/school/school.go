package school

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
	"go.uber.org/zap"
)

type school struct {
	schoolModule   module.School
	log            logger.Logger
	contextTimeout time.Duration
}

func Init(ctx context.Context, schoolModule module.School, log logger.Logger, contextTimeout time.Duration) rest.School {
	return &school{
		schoolModule:   schoolModule,
		log:            log,
		contextTimeout: contextTimeout,
	}
}

func (s *school) CreateSchool(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	var sc dto.School
	if err := c.ShouldBind(&sc); err != nil {
		err = errors.ErrValidationError.Wrap(err, "invalid input")
		s.log.Error(ctx, "invalid  user input", zap.Error(err))
		_ = c.Error(err)

		return
	}
	schoolDetail, err := s.schoolModule.CreateSchool(ctx, sc)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response.SendSuccessResponse(c, http.StatusCreated, schoolDetail, nil)
}
