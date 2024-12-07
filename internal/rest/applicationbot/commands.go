package applicationbot

import (
	"fmt"
	"strings"
)

var (
	StartCommand         = "/start"
	CancelCommand        = "/cancel"
	YesCommand           = "/yes"
	NoCommand            = "/no"
	approveCommandPrefix = "/approve/"
	rejectCommandPrefix  = "/reject/"
)

func getApproveCommand(applicationID uint64) string {
	return fmt.Sprintf("%s%d", approveCommandPrefix, applicationID)
}

func getRejectCommand(applicationID uint64) string {
	return fmt.Sprintf("%s%d", rejectCommandPrefix, applicationID)
}

func isApproveCommand(text string) bool {
	return strings.HasPrefix(text, approveCommandPrefix)
}

func isRejectCommand(text string) bool {
	return strings.HasPrefix(text, rejectCommandPrefix)
}
