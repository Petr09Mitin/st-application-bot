package main

import (
	"fmt"
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
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	zapLogger, err := logger.NewZapLogger(cfg.Logger.OutputPaths)
	if err != nil {
		fmt.Println(err)
		return
	}

	token := cfg.Telegram.BotToken
	if token == "" {
		zapLogger.Error(domain.ErrInvalidEnv.Error())
		return
	}

	pgConnStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		cfg.Postgres.User,
		cfg.Postgres.DBName,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
	)
	db, err := postgres.NewPool(pgConnStr)
	if err != nil {
		zapLogger.Error(err.Error())
		return
	}
	defer db.Close()

	applicationStorage := postgres.NewApplication(db)

	rootRouter := mux.NewRouter()

	if !cfg.Server.IsAPIDisabled {
		go router.MountRootRouter(rootRouter, applicationStorage, zapLogger, cfg)
	}

	adminChatID := cfg.Telegram.AdminChatID

	dsp := echotron.NewDispatcher(token, applicationbot.GetApplicationBotConstructor(token, adminChatID, applicationStorage, zapLogger, cfg))
	for {
		zapLogger.Info("Started polling...")
		err = dsp.Poll()
		if err != nil {
			zapLogger.Error(err.Error())
		}
		time.Sleep(cfg.Telegram.PollingTimeout)
	}
}
