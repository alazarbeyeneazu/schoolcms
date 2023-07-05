package initiator

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	persistencedb "loyalty/internal/constant/persistenceDB"
	"loyalty/internal/handler/middleware"
	"loyalty/platform/logger"
	"loyalty/platform/routine"
	"loyalty/platform/wait"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Initiate() {
	sampleLogger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf(`{"level":"fatal","msg":"failed to initialize sample logger: %v"}
`, err)
		os.Exit(1)
	}

	sampleLogger.Info("initializing config")
	configName := "config"
	if name := os.Getenv("CONFIG_NAME"); name != "" {
		configName = name
		sampleLogger.Info(fmt.Sprintf("config name is set to %s", configName))
	} else {
		sampleLogger.Info("using default config name 'config'")
	}
	InitConfig(configName, "config", sampleLogger)
	sampleLogger.Info("config initialized")

	log := logger.New(InitLogger())
	log.Info(context.Background(), "logger initialized")

	log.Info(context.Background(), "initializing database")
	pgxConn := InitDB(viper.GetString("database.url"), log)
	log.Info(context.Background(), "database initialized")

	if viper.GetBool("migration.active") {
		log.Info(context.Background(), "initializing migration")
		m := InitiateMigration(viper.GetString("migration.path"), viper.GetString("database.url"), log)
		UpMigration(m, log)
		log.Info(context.Background(), "migration initialized")
	}

	log.Info(context.Background(), "initializing persistence layer")
	persistence := InitPersistence(persistencedb.New(pgxConn, log), log)
	log.Info(context.Background(), "persistence layer initialized")

	log.Info(context.Background(), "initializing module")
	module := InitModule(persistence, log)
	log.Info(context.Background(), "module initialized")

	log.Info(context.Background(), "initializing handler")
	handler := InitHandler(module, log, viper.GetDuration("server.timeout"))
	log.Info(context.Background(), "handler initialized")

	log.Info(context.Background(), "initializing server")
	server := gin.New()
	gin.SetMode(gin.ReleaseMode)

	server.Use(middleware.GinLogger(log.Named("gin")))
	server.Use(ginzap.RecoveryWithZap(log.GetZapLogger().Named("gin.recovery"), true))

	server.Use(middleware.ErrorHandler())
	server.Use(InitCORS())
	log.Info(context.Background(), "server initialized")

	log.Info(context.Background(), "initializing router")

	InitRouter(server.Group("/api/v1"), handler, module, log)
	log.Info(context.Background(), "router initialized")
	srv := &http.Server{
		Addr:              viper.GetString("server.host") + ":" + viper.GetString("server.port"),
		ReadHeaderTimeout: viper.GetDuration("read_header_timeout"),
		Handler:           server,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	routine.ExecuteRoutine(context.Background(), routine.Routine{
		Name:   "server",
		NoWait: true,
		Operation: func(ctx context.Context, log logger.Logger) {
			log.Info(ctx, "server started",
				zap.String("host", viper.GetString("server.host")),
				zap.Int("port", viper.GetInt("server.port")))
			log.Info(ctx, fmt.Sprintf("server stopped with error %v", srv.ListenAndServe()))
		},
	}, log)

	sig := <-quit
	log.Info(context.Background(), fmt.Sprintf("server shutting down with signal %v", sig))
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("server.timeout"))
	defer cancel()

	log.Info(ctx, "shutting down server")

	err = srv.Shutdown(ctx)

	// wait for separate routines
	wait.RoutineWaitGroup.Wait()

	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("error while shutting down server: %v", err))
	} else {
		log.Info(context.Background(), "server shutdown complete")
	}
}
