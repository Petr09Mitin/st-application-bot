package applicationbot

import (
	"context"
	"strings"

	"github.com/NicoNex/echotron/v3"
	"github.com/st-application-bot/domain"
	"github.com/st-application-bot/pkg/validators"
	"github.com/st-application-bot/usecase/application"
	"go.uber.org/zap"
)

type ConversationState int

var (
	InitialState                   ConversationState = 0
	WaitingForName                 ConversationState = 1
	WaitingForDistrict             ConversationState = 2
	WaitingForIsSevChelParticipant ConversationState = 3
	WaitingForPurpose              ConversationState = 4
)

type ApplicationBot struct {
	chatID       int64
	service      *application.Service
	application  *domain.Application
	currentState ConversationState
	logger       *zap.Logger
	echotron.API
}

func GetApplicationBotConstructor(token string, storage application.ApplicationStorage, logger *zap.Logger) echotron.NewBotFn {
	return func(chatID int64) echotron.Bot {
		return &ApplicationBot{
			chatID,
			application.NewService(storage),
			&domain.Application{},
			0,
			logger,
			echotron.NewAPI(token),
		}
	}
}

func (b *ApplicationBot) Update(update *echotron.Update) {
	var err error
	if update.Message == nil && update.CallbackQuery == nil {
		return
	}

	messageText := ""

	if update.Message != nil {
		if update.Message.Chat.Type != "private" {
			return
		}

		messageText, err = b.ValidateTextMessage(update)
		if err != nil {
			b.logger.Error(err.Error())
			return
		}
	} else if update.CallbackQuery != nil {
		if update.CallbackQuery.Message.Chat.Type != "private" {
			return
		}

		messageText, err = b.ValidateCallbackQuery(update)
		if err != nil {
			b.logger.Error(err.Error())
			return
		}
	}

	err = b.HandleTextMessage(messageText)
	if err != nil {
		b.logger.Error(err.Error())
		return
	}
}

func (b *ApplicationBot) ValidateTextMessage(update *echotron.Update) (messageText string, err error) {
	if update.Message.From == nil {
		b.currentState = InitialState
		_, err = b.SendMessage(OnlyRealUsersText, b.chatID, OptionsEmpty)
		if err != nil {
			return
		}
	}

	b.application.TelegramUserID = update.Message.From.ID

	if update.Message.From.Username == "" {
		b.currentState = InitialState
		_, err = b.SendMessage(OnlyUsersWithUsernameText, b.chatID, OptionsEmpty)
		if err != nil {
			return
		}

		return
	}

	b.application.TelegramUsername = update.Message.From.Username

	messageText = strings.TrimSpace(update.Message.Text)
	if messageText == "" {
		_, err = b.SendMessage(EnterUnemptyMessageText, b.chatID, OptionsEmpty)
		if err != nil {
			return
		}
	}

	return
}

func (b *ApplicationBot) ValidateCallbackQuery(update *echotron.Update) (messageText string, err error) {
	if update.CallbackQuery.From == nil {
		b.currentState = InitialState
		_, err = b.SendMessage(OnlyRealUsersText, b.chatID, OptionsEmpty)
		if err != nil {
			return
		}
	}

	messageText = strings.TrimSpace(update.CallbackQuery.Data)
	if messageText == "" {
		_, err = b.SendMessage(EnterUnemptyMessageText, b.chatID, OptionsEmpty)
		if err != nil {
			return
		}
	}

	return
}

func (b *ApplicationBot) HandleTextMessage(messageText string) (err error) {
	switch messageText {
	case StartCommand:
		_, err = b.service.GetApplicationByTgUserID(context.Background(), b.application.TelegramUserID)
		if err != domain.ErrApplicationNotFound {
			if err != nil {
				_, err = b.SendMessage(InternalErrorText, b.chatID, OptionsWithStartKeyboard)
				return
			}

			_, err = b.SendMessage(ApplicationAlreadyExistsText, b.chatID, OptionsEmpty)
			return
		}

		_, err = b.SendMessage(IntroductionText, b.chatID, OptionsEmpty)
		if err != nil {
			return
		}

		_, err = b.SendMessage(FullNamePromptText, b.chatID, OptionsWithCancelKeyboard)
		if err != nil {
			return
		}

		b.currentState = WaitingForName
	case CancelCommand:
		_, err = b.SendMessage(CancelText, b.chatID, OptionsWithStartKeyboard)
		if err != nil {
			return
		}
		b.currentState = InitialState
	default:
		err = b.HandlePlainTextMessage(messageText)
		if err != nil {
			return
		}
	}

	return
}

func (b *ApplicationBot) HandlePlainTextMessage(messageText string) (err error) {
	switch b.currentState {
	case InitialState:
		_, err = b.SendMessage(PressStartText, b.chatID, OptionsWithStartKeyboard)
		if err != nil {
			return
		}
	case WaitingForName:
		if !validators.IsValidFullName(messageText) {
			_, err = b.SendMessage(IncorrectFullNameFormatText, b.chatID, OptionsWithCancelKeyboard)
			if err != nil {
				return
			}

			return
		}

		b.application.FullName = messageText

		_, err = b.SendMessage(DistrictPromptText, b.chatID, OptionsWithCancelKeyboard)
		if err != nil {
			return
		}

		b.currentState = WaitingForDistrict
	case WaitingForDistrict:
		b.application.District = messageText

		_, err = b.SendMessage(IsSevChelParticipantPromptText, b.chatID, OptionsWithYesOrNoWithCancelKeyboard)
		if err != nil {
			return
		}

		b.currentState = WaitingForIsSevChelParticipant
	case WaitingForIsSevChelParticipant:
		b.application.IsSevChelParticipant = messageText == YesCommand

		_, err = b.service.CreateApplication(context.Background(), b.application)
		if err != nil {
			if err == domain.ErrApplicationAlreadyExists {
				_, err = b.SendMessage(ApplicationAlreadyExistsText, b.chatID, OptionsEmpty)
				return
			}

			_, err = b.SendMessage(InternalErrorText, b.chatID, OptionsWithStartKeyboard)
			return
		}

		_, err = b.SendMessage(ConclusionText, b.chatID, OptionsEmpty)
		if err != nil {
			return
		}
		b.currentState = InitialState
	default:
		b.currentState = InitialState
		_, err = b.SendMessage(PressStartText, b.chatID, OptionsWithStartKeyboard)
		if err != nil {
			return
		}
	}

	return
}
