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
	InitialState                ConversationState = 0
	WaitingForName              ConversationState = 1
	WaitingForPhone             ConversationState = 2
	WaitingForDistrict          ConversationState = 3
	WaitingForPurpose           ConversationState = 4
	WaitingForInformationSource ConversationState = 5
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
		_, err := b.SendMessage(NoMessageText, b.chatID, OptionsWithStartKeyboard)
		if err != nil {
			return
		}
	}

	messageText := ""

	if update.Message != nil {
		messageText, err = b.ValidateTextMessage(update)
		if err != nil {
			b.logger.Error(err.Error())
			return
		}
	} else if update.CallbackQuery != nil {
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
		_, err = b.SendMessage(IntroductionText, b.chatID, OptionsWithCancelKeyboard)
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
		_, err = b.SendMessage(PhonePromptText, b.chatID, OptionsWithCancelKeyboard)
		if err != nil {
			return
		}
		b.currentState = WaitingForPhone
	case WaitingForPhone:
		if !validators.IsValidPhone(messageText) {
			_, err = b.SendMessage(IncorrectPhoneFormatText, b.chatID, OptionsWithCancelKeyboard)
			if err != nil {
				return
			}

			return
		}

		b.application.PhoneNumber = messageText
		_, err = b.SendMessage(DistrictPromptText, b.chatID, OptionsWithCancelKeyboard)
		if err != nil {
			return
		}
		b.currentState = WaitingForDistrict
	case WaitingForDistrict:
		b.application.District = messageText
		_, err = b.SendMessage(PurposePromptText, b.chatID, OptionsWithCancelKeyboard)
		if err != nil {
			return
		}
		b.currentState = WaitingForPurpose
	case WaitingForPurpose:
		b.application.PurposeOfApplication = messageText
		_, err = b.SendMessage(InformationSourcePromptText, b.chatID, OptionsWithCancelKeyboard)
		if err != nil {
			return
		}
		b.currentState = WaitingForInformationSource
	case WaitingForInformationSource:
		b.application.SourceOfInformation = messageText
		_, err = b.service.CreateApplication(context.Background(), b.application)
		if err != nil {
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
