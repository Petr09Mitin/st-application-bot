package router

import (
	"net/http"

	"github.com/NicoNex/echotron/v3"
	"github.com/gorilla/mux"
	applicationrest "github.com/st-application-bot/internal/rest/application_rest"
	"github.com/st-application-bot/usecase/application"
	"go.uber.org/zap"
)

func MountApplicationRouter(root *mux.Router, applicationStorage application.ApplicationStorage, logger *zap.Logger, api *echotron.API) {
	r := root.PathPrefix("/applications").Subrouter()

	h := applicationrest.NewApplicationHandler(applicationStorage, logger, api)

	r.HandleFunc("/", h.HandleGetUnresolvedApplications).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/{applicationID:[0-9]+}/approve", h.HandleApproveApplication).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/{applicationID:[0-9]+}/reject", h.HandleRejectApplication).Methods(http.MethodPost, http.MethodOptions)
}
