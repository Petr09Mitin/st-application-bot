package applicationbot

import (
	"context"
	"strings"

	"github.com/NicoNex/echotron/v3"
	"github.com/st-application-bot/domain"
	"github.com/st-application-bot/internal/config"
	"github.com/st-application-bot/pkg/validators"
	"go.uber.org/zap/zapcore"
)

const (
	TgPrivateChatType = "private"
)

type ConversationState int

var (
	InitialState                   ConversationState = 0
	WaitingForName                 ConversationState = 1
	WaitingForDistrict             ConversationState = 2
	WaitingForIsSevChelParticipant ConversationState = 3
	WaitingForPurpose              ConversationState = 4
)

type ApplicationService interface {
	GetApplicationByTgUserID(ctx context.Context, tgUserID int64) (application *domain.Application, err error)
	CreateApplication(ctx context.Context, application *domain.Application) (newApplication *domain.Application, err error)
	ApproveApplication(ctx context.Context, applicationID uint) (application *domain.Application, err error)
	RejectApplication(ctx context.Context, applicationID uint) (application *domain.Application, err error)
}

type BotAPI interface {
	SendMessage(text string, chatID int64, opts *echotron.MessageOptions) (res echotron.APIResponseMessage, err error)
	DeleteMessage(chatID int64, messageID int) (res echotron.APIResponseBase, err error)
}

type Logger interface {
	Error(msg string, fields ...zapcore.Field)
}

type ApplicationBot struct {
	chatID       int64
	adminChatID  int64
	service      ApplicationService
	application  *domain.Application
	currentState ConversationState
	logger       Logger
	cfg          *config.Config
	BotAPI
}

func GetApplicationBotConstructor(adminChatID int64, api BotAPI, service ApplicationService, logger Logger, cfg *config.Config) echotron.NewBotFn {
	return func(chatID int64) echotron.Bot {
		return &ApplicationBot{
			chatID,
			adminChatID,
			service,
			&domain.Application{},
			InitialState,
			logger,
			cfg,
			api,
		}
	}
}

// Update handles the event that happens in tg chat between bot and possible user
// If error happens in any subprogramm that Update calls, Update should log it and maybe process in any other way
// No logging should be present in any of Update subprogramms
func (b *ApplicationBot) Update(update *echotron.Update) {
	if update == nil {
		b.logger.Error(domain.ErrInvalidUpdate.Error())
		return
	}
	var err error
	messageText := ""
	messageID := 0
	// If this condition is true, then update is not relevant for the purpose of processing an application
	if update.Message == nil && update.CallbackQuery == nil {
		b.logger.Error(domain.ErrInvalidUpdate.Error())
		return
	} else if update.Message != nil {
		messageText, messageID, err = b.processMessageUpdate(update)
		if err != nil {
			b.logger.Error(err.Error())
			return
		}
	} else if update.CallbackQuery != nil {
		messageText, messageID, err = b.processCallbackQueryUpdate(update)
		if err != nil {
			b.logger.Error(err.Error())
			return
		}
	} // else is unreachable: a&&b || !a || !b is equivalent to true (a := update.Message == nil, b := update.CallbackQuery == nil)

	err = b.handleTextMessage(messageText, messageID)
	if err != nil {
		b.logger.Error(err.Error())
		return
	}
}

// ValidateIncomingChat checks if current chat is private, and if it is not, returns domain.ErrChatNotPrivate
func (b *ApplicationBot) validateIncomingChat(c echotron.Chat) error {
	if c.Type != TgPrivateChatType {
		return domain.ErrChatNotPrivate
	}

	return nil
}

// SetSenderInfo validates *echotron.User, with whom the chat is held,
// and sets their tg ID and username as b.application.TelegramUserID and b.application.TelegramUsername.
func (b *ApplicationBot) setSenderInfo(s *echotron.User) (err error) {
	// if s == nil, then we cannot process application as the user is unknown
	// domain.ErrMessageSenderNotSpecified is returned if we can send corresponding message to user
	// in other case, internal error of sending message is returned
	if s == nil {
		b.currentState = InitialState
		_, err = b.SendMessage(OnlyRealUsersText, b.chatID, OptionsEmpty)
		if err != nil {
			return
		}

		err = domain.ErrMessageSenderNotSpecified
		return
	}

	b.application.TelegramUserID = s.ID

	// If sender username is not present, we should notify them that tg username is mandatory
	if s.Username == "" {
		b.currentState = InitialState
		_, err = b.SendMessage(OnlyUsersWithUsernameText, b.chatID, OptionsEmpty)
		if err != nil {
			return
		}

		err = domain.ErrEmptyUsername
		return
	}

	b.application.TelegramUsername = s.Username
	return nil
}

