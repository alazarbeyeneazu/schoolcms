package initiator

import (
	persistencedb "loyalty/internal/constant/persistenceDB"
	"loyalty/platform/logger"
)

type Persistence struct {
}

func InitPersistence(db persistencedb.PersistenceDB, log logger.Logger) Persistence {
	return Persistence{}
}
