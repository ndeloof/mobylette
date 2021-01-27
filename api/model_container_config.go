/*
 * Docker Engine API
 *
 * The Engine API is an HTTP API served by Docker Engine. It is the API the Docker client uses to communicate with the Engine, so everything the Docker client can do can be done with the API.  Most of the client's commands map directly to API endpoints (e.g. `docker ps` is `GET /containers/json`). The notable exception is running containers, which consists of several API calls.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```  # Versioning  The API is usually changed in each release, so API calls are versioned to ensure that clients don't break. To lock to a specific version of the API, you prefix the URL with its version, for example, call `/v1.30/info` to use the v1.30 version of the `/info` endpoint. If the API version specified in the URL is not supported by the daemon, a HTTP `400 Bad Request` error message is returned.  If you omit the version-prefix, the current version of the API (v1.41) is used. For example, calling `/info` is the same as calling `/v1.41/info`. Using the API without a version-prefix is deprecated and will be removed in a future release.  Engine releases in the near future should support this version of the API, so your client will continue to work even if it is talking to a newer Engine.  The API uses an open schema model, which means server may add extra properties to responses. Likewise, the server will ignore any extra query parameters and request body properties. When you write clients, you need to ignore additional properties in responses to ensure they do not break when talking to newer daemons.   # Authentication  Authentication for registries is handled client side. The client has to send authentication details to various endpoints that need to communicate with registries, such as `POST /images/(name)/push`. These are sent as `X-Registry-Auth` header as a [base64url encoded](https://tools.ietf.org/html/rfc4648#section-5) (JSON) string with the following structure:  ``` {   \"username\": \"string\",   \"password\": \"string\",   \"email\": \"string\",   \"serveraddress\": \"string\" } ```  The `serveraddress` is a domain/IP without a protocol. Throughout this structure, double quotes are required.  If you have already got an identity token from the [`/auth` endpoint](#operation/SystemAuth), you can just pass this instead of credentials:  ``` {   \"identitytoken\": \"9cbaf023786cd7...\" } ``` 
 *
 * API version: 1.41
 * Generated by: api Generator (https://api-generator.tech)
 */

package api

// ContainerConfig - Configuration for a container that is portable between hosts
type ContainerConfig struct {

	// The hostname to use for the container, as a valid RFC 1123 hostname.
	Hostname string `json:"Hostname,omitempty"`

	// The domain name to use for the container.
	Domainname string `json:"Domainname,omitempty"`

	// The user that commands are run as inside the container.
	User string `json:"User,omitempty"`

	// Whether to attach to `stdin`.
	AttachStdin bool `json:"AttachStdin,omitempty"`

	// Whether to attach to `stdout`.
	AttachStdout bool `json:"AttachStdout,omitempty"`

	// Whether to attach to `stderr`.
	AttachStderr bool `json:"AttachStderr,omitempty"`

	// An object mapping ports to an empty object in the form:  `{\"<port>/<tcp|udp|sctp>\": {}}` 
	ExposedPorts map[string]map[string]interface{} `json:"ExposedPorts,omitempty"`

	// Attach standard streams to a TTY, including `stdin` if it is not closed. 
	Tty bool `json:"Tty,omitempty"`

	// Open `stdin`
	OpenStdin bool `json:"OpenStdin,omitempty"`

	// Close `stdin` after one attached client disconnects
	StdinOnce bool `json:"StdinOnce,omitempty"`

	// A list of environment variables to set inside the container in the form `[\"VAR=value\", ...]`. A variable without `=` is removed from the environment, rather than to have an empty value. 
	Env []string `json:"Env,omitempty"`

	// Command to run specified as a string or an array of strings. 
	Cmd []string `json:"Cmd,omitempty"`

	Healthcheck HealthConfig `json:"Healthcheck,omitempty"`

	// Command is already escaped (Windows only)
	ArgsEscaped bool `json:"ArgsEscaped,omitempty"`

	// The name of the image to use when creating the container/ 
	Image string `json:"Image,omitempty"`

	// An object mapping mount point paths inside the container to empty objects. 
	Volumes map[string]map[string]interface{} `json:"Volumes,omitempty"`

	// The working directory for commands to run in.
	WorkingDir string `json:"WorkingDir,omitempty"`

	// The entry point for the container as a string or an array of strings.  If the array consists of exactly one empty string (`[\"\"]`) then the entry point is reset to system default (i.e., the entry point used by docker when there is no `ENTRYPOINT` instruction in the `Dockerfile`). 
	Entrypoint []string `json:"Entrypoint,omitempty"`

	// Disable networking for the container.
	NetworkDisabled bool `json:"NetworkDisabled,omitempty"`

	// MAC address of the container.
	MacAddress string `json:"MacAddress,omitempty"`

	// `ONBUILD` metadata that were defined in the image's `Dockerfile`. 
	OnBuild []string `json:"OnBuild,omitempty"`

	// User-defined key/value metadata.
	Labels map[string]string `json:"Labels,omitempty"`

	// Signal to stop a container as a string or unsigned integer. 
	StopSignal string `json:"StopSignal,omitempty"`

	// Timeout to stop a container in seconds.
	StopTimeout int32 `json:"StopTimeout,omitempty"`

	// Shell for when `RUN`, `CMD`, and `ENTRYPOINT` uses a shell. 
	Shell []string `json:"Shell,omitempty"`
}
