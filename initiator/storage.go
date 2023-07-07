package initiator

import (
	persistencedb "schoolcms/internal/constant/persistenceDB"
	"schoolcms/internal/storage"
	"schoolcms/internal/storage/user"

	"schoolcms/platform/logger"
)

type Persistence struct {
	User storage.User
}

func InitPersistence(db persistencedb.PersistenceDB, log logger.Logger) Persistence {
	return Persistence{
		User: user.Init(db, log.Named("user persistant")),
	}
}
