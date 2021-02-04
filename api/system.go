package api


import (
	"fmt"
	"strconv"
	"net/http"
)


// SystemPing - Ping
func (s *Server) Ping(w http.ResponseWriter, r *http.Request) { 
	setHeaders(w)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	ok, err := s.ctr.IsServing(r.Context())
	if ok {
		w.Header().Set("Content-Length", "2")
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusTeapot)
		r := []byte(fmt.Sprintf("Sorry, can't connect to containerd: %s", err.Error()))
		w.Header().Set("Content-Length", strconv.Itoa(len(r)))
		w.Write(r)
	}
}

