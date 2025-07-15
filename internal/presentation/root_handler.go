package presentation

import "net/http"

func NewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
		w.WriteHeader(http.StatusOK)
	})
}
