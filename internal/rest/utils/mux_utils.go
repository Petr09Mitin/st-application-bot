package utils

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/st-application-bot/domain"
)

const (
	ApplicationIDMuxVarName = "applicationID"
)

func GetApplicationIDMuxVar(r *http.Request) (applicationID uint, err error) {
	applicationIDData, ok := mux.Vars(r)[ApplicationIDMuxVarName]
	if !ok {
		err = domain.ErrInvalidData
		return
	}

	applicationIDInt, err := strconv.Atoi(applicationIDData)
	if err != nil {
		err = domain.ErrInvalidData
		return
	}

	applicationID = uint(applicationIDInt)

	return
}
