package teacher

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

type teacher struct {
	TeacherModule  module.Teacher
	ContextTimeOut time.Duration
	log            logger.Logger
}

func Init(ctx context.Context, teacherModule module.Teacher, log logger.Logger, contextTimeOut time.Duration) rest.Teacher {
	return &teacher{
		TeacherModule:  teacherModule,
		ContextTimeOut: contextTimeOut,
		log:            log,
	}
}
func (t *teacher) CreateTeacher(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, t.ContextTimeOut)
	defer cancel()
	var tchr dto.Teacher
	if err := c.ShouldBind(&tchr); err != nil {
		err = errors.ErrValidationError.Wrap(err, "invalid input")
		t.log.Error(ctx, "unable to bind user input to dto.Teacher object", zap.Error(err))
		_ = c.Error(err)
		return
	}
	teacher, err := t.TeacherModule.CreateTeacher(ctx, tchr)
	if err != nil {
		_ = c.Error(err)
		return
	}
	teacher.ID = uuid.Nil
	response.SendSuccessResponse(c, http.StatusCreated, teacher, nil)

}
