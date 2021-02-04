package api

import (
	"net/http"
)

func setHeaders(w http.ResponseWriter) {
	w.Header().Add("Api-Version", "1.41")
	w.Header().Set("Server", "Mobylette/0.0.1 (linux)")
	w.Header().Add("Ostype:", "linux")
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Add("Pragma", "no-cache")
}