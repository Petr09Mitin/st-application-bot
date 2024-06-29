package applicationrest

import (
	"net/http"
	"sync"

	"github.com/NicoNex/echotron/v3"
	"github.com/st-application-bot/domain"
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

	targetChats, err := a.service.GetAllTargetChats(r.Context())
	if err != nil {
		a.logger.Error(err.Error())
		json.ServeJSONError(r.Context(), w, domain.ErrUnableToApproveApplication)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(targetChats))

	for _, chat := range targetChats {
		go func(chat *domain.TargetChat) {
			defer wg.Done()
			_, err = a.api.ApproveChatJoinRequest(chat.TelegramChatID, application.TelegramUserID)
			if err != nil {
				a.logger.Error(err.Error())
				return
			}

		}(chat)
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

	targetChats, err := a.service.GetAllTargetChats(r.Context())
	if err != nil {
		a.logger.Error(err.Error())
		json.ServeJSONError(r.Context(), w, domain.ErrUnableToApproveApplication)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(targetChats))

	for _, chat := range targetChats {
		go func(chat *domain.TargetChat) {
			defer wg.Done()
			a.logger.Info("tg chat id: ", zap.Int64("tg chat id", chat.TelegramChatID))
			_, err = a.api.DeclineChatJoinRequest(chat.TelegramChatID, application.TelegramUserID)
			if err != nil {
				a.logger.Error(err.Error())
				return
			}

		}(chat)
	}

	json.ServeJSONBody(r.Context(), w, application, http.StatusOK)
}
