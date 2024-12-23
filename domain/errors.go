package domain

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidEnv                 = errors.New("invalid env variables")
	ErrApplicationAlreadyExists   = errors.New("application already exists")
	ErrApplicationNotFound        = errors.New("application not found")
	ErrJSONMarshalling            = errors.New("unable to unmarshal JSON")
	ErrInternal                   = errors.New("internal server error")
	ErrInvalidData                = errors.New("invalid request data provided")
	ErrUnableToApproveApplication = errors.New("unable to approve application")
	ErrUnableToDeclineApplication = errors.New("unable to decline application")
	ErrUnableToSendMessage        = errors.New("unable to send message through telegram private chat")
	ErrUnauthorized               = errors.New("unauthorized")
	ErrChatNotPrivate             = errors.New("chat with bot is not a private chat")
	ErrInvalidUpdate              = errors.New("invalid echotron.Update argument")
	ErrMessageSenderNotSpecified  = errors.New("message sender not specified in update")
	ErrEmptyUsername              = errors.New("empty tg username")
	ErrEmptyMessageText           = errors.New("empty message text")
	ErrInvalidMessageID           = errors.New("invalid message id")
	ErrInvalidFullName            = errors.New("invalid full name")
	ErrNoCallbackQueryMessage     = errors.New("no message in callback query")
	ErrNoMessage                  = errors.New("no message provided")
	ErrInvalidCallbackQuery       = errors.New("invalid callback query")
)

type HTTPError struct {
	Error string `json:"error"`
}

var HTTPErrors = map[error]int{
	ErrInvalidEnv:                 http.StatusInternalServerError,
	ErrApplicationAlreadyExists:   http.StatusBadRequest,
	ErrApplicationNotFound:        http.StatusNotFound,
	ErrJSONMarshalling:            http.StatusInternalServerError,
	ErrInternal:                   http.StatusInternalServerError,
	ErrInvalidData:                http.StatusBadRequest,
	ErrUnableToApproveApplication: http.StatusInternalServerError,
	ErrUnableToDeclineApplication: http.StatusInternalServerError,
	ErrUnableToSendMessage:        http.StatusInternalServerError,
	ErrUnauthorized:               http.StatusUnauthorized,
}

func ParseHTTPError(err error) (msg string, status int) {
	if err == nil {
		err = ErrInternal
	}

	if err.Error() == "" {
		err = ErrInternal
	}

	status, ok := HTTPErrors[err]
	if !ok {
		status = 500
		err = ErrInternal
	}

	msg = err.Error()

	return
}
