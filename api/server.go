package api

import (
	"encoding/json"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/defaults"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"net"
	"net/http"
	"os"
)

const (
	defaultNamespace = "com.docker.mobylette"
	defaultRuntime   = "io.containerd.runc.v2"
)

func NewServer() (*server, error) {
	client, err := containerd.New(
		defaults.DefaultAddress,
		containerd.WithDefaultNamespace(defaultNamespace),
		containerd.WithDefaultRuntime(defaultRuntime),
	)
	if err != nil {
		return nil, err
	}
	l, err := net.Listen("unix", "/var/run/moby.sock")
	if err != nil {
		return nil, err
	}

	router := mux.NewRouter().UseEncodedPath()
	server := &server{
		Server: http.Server{
			Addr:    "/var/run/moby.sock",
			Handler: handlers.LoggingHandler(os.Stdout, router),
		},
		listener: l,
		Decoder:  NewAPIDecoder(),
		ctr:      client,
	}

	server.registerPingHandlers(router)
	server.registerVersionHandlers(router)

	return server, nil
}

type server struct {
	http.Server
	listener net.Listener
	*schema.Decoder

	ctr *containerd.Client
}

func (s *server) Serve() error {
	defer s.listener.Close()
	return s.Server.Serve(s.listener)
}

// APIHandler wrap a minimalist HandlerFunc with all common API stuff
func (s *server) APIHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Api-Version", "1.41")
		w.Header().Set("server", "Mobylette/0.0.1 (linux)")
		w.Header().Add("Ostype:", "linux")
		w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Add("Pragma", "no-cache")
	}
}

type JSONHandlerFunc func(r *http.Request) (interface{}, *int, error)

// JSONAPIHandler wrap a minimalist HandlerFunc with all common API stuff returning json data
func (s *server) JSONAPIHandler(h JSONHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status, err := h(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if status != nil {
			w.WriteHeader(*status)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		json.NewEncoder(w).Encode(data)
	}
}
