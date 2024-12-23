package applicationbot

import "github.com/NicoNex/echotron/v3"

var (
	StartKeyboard = echotron.InlineKeyboardMarkup{
		InlineKeyboard: [][]echotron.InlineKeyboardButton{
			{
				{
					Text:         "Начать",
					CallbackData: startCommand,
				},
			},
		},
	}
	CancelKeyboard = echotron.InlineKeyboardMarkup{
		InlineKeyboard: [][]echotron.InlineKeyboardButton{
			{
				{
					Text:         "Отмена",
					CallbackData: cancelCommand,
				},
			},
		},
	}
	YesOrNoWithCancelKeyboard = echotron.InlineKeyboardMarkup{
		InlineKeyboard: [][]echotron.InlineKeyboardButton{
			{
				{
					Text:         "Да",
					CallbackData: yesCommand,
				},
				{
					Text:         "Нет",
					CallbackData: noCommand,
				},
			},
			{
				{
					Text:         "Отмена",
					CallbackData: cancelCommand,
				},
			},
		},
	}
	OptionsEmpty = &echotron.MessageOptions{
		ParseMode: echotron.HTML,
	}
	OptionsWithStartKeyboard = &echotron.MessageOptions{
		ParseMode:   echotron.HTML,
		ReplyMarkup: StartKeyboard,
	}
	OptionsWithCancelKeyboard = &echotron.MessageOptions{
		ParseMode:   echotron.HTML,
		ReplyMarkup: CancelKeyboard,
	}
	OptionsWithYesOrNoWithCancelKeyboard = &echotron.MessageOptions{
		ParseMode:   echotron.HTML,
		ReplyMarkup: YesOrNoWithCancelKeyboard,
	}
)

func getApproveOrRejectKeyboard(applicationID uint64) echotron.InlineKeyboardMarkup {
	return echotron.InlineKeyboardMarkup{
		InlineKeyboard: [][]echotron.InlineKeyboardButton{
			{
				{
					Text:         "Одобрить",
					CallbackData: getApproveCommand(applicationID),
				},
				{
					Text:         "Отклонить",
					CallbackData: getRejectCommand(applicationID),
				},
			},
		},
	}
}

func getOptionsWithApproveOrRejectKeyboard(applicationID uint64) *echotron.MessageOptions {
	return &echotron.MessageOptions{
		ParseMode:   echotron.HTML,
		ReplyMarkup: getApproveOrRejectKeyboard(applicationID),
	}
}
