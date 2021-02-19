package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *server) registerPingHandlers(r *mux.Router) {
	r.Handle("/_ping", s.APIHandler(s.ping)).Methods(http.MethodGet, http.MethodHead)
}

func (s *server) ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	ok, err := s.ctr.IsServing(r.Context())
	if !ok || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Length", "2")
		w.Write([]byte("OK"))
	}
}