func (b *ApplicationBot) parseUpdateText(t string) (messageText string, err error) {
	messageText = strings.TrimSpace(t)
	// there is no point in processing message that is empty or contains only whitespaces
	if messageText == "" {
		_, err = b.SendMessage(EnterUnemptyMessageText, b.chatID, OptionsEmpty)
		if err != nil {
			return
		}

		err = domain.ErrEmptyMessageText
		return
	}

	return
}

// ProcessMessageUpdate processes echotron.Update object.
// It returns plain text of sent message, messageID (for purposes of interacting with the message in future if needed) and err
func (b *ApplicationBot) processMessageUpdate(update *echotron.Update) (messageText string, messageID int, err error) {
	if update == nil {
		err = domain.ErrInvalidUpdate
		return
	}

	if update.Message == nil {
		err = domain.ErrNoMessage
		return
	}

	err = b.validateIncomingChat(update.Message.Chat)
	if err != nil {
		return
	}

	err = b.setSenderInfo(update.Message.From)
	if err != nil {
		return
	}

	messageText, err = b.parseUpdateText(update.Message.Text)
	if err != nil {
		return
	}

	messageID = update.Message.ID

	return
}

// ProcessCallbackQueryUpdate processes echotron.Update object.
// It returns plain text of sent callback query, messageID (for purposes of interacting with the message in future if needed) and err.
func (b *ApplicationBot) processCallbackQueryUpdate(update *echotron.Update) (messageText string, messageID int, err error) {
	if update == nil {
		err = domain.ErrInvalidUpdate
		return
	}
	if update.CallbackQuery == nil {
		err = domain.ErrInvalidCallbackQuery
		return
	}
	// Callback query should always be initiated by keyboard click on the bot-sent private message.
	if update.CallbackQuery.Message == nil {
		err = domain.ErrNoCallbackQueryMessage
		return
	}

	err = b.validateIncomingChat(update.CallbackQuery.Message.Chat)
	if err != nil {
		return
	}

	err = b.setSenderInfo(update.CallbackQuery.From)
	if err != nil {
		return
	}

	messageText, err = b.parseUpdateText(update.CallbackQuery.Data)
	if err != nil {
		return
	}

	messageID = update.CallbackQuery.Message.ID

	return
}

// HandleTextMessage handles parsed plain text and message id from either ValidateCallbackQuery or ValidateTextMessage call.
// It recognizes command patterns and, if text is a command, calls corresponding command handler.
// If text is not a command, then HandleTextMessage passes it on to the state machine that implements conversation flow (HandlePlainTextMessage method).
// If any error happens while processing the message or command, state machine is returned to InitialState.
func (b *ApplicationBot) handleTextMessage(messageText string, messageID int) (err error) {
	switch {
	case isApproveCommand(messageText):
		err = b.handleApproveCommand(messageText, messageID)
		if err != nil {
			b.currentState = InitialState
			return
		}
	case isRejectCommand(messageText):
		err = b.handleRejectCommand(messageText, messageID)
		if err != nil {
			b.currentState = InitialState
			return
		}
	case isStartCommand(messageText):
		err = b.handleStartCommand()
		if err != nil {
			b.currentState = InitialState
			return
		}
	case isCancelCommand(messageText):
		err = b.handleCancelCommand()
		if err != nil {
			b.currentState = InitialState
			return
		}
	default:
		err = b.HandlePlainTextMessage(messageText)
		if err != nil {
			b.currentState = InitialState
			return
		}
	}

	return
}

// handleStartCommand handles the "/start" bot command.
func (b *ApplicationBot) handleStartCommand() (err error) {
	// checks that this user did not apply twice or more
	_, err = b.service.GetApplicationByTgUserID(context.Background(), b.application.TelegramUserID)
	// if err != domain.ErrApplicationNotFound, then application was found or internal error happened
	if err != domain.ErrApplicationNotFound {
		// internal error happened
		if err != nil {
			_, err = b.SendMessage(InternalErrorText, b.chatID, OptionsWithStartKeyboard)
			if err != nil {
				return
			}

			return
		}

		// notify user that application with his tgUserID already exists
		_, err = b.SendMessage(ApplicationAlreadyExistsText, b.chatID, OptionsEmpty)
		if err != nil {
			return
		}

		err = domain.ErrApplicationAlreadyExists
		return
	}

	// user can apply, send introduction
	_, err = b.SendMessage(IntroductionText, b.chatID, OptionsEmpty)
	if err != nil {
		return
	}

	_, err = b.SendMessage(FullNamePromptText, b.chatID, OptionsWithCancelKeyboard)
	if err != nil {
		return
	}

	// switch state machine into WaitingForName state
	b.currentState = WaitingForName
	return
}

