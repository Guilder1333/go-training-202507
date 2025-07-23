package presentation

import (
	"hands_on_go/internal/statuserr"
	"net/http"

	"github.com/rs/zerolog/log"
)

type ErrorResponseHandlerFunc func(w http.ResponseWriter, r *http.Request) error

func WithErrorResponse(next ErrorResponseHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err == nil {
			return
		}

		kind := statuserr.GetKind(err)
		var msg string
		var status int
		switch kind {
		case statuserr.KindInvalidRequest:
			status = http.StatusBadRequest
			msg = "invalid request"
		case statuserr.KindUserNotFound:
			status = http.StatusNotFound
			msg = "user not found"
		case statuserr.KindCreateUserFailed:
			status = http.StatusInternalServerError
			msg = "failed to create user"
		default:
			status = http.StatusInternalServerError
			msg = "something went wrong"
		}

		errmsg, ok := statuserr.GetMessage(err)
		if ok {
			msg = errmsg
		}

		log.Warn().Err(err).Msg("Error happended during request: " + msg)
		w.WriteHeader(status)
		w.Write([]byte(msg))
	}
}
