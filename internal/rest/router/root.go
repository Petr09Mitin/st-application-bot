package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/NicoNex/echotron/v3"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/st-application-bot/internal/config"
	"github.com/st-application-bot/internal/rest/middleware"
	"github.com/st-application-bot/usecase/application"
	"go.uber.org/zap"
)

func MountRootRouter(rootRouter *mux.Router, applicationStorage application.ApplicationStorage, zapLogger *zap.Logger) {
	r := rootRouter.PathPrefix("/api/v1").Subrouter()

	tgBotToken := os.Getenv(config.TGBotTokenEnvName)
	api := echotron.NewAPI(tgBotToken)

	MountApplicationRouter(r, applicationStorage, zapLogger, &api)

	r.Use(middleware.Recovery)

	handler := cors.New(cors.Options{
		AllowedOrigins:   config.ALLOWED_ORIGINS,
		AllowedMethods:   config.ALLOWED_METHODS,
		AllowedHeaders:   config.ALLOWED_HEADERS,
		AllowCredentials: true,
	}).Handler(rootRouter)

	appPort := os.Getenv(config.AppPortEnvName)
	zapLogger.Info(fmt.Sprintf("started on port %s", appPort))
	err := http.ListenAndServe(appPort, handler)
	if err != nil {
		zapLogger.Error(err.Error())
		return
	}
	zapLogger.Info(fmt.Sprintf("app stopped listening on port %s", appPort))
}
