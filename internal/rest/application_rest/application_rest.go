package applicationrest

import (
	"net/http"

	"github.com/NicoNex/echotron/v3"
	"github.com/st-application-bot/domain"
	"github.com/st-application-bot/internal/rest/applicationbot"
	"github.com/st-application-bot/internal/rest/utils"
	"github.com/st-application-bot/pkg/json"
	"github.com/st-application-bot/usecase/application"
	"go.uber.org/zap"
)

type ApplicationHandler struct {
	logger  *zap.Logger
	service *application.Service
	api     *echotron.API
}

func NewApplicationHandler(applicationStorage application.ApplicationStorage, logger *zap.Logger, api *echotron.API) *ApplicationHandler {
	return &ApplicationHandler{
		logger:  logger,
		service: application.NewService(applicationStorage),
		api:     api,
	}
}

func (a *ApplicationHandler) HandleGetUnresolvedApplications(w http.ResponseWriter, r *http.Request) {
	applications, err := a.service.GetUnresolvedApplications(r.Context())
	if err != nil {
		a.logger.Error(err.Error())
		json.ServeJSONError(r.Context(), w, err)
		return
	}

	json.ServeJSONBody(r.Context(), w, applications, http.StatusOK)
}

func (a *ApplicationHandler) HandleApproveApplication(w http.ResponseWriter, r *http.Request) {
	applicationID, err := utils.GetApplicationIDMuxVar(r)
	if err != nil {
		a.logger.Error(err.Error())
		json.ServeJSONError(r.Context(), w, err)
		return
	}

	application, err := a.service.ApproveApplication(r.Context(), applicationID)
	if err != nil {
		a.logger.Error(err.Error())
		json.ServeJSONError(r.Context(), w, err)
		return
	}

	_, err = a.api.SendMessage(applicationbot.ApproveText, application.TelegramUserID, applicationbot.OptionsEmpty)
	if err != nil {
		a.logger.Error(err.Error())
		json.ServeJSONError(r.Context(), w, domain.ErrUnableToSendMessage)
		return
	}

	json.ServeJSONBody(r.Context(), w, application, http.StatusOK)
}

func (a *ApplicationHandler) HandleRejectApplication(w http.ResponseWriter, r *http.Request) {
	applicationID, err := utils.GetApplicationIDMuxVar(r)
	if err != nil {
		a.logger.Error(err.Error())
		json.ServeJSONError(r.Context(), w, err)
		return
	}

	application, err := a.service.RejectApplication(r.Context(), applicationID)
	if err != nil {
		a.logger.Error(err.Error())
		json.ServeJSONError(r.Context(), w, err)
		return
	}

	_, err = a.api.SendMessage(applicationbot.RejectText, application.TelegramUserID, applicationbot.OptionsEmpty)
	if err != nil {
		a.logger.Error(err.Error())
		json.ServeJSONError(r.Context(), w, domain.ErrUnableToSendMessage)
		return
	}

	json.ServeJSONBody(r.Context(), w, application, http.StatusOK)
}
