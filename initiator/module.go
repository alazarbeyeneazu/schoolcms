package initiator

import (
	"context"
	"schoolcms/internal/module"
	"schoolcms/internal/module/grade"
	"schoolcms/internal/module/school"
	"schoolcms/internal/module/teacher"
	"schoolcms/internal/module/user"

	"schoolcms/platform/logger"
)

type Module struct {
	User    module.User
	School  module.School
	Teacher module.Teacher
	Grade   module.Grade
}

func InitModule(persistence Persistence,
	log logger.Logger) Module {

	return Module{
		User:    user.Init(log.Named("user module"), persistence.User),
		School:  school.Init(persistence.School, log.Named("school module")),
		Teacher: teacher.Init(log.Named("teacher module"), persistence.Teacher),
		Grade:   grade.Init(context.Background(), persistence.Grade, log.Named("grade module")),
	}
}
