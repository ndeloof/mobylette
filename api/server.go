package api

import (
	"strings"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/defaults"
)


const (
	defaultNamespace = "com.docker.mobylette"
	defaultRuntime   = "io.containerd.runc.v2"
)

func NewServer() (*Server, error) {
	client, err := containerd.New(
		defaults.DefaultAddress,
		containerd.WithDefaultNamespace(defaultNamespace),
		containerd.WithDefaultRuntime(defaultRuntime),
	)
	if err != nil {
		return nil, err
	}
	return &Server{client}, nil
}

type Server struct {
	ctr *containerd.Client
}

func (a *Server) Routes() Routes {
	return Routes{ 
		{
			"Ping",
			strings.ToUpper("Get"),
			"/_ping",
			a.Ping,
		},
	}
}