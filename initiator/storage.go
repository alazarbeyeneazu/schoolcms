package initiator

import (
	persistencedb "schoolcms/internal/constant/persistenceDB"

	"schoolcms/internal/storage"
	"schoolcms/internal/storage/grade"
	"schoolcms/internal/storage/school"
	"schoolcms/internal/storage/teacher"
	"schoolcms/internal/storage/user"
	"schoolcms/platform/logger"
)

type Persistence struct {
	User    storage.User
	School  storage.School
	Teacher storage.Teacher
	Grade   storage.Grade
}

func InitPersistence(db persistencedb.PersistenceDB, log logger.Logger) Persistence {
	return Persistence{
		User:    user.Init(db, log.Named("user persistant")),
		School:  school.Init(db, log.Named("school persistant")),
		Teacher: teacher.Init(db, log.Named("teacher persistant")),
		Grade:   grade.Init(db, log.Named("grade persistant")),
	}
}
