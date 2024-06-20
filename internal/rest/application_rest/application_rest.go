package applicationrest

import (
	"net/http"

	"github.com/st-application-bot/pkg/json"
	"github.com/st-application-bot/usecase/application"
)

type ApplicationHandler struct {
	service *application.Service
}

func NewApplicationHandler(applicationStorage application.ApplicationStorage) *ApplicationHandler {
	return &ApplicationHandler{
		service: application.NewService(applicationStorage),
	}
}

func (a *ApplicationHandler) HandleGetAllApplications(w http.ResponseWriter, r *http.Request) {
	applications, err := a.service.GetAllApplications(r.Context())
	if err != nil {
		json.ServeJSONError(r.Context(), w, err)
		return
	}

	json.ServeJSONBody(r.Context(), w, applications, http.StatusOK)
}
