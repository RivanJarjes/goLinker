package httprouter

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/status", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok\n"))
	})
	return r
}
