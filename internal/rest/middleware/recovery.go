package middleware

import (
	"net/http"

	"github.com/st-application-bot/domain"
	"github.com/st-application-bot/pkg/json"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				json.ServeJSONError(r.Context(), w, domain.ErrInternal)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
