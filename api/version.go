package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

const Versioned = "/v{version:[0-9][0-9.]*}"

func (s *server) registerVersionHandlers(r *mux.Router) {
	r.Handle("/version", s.JSONAPIHandler(s.version)).Methods(http.MethodGet)
	r.Handle(Versioned+"/version", s.JSONAPIHandler(s.version)).Methods(http.MethodGet)
}

func (s *server) version(r *http.Request) (interface{}, *int, error) {
	version, err := s.ctr.Version(r.Context())
	if err != nil {
		return nil, nil, err
	}

	return SystemVersion{
		Version:       "0.0.1",
		ApiVersion:    "1.41",
		MinAPIVersion: "1.12",
		Experimental:  true,
		Components: []SystemVersionComponents{
			{
				Name:    "containerd",
				Version: version.Version,
			},
		},
	}, nil, nil
}

// SystemVersion - Response of Engine API: GET \"/version\"
type SystemVersion struct {
	Platform SystemVersionPlatform `json:"Platform,omitempty"`

	// Information about system components
	Components []SystemVersionComponents `json:"Components,omitempty"`

	// The version of the daemon
	Version string `json:"Version,omitempty"`

	// The default (and highest) API version that is supported by the daemon
	ApiVersion string `json:"ApiVersion,omitempty"`

	// The minimum API version that is supported by the daemon
	MinAPIVersion string `json:"MinAPIVersion,omitempty"`

	// The Git commit of the source code that was used to build the daemon
	GitCommit string `json:"GitCommit,omitempty"`

	// The version Go used to compile the daemon, and the version of the Go runtime in use.
	GoVersion string `json:"GoVersion,omitempty"`

	// The operating system that the daemon is running on (\"linux\" or \"windows\")
	Os string `json:"Os,omitempty"`

	// The architecture that the daemon is running on
	Arch string `json:"Arch,omitempty"`

	// The kernel version (`uname -r`) that the daemon is running on.  This field is omitted when empty.
	KernelVersion string `json:"KernelVersion,omitempty"`

	// Indicates if the daemon is started with experimental features enabled.  This field is omitted when empty / false.
	Experimental bool `json:"Experimental,omitempty"`

	// The date and time that the daemon was compiled.
	BuildTime string `json:"BuildTime,omitempty"`
}

type SystemVersionPlatform struct {
	Name string `json:"Name"`
}

type SystemVersionComponents struct {

	// Name of the component
	Name string `json:"Name"`

	// Version of the component
	Version string `json:"Version"`

	// Key/value pairs of strings with additional information about the component. These values are intended for informational purposes only, and their content is not defined, and not part of the API specification.  These messages can be printed by the client as information to the user.
	Details *map[string]interface{} `json:"Details,omitempty"`
}
