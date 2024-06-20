package json

import (
	"context"
	defJSON "encoding/json"
	"net/http"

	"github.com/st-application-bot/domain"
)

type JSONResponse struct {
	Body any `json:"body"`
}

func MarshalResponseBody(value any) (data []byte, err error) {
	data, err = defJSON.Marshal(&JSONResponse{Body: value})
	if err != nil {
		err = domain.ErrJSONMarshalling
		return
	}

	return
}

func MarshalResponseError(errMsg string) (data []byte, err error) {
	data, err = defJSON.Marshal(&domain.HTTPError{Error: errMsg})
	if err != nil {
		err = domain.ErrJSONMarshalling
		return
	}

	return
}

func ServeJSONBody(ctx context.Context, w http.ResponseWriter, value any, statusCode int) {
	data, err := MarshalResponseBody(value)
	if err != nil {
		ServeJSONError(ctx, w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json;")

	w.WriteHeader(statusCode)
	_, err = w.Write(data)
	if err != nil {
		ServeJSONError(ctx, w, err)
		return
	}
}

func ServeJSONError(ctx context.Context, w http.ResponseWriter, err error) {
	msg, status := domain.ParseHTTPError(err)

	w.Header().Set("Content-Type", "application/json;")

	data, err := MarshalResponseError(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	_, err = w.Write(data)
	if err != nil {
		return
	}
}
