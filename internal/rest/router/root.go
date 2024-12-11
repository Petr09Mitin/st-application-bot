package router

import (
	"fmt"
	"net/http"

	"github.com/NicoNex/echotron/v3"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/st-application-bot/internal/config"
	"github.com/st-application-bot/internal/rest/middleware"
	"github.com/st-application-bot/usecase/application"
	"go.uber.org/zap"
)

func MountRootRouter(rootRouter *mux.Router, applicationStorage application.ApplicationStorage, zapLogger *zap.Logger, cfg *config.Config) {
	r := rootRouter.PathPrefix(cfg.Server.APIPrefix).Subrouter()

	tgBotToken := cfg.Telegram.BotToken
	api := echotron.NewAPI(tgBotToken)

	MountApplicationRouter(r, applicationStorage, zapLogger, &api)

	r.Use(middleware.Recovery)

	handler := cors.New(cors.Options{
		AllowedOrigins:   cfg.Server.CORS.AllowedOrigins,
		AllowedMethods:   cfg.Server.CORS.AllowedMethods,
		AllowedHeaders:   cfg.Server.CORS.AllowedHeaders,
		AllowCredentials: cfg.Server.CORS.AllowedCredentials,
	}).Handler(rootRouter)

	appPort := cfg.Server.Port
	zapLogger.Info(fmt.Sprintf("started on port %s", appPort))
	err := http.ListenAndServe(appPort, handler)
	if err != nil {
		zapLogger.Error(err.Error())
		return
	}
	zapLogger.Info(fmt.Sprintf("app stopped listening on port %s", appPort))
}
