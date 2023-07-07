package initiator

import (
	"schoolcms/internal/module"
	"schoolcms/internal/module/user"
	"schoolcms/platform/logger"
)

type Module struct {
	User module.User
}

func InitModule(persistence Persistence,
	log logger.Logger) Module {

	return Module{
		User: user.Init(log.Named("user module"), persistence.User),
	}
}
