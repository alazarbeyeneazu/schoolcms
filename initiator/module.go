package initiator

import (
	"schoolcms/internal/module"
	"schoolcms/internal/module/family"
	"schoolcms/internal/module/grade"
	"schoolcms/internal/module/school"
	"schoolcms/internal/module/student"
	"schoolcms/internal/module/teacher"
	"schoolcms/internal/module/user"

	"schoolcms/platform/logger"
)

type Module struct {
	User    module.User
	School  module.School
	Teacher module.Teacher
	Grade   module.Grade
	Student module.Student
	Family  module.Family
}

func InitModule(persistence Persistence,
	log logger.Logger) Module {

	return Module{
		User:    user.Init(log.Named("user module"), persistence.User),
		School:  school.Init(persistence.School, log.Named("school module")),
		Teacher: teacher.Init(log.Named("teacher module"), persistence.Teacher),
		Grade:   grade.Init(persistence.Grade, log.Named("grade module")),
		Student: student.Init(persistence.Student, log.Named("student persistent")),
		Family:  family.Init(persistence.Family, log.Named("family module")),
	}
}
