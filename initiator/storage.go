package initiator

import (
	persistencedb "schoolcms/internal/constant/persistenceDB"

	"schoolcms/internal/storage"
	"schoolcms/internal/storage/family"
	"schoolcms/internal/storage/grade"
	"schoolcms/internal/storage/school"
	"schoolcms/internal/storage/student"
	"schoolcms/internal/storage/teacher"
	"schoolcms/internal/storage/user"
	"schoolcms/platform/logger"
)

type Persistence struct {
	User    storage.User
	School  storage.School
	Teacher storage.Teacher
	Grade   storage.Grade
	Student storage.Student
	Family  storage.Family
}

func InitPersistence(db persistencedb.PersistenceDB, log logger.Logger) Persistence {
	return Persistence{
		User:    user.Init(db, log.Named("user persistent")),
		School:  school.Init(db, log.Named("school persistent")),
		Teacher: teacher.Init(db, log.Named("teacher persistent")),
		Grade:   grade.Init(db, log.Named("grade persistent")),
		Student: student.Init(db, log.Named("student persistent")),
		Family:  family.Init(db, log.Named("family persistent")),
	}
}
