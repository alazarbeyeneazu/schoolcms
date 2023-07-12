package test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"schoolcms/initiator"

	persistencedb "schoolcms/internal/constant/persistenceDB"
	"schoolcms/internal/handler/middleware"
	"schoolcms/platform/logger"
	"schoolcms/platform/utils"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

type PgxSeeder struct {
	DB *pgxpool.Pool
}

type PgxSeederTable struct {
	TableName      string
	PrimaryKeyName string
	Data           any
}

func (p *PgxSeeder) Feed(data interface{}) error {
	table, ok := data.(PgxSeederTable)
	if !ok {
		return fmt.Errorf("data must be a type of PgxSeederTable")
	}
	js, err := json.Marshal(table.Data)
	if err != nil {
		return err
	}

	var dataMap map[string]any
	err = json.Unmarshal(js, &dataMap)
	if err != nil {
		return err
	}

	fields := make([]string, 0, len(dataMap))
	values := make([]string, 0, len(dataMap))

	for k, v := range dataMap {
		fields = append(fields, k)
		values = append(values, fmt.Sprintf("'%v'", v))
	}

	_, err = p.DB.Exec(context.Background(), fmt.Sprintf("INSERT INTO %s (%s) values(%s)",
		table.TableName,
		strings.Join(fields, ","),
		strings.Join(values, ",")))
	if err != nil {
		return err
	}

	return nil
}

func (p *PgxSeeder) Starve(data interface{}) error {
	table, ok := data.(PgxSeederTable)
	if !ok {
		return fmt.Errorf("data must be a type of PgxSeederTable")
	}

	js, err := json.Marshal(table.Data)
	if err != nil {
		return err
	}

	var dataMap map[string]any
	err = json.Unmarshal(js, &dataMap)
	if err != nil {
		return err
	}

	_, err = p.DB.Exec(context.Background(),
		fmt.Sprintf("DELETE FROM %s WHERE %s='%s'",
			table.TableName,
			table.PrimaryKeyName,
			dataMap[table.PrimaryKeyName]))

	return err
}

type Instance struct {
	DBName    string
	Server    *gin.Engine
	Module    initiator.Module
	Logger    logger.Logger
	Conn      *pgxpool.Pool
	PersistDB persistencedb.PersistenceDB
}

func Initiate(path string) (Instance, func() error) {
	log := logger.New(initiator.InitLogger())
	log.Info(context.Background(), "logger initialized")

	configName := "config"
	if name := os.Getenv("CONFIG_NAME"); name != "" {
		configName = name
		log.Info(context.Background(), fmt.Sprintf("config name is set to %s", configName))
	} else {
		log.Info(context.Background(), "using default config name 'config'")
	}
	log.Info(context.Background(), "initializing config")
	initiator.InitConfig(configName, path+"config", log.GetZapLogger())
	log.Info(context.Background(), "config initialized")

	log.Info(context.Background(), "initializing database")
	pgxConn := initiator.InitDB(viper.GetString("database.url"), log)
	log.Info(context.Background(), "database initialized")
	// create database for this specific test
	log.Info(context.Background(), "initializing test database")
	dbName := utils.GenerateCustomRandomString(utils.SmallLetters, 10)
	_, err := pgxConn.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %s", dbName))
	if err != nil {
		log.Fatal(context.Background(), "failed to create test database")
	}
	log.Info(context.Background(), "initializing kafka dailer")

	testConnURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("test_database.user"),
		viper.GetString("test_database.password"),
		viper.GetString("test_database.host"),
		viper.GetString("test_database.port"),
		dbName)
	testConn := initiator.InitDB(testConnURL, log)
	log.Info(context.Background(), "test database initialized")

	if viper.GetBool("migration.active") {
		log.Info(context.Background(), "initializing migration")
		m := initiator.InitiateMigration(path+viper.GetString("migration.path"), testConnURL, log)
		initiator.UpMigration(m, log)
		log.Info(context.Background(), "migration initialized")
	}

	log.Info(context.Background(), "initializing persistence layer")
	persistDB := persistencedb.New(testConn, log)
	persistence := initiator.InitPersistence(persistDB, log)
	log.Info(context.Background(), "persistence layer initialized")

	log.Info(context.Background(), "initializing module")
	module := initiator.InitModule(persistence, log)
	log.Info(context.Background(), "module initialized")

	log.Info(context.Background(), "initializing handler")
	handler := initiator.InitHandler(module, log, viper.GetDuration("server.timeout"))
	log.Info(context.Background(), "handler initialized")

	log.Info(context.Background(), "initializing server")
	server := gin.New()
	server.Use(middleware.GinLogger(log))
	server.Use(ginzap.RecoveryWithZap(log.GetZapLogger().Named("gin.recovery"), true))

	server.Use(middleware.ErrorHandler())
	log.Info(context.Background(), "server initialized")

	log.Info(context.Background(), "initializing router")
	v1 := server.Group("/v1")
	state := initiator.InitState()
	initiator.InitRouter(v1, handler, log, state.AuthDomains)
	log.Info(context.Background(), "router initialized")

	return Instance{
			DBName:    dbName,
			Server:    server,
			Module:    module,
			Logger:    log,
			Conn:      testConn,
			PersistDB: persistDB,
		}, func() error {
			_, err = pgxConn.Exec(context.Background(), fmt.Sprintf("DROP DATABASE %s cascade", dbName))
			if err != nil {
				return err
			}

			testConn.Close()
			log.Info(context.Background(), fmt.Sprintf("dropped test database %s", dbName))
			return nil
		}
}
