package applicationbot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/st-application-bot/domain"
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

func getApplicationIDFromApproveCommand(text string) (id uint64, err error) {
	data, ok := strings.CutPrefix(text, approveCommandPrefix)
	if !ok {
		err = domain.ErrInvalidData
		return
	}

	return strconv.ParseUint(data, 10, 64)
}

func getApplicationIDFromRejectCommand(text string) (id uint64, err error) {
	data, ok := strings.CutPrefix(text, rejectCommandPrefix)
	if !ok {
		err = domain.ErrInvalidData
		return
	}

	return strconv.ParseUint(data, 10, 64)
}
