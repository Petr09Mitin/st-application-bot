package main

import (
	"fmt"
	"os"
	"time"

	"github.com/NicoNex/echotron/v3"
	"github.com/gorilla/mux"
	"github.com/st-application-bot/domain"
	"github.com/st-application-bot/internal/config"
	"github.com/st-application-bot/internal/repository/postgres"
	"github.com/st-application-bot/internal/rest/applicationbot"
	"github.com/st-application-bot/internal/rest/router"
	"github.com/st-application-bot/pkg/logger"
)

func main() {
	zapLogger, err := logger.NewZapLogger(logger.DefaultOutputPaths)
	if err != nil {
		fmt.Println(err)
		return
	}

	token := os.Getenv(config.TGBotTokenEnvName)
	if token == "" {
		zapLogger.Error(domain.ErrInvalidEnv.Error())
		return
	}

	pgConnStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv(config.PgUserEnvName),
		os.Getenv(config.PgDBNameEnvName),
		os.Getenv(config.PgPasswordEnvName),
		os.Getenv(config.PgHostEnvName),
		os.Getenv(config.PgPortEnvName),
	)
	db, err := postgres.NewPool(pgConnStr)
	if err != nil {
		zapLogger.Error(err.Error())
		return
	}
	defer db.Close()

	applicationStorage := postgres.NewApplication(db)

	rootRouter := mux.NewRouter()

	go router.MountRootRouter(rootRouter, applicationStorage, zapLogger)

	dsp := echotron.NewDispatcher(token, applicationbot.GetApplicationBotConstructor(token, applicationStorage, zapLogger))
	for {
		zapLogger.Info("Started polling...")
		err = dsp.Poll()
		if err != nil {
			zapLogger.Error(err.Error())
		}
		time.Sleep(3 * time.Second)
	}
}
