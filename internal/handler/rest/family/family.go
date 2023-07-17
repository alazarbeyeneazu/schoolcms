package family

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

type family struct {
	log            logger.Logger
	familyModule   module.Family
	contextTImeOut time.Duration
}

func Init(log logger.Logger, familyModule module.Family, contextTimeOut time.Duration) rest.Family {
	return &family{
		log:            log,
		familyModule:   familyModule,
		contextTImeOut: contextTimeOut,
	}
}
func (f *family) CreateFamily(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, f.contextTImeOut)
	defer cancel()

	var fam dto.Family
	if err := c.ShouldBind(&fam); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while binding user input to Family")
		f.log.Error(ctx, "error while binding user input ", zap.Error(err))
		_ = c.Error(err)
		return
	}
	retFam, err := f.familyModule.CreateFamily(ctx, fam)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusCreated, retFam, nil)
}

func (f *family) AssignFamilyToStudent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, f.contextTImeOut)
	defer cancel()

	var fam dto.FamilyToStudent
	if err := c.ShouldBind(&fam); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while binding user input to Family Student")
		f.log.Error(ctx, "error while binding user input ", zap.Error(err))
		_ = c.Error(err)
		return
	}
	retFam, err := f.familyModule.AssignFamilyToStudent(ctx, fam)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessResponse(c, http.StatusCreated, retFam, nil)
}
