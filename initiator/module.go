package initiator

import (
	"loyalty/platform/logger"
)

type Module struct {
}

func InitModule(persistence Persistence,
	log logger.Logger) Module {

	return Module{}
}
