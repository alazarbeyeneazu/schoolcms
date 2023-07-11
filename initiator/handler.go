package initiator

import (
	"context"
	"schoolcms/internal/handler/rest"
	"schoolcms/internal/handler/rest/grade"
	"schoolcms/internal/handler/rest/school"
	"schoolcms/internal/handler/rest/student"
	"schoolcms/internal/handler/rest/teacher"
	"schoolcms/internal/handler/rest/user"
	"schoolcms/platform/logger"
	"time"
)

type Handler struct {
	User    rest.User
	School  rest.School
	Teacher rest.Teacher
	Grade   rest.Grade
	Student rest.Student
}

func InitHandler(ctx context.Context, module Module, log logger.Logger, timeout time.Duration) Handler {
	return Handler{
		User:    user.Init(ctx, module.User, log.Named("user handler"), timeout),
		School:  school.Init(ctx, module.School, log.Named("school handler"), timeout),
		Teacher: teacher.Init(ctx, module.Teacher, log.Named("teacher handler"), timeout),
		Grade:   grade.Init(ctx, module.Grade, timeout, log.Named("grade handler")),
		Student: student.Init(module.Student, log.Named("student handler"), timeout),
	}
}
