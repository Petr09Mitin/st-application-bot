package utils

import "fmt"

func FormatChatID(chatID int64) string {
	return fmt.Sprintf("-%d", chatID)
}