// handleCancelCommand handles "/cancel" command.
// From any state it switches the state machine into InitialState, notifying user that their application is cancelled.
func (b *ApplicationBot) handleCancelCommand() (err error) {
	_, err = b.SendMessage(CancelText, b.chatID, OptionsWithStartKeyboard)
	if err != nil {
		return
	}

	b.currentState = InitialState
	return
}

// handleCancelCommand handles plain text message that is the part of the conversation flow while applying.
// It is a state machine implementation.
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

			err = domain.ErrInvalidFullName
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
		b.application.IsSevChelParticipant = isYesCommand(messageText)

		var newApplication *domain.Application
		newApplication, err = b.service.CreateApplication(context.Background(), b.application)
		if err != nil {
			if err == domain.ErrApplicationAlreadyExists {
				_, err = b.SendMessage(ApplicationAlreadyExistsText, b.chatID, OptionsEmpty)
				if err != nil {
					return
				}

				err = domain.ErrApplicationAlreadyExists
				return
			}

			_, msgErr := b.SendMessage(InternalErrorText, b.chatID, OptionsWithStartKeyboard)
			if msgErr != nil {
				err = msgErr
				return
			}

			return
		}

		_, err = b.SendMessage(getConclusionText(b.cfg.Telegram.IsSecondRoster), b.chatID, OptionsEmpty)
		if err != nil {
			return
		}
		b.currentState = InitialState

		err = b.sendApplicationToAdmin(newApplication)
		if err != nil {
			return
		}
	default:
		b.currentState = InitialState
		_, err = b.SendMessage(PressStartText, b.chatID, OptionsWithStartKeyboard)
		if err != nil {
			return
		}
	}

	return
}

// sendApplicationToAdmin sends message text formatted by getApplicationMessageText to admin account.
func (b *ApplicationBot) sendApplicationToAdmin(application *domain.Application) (err error) {
	_, err = b.SendMessage(
		getApplicationMessageText(application),
		b.adminChatID,
		getOptionsWithApproveOrRejectKeyboard(application.ID),
	)
	if err != nil {
		return
	}

	return
}

// isSenderAdmin reports whether current chatID matches adminChatID.
func (b *ApplicationBot) isSenderAdmin() bool {
	return b.chatID == b.adminChatID
}

// handleApproveCommand approves application if current chat is chat between admin and bot.
// Then it sends approve text to the user who filed the application,
// sends confirm message to admin (to indicate that approve was successful),
// and deletes the message that caused handleApproveCommand call.
func (b *ApplicationBot) handleApproveCommand(text string, messageID int) (err error) {
	if !b.isSenderAdmin() {
		err = domain.ErrUnauthorized
		return
	}

	id, err := getApplicationIDFromApproveCommand(text)
	if err != nil {
		return
	}

	application, err := b.service.ApproveApplication(context.Background(), uint(id))
	if err != nil {
		return
	}

	_, err = b.SendMessage(ApproveText, application.TelegramUserID, OptionsEmpty)
	if err != nil {
		return
	}

	_, err = b.SendMessage(
		getApproveAdminMessageText(application.ID, application.TelegramUsername),
		b.adminChatID,
		OptionsEmpty,
	)
	if err != nil {
		return
	}

	_, err = b.DeleteMessage(b.adminChatID, messageID)
	if err != nil {
		return
	}

	return
}

// handleRejectCommand rejects application if current chat is chat between admin and bot.
// Then it sends reject text to the user who filed the application,
// sends confirm message to admin (to indicate that reject was successful),
// and deletes the message that caused handleRejectCommand call.
func (b *ApplicationBot) handleRejectCommand(text string, messageID int) (err error) {
	if !b.isSenderAdmin() {
		err = domain.ErrUnauthorized
		return
	}

	id, err := getApplicationIDFromRejectCommand(text)
	if err != nil {
		return
	}

	application, err := b.service.RejectApplication(context.Background(), uint(id))
	if err != nil {
		return
	}

	_, err = b.SendMessage(RejectText, application.TelegramUserID, OptionsEmpty)
	if err != nil {
		return
	}

	_, err = b.SendMessage(
		getRejectAdminMessageText(application.ID, application.TelegramUsername),
		b.adminChatID,
		OptionsEmpty,
	)
	if err != nil {
		return
	}

	_, err = b.DeleteMessage(b.adminChatID, messageID)
	if err != nil {
		return
	}

	return
}
