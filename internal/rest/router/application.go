package router

import (
	"net/http"

	"github.com/gorilla/mux"
	applicationrest "github.com/st-application-bot/internal/rest/application_rest"
	"github.com/st-application-bot/usecase/application"
	"go.uber.org/zap"
)

func MountApplicationRouter(root *mux.Router, applicationStorage application.ApplicationStorage, logger *zap.Logger) {
	r := root.PathPrefix("/applications").Subrouter()

	h := applicationrest.NewApplicationHandler(applicationStorage, logger)

	r.HandleFunc("/", h.HandleGetAllApplications).Methods(http.MethodGet, http.MethodOptions)
}
