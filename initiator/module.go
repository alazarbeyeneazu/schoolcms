package initiator

import (
	"schoolcms/internal/module"
	"schoolcms/internal/module/school"
	"schoolcms/internal/module/user"

	"schoolcms/platform/logger"
)

type Module struct {
	User   module.User
	School module.School
}

func InitModule(persistence Persistence,
	log logger.Logger) Module {

	return Module{
		User:   user.Init(log.Named("user module"), persistence.User),
		School: school.Init(persistence.School, log.Named("school module")),
	}
}
