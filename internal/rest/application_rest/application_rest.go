package applicationrest

import (
	"net/http"

	"github.com/st-application-bot/pkg/json"
	"github.com/st-application-bot/usecase/application"
	"go.uber.org/zap"
)

type ApplicationHandler struct {
	logger  *zap.Logger
	service *application.Service
}

func NewApplicationHandler(applicationStorage application.ApplicationStorage, logger *zap.Logger) *ApplicationHandler {
	return &ApplicationHandler{
		logger:  logger,
		service: application.NewService(applicationStorage),
	}
}

func (a *ApplicationHandler) HandleGetAllApplications(w http.ResponseWriter, r *http.Request) {
	applications, err := a.service.GetAllApplications(r.Context())
	if err != nil {
		a.logger.Error(err.Error())
		json.ServeJSONError(r.Context(), w, err)
		return
	}

	json.ServeJSONBody(r.Context(), w, applications, http.StatusOK)
}
