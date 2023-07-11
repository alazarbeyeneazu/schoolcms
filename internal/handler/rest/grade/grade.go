package grade

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

type grade struct {
	log            logger.Logger
	gradeModule    module.Grade
	contextTimeOut time.Duration
}

func Init(ctx context.Context, gradeModule module.Grade, timeout time.Duration, log logger.Logger) rest.Grade {
	return &grade{}
}
func (g *grade) CreateGrade(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, g.contextTimeOut)
	defer cancel()
	var grd dto.Grade
	if err := c.ShouldBind(&grd); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while binding user input to dto.Grade")
		g.log.Error(c, "error while validating user input to dto.Grade", zap.Error(err))
		return
	}
	grad, err := g.gradeModule.CreateGrade(ctx, grd)
	if err != nil {
		_ = c.Error(err)
		return
	}
	grad.ID = uuid.Nil
	response.SendSuccessResponse(c, http.StatusCreated, grad, nil)

}
