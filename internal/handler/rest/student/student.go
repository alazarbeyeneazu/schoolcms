package student

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

type student struct {
	studentModule  module.Student
	log            logger.Logger
	ContextTimeOut time.Duration
}

func Init(studentModule module.Student, log logger.Logger, contextTimeOut time.Duration) rest.Student {
	return &student{
		studentModule:  studentModule,
		log:            log,
		ContextTimeOut: contextTimeOut,
	}
}
func (s *student) CreateStudent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, s.ContextTimeOut)
	defer cancel()
	var std dto.Student
	if err := c.ShouldBind(&std); err != nil {
		err = errors.ErrValidationError.Wrap(err, "error while binding user to dto.Student")
		s.log.Error(ctx, "error while binding dto.Student ", zap.Error(err))
		_ = c.Error(err)
		return
	}
	studentRespo, err := s.studentModule.CreateStudent(ctx, std)
	if err != nil {
		_ = c.Error(err)
		return
	}
	studentRespo.ID = uuid.Nil
	response.SendSuccessResponse(c, http.StatusCreated, studentRespo, nil)
}
