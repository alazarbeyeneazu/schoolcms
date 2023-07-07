package persistencedb

import (
	"schoolcms/internal/constant/model/db"
	"schoolcms/platform/logger"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PersistenceDB struct {
	*db.Queries
	pool *pgxpool.Pool
	log  logger.Logger
}

type Sibling string

func New(pool *pgxpool.Pool, log logger.Logger) PersistenceDB {
	return PersistenceDB{
		Queries: db.New(pool),
		pool:    pool,
		log:     log,
	}
}
