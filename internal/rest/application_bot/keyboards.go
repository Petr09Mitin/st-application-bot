package applicationbot

import "github.com/NicoNex/echotron/v3"

var (
	StartKeyboard = echotron.InlineKeyboardMarkup{
		InlineKeyboard: [][]echotron.InlineKeyboardButton{
			{
				{
					Text:         "Начать",
					CallbackData: StartCommand,
				},
			},
		},
	}
	CancelKeyboard = echotron.InlineKeyboardMarkup{
		InlineKeyboard: [][]echotron.InlineKeyboardButton{
			{
				{
					Text:         "Отмена",
					CallbackData: CancelCommand,
				},
			},
		},
	}
	YesOrNoWithCancelKeyboard = echotron.InlineKeyboardMarkup{
		InlineKeyboard: [][]echotron.InlineKeyboardButton{
			{
				{
					Text:         "Да",
					CallbackData: YesCommand,
				},
				{
					Text:         "Нет",
					CallbackData: NoCommand,
				},
			},
			{
				{
					Text:         "Отмена",
					CallbackData: CancelCommand,
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
