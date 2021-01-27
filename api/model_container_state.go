/*
 * Docker Engine API
 *
 * The Engine API is an HTTP API served by Docker Engine. It is the API the Docker client uses to communicate with the Engine, so everything the Docker client can do can be done with the API.  Most of the client's commands map directly to API endpoints (e.g. `docker ps` is `GET /containers/json`). The notable exception is running containers, which consists of several API calls.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```  # Versioning  The API is usually changed in each release, so API calls are versioned to ensure that clients don't break. To lock to a specific version of the API, you prefix the URL with its version, for example, call `/v1.30/info` to use the v1.30 version of the `/info` endpoint. If the API version specified in the URL is not supported by the daemon, a HTTP `400 Bad Request` error message is returned.  If you omit the version-prefix, the current version of the API (v1.41) is used. For example, calling `/info` is the same as calling `/v1.41/info`. Using the API without a version-prefix is deprecated and will be removed in a future release.  Engine releases in the near future should support this version of the API, so your client will continue to work even if it is talking to a newer Engine.  The API uses an open schema model, which means server may add extra properties to responses. Likewise, the server will ignore any extra query parameters and request body properties. When you write clients, you need to ignore additional properties in responses to ensure they do not break when talking to newer daemons.   # Authentication  Authentication for registries is handled client side. The client has to send authentication details to various endpoints that need to communicate with registries, such as `POST /images/(name)/push`. These are sent as `X-Registry-Auth` header as a [base64url encoded](https://tools.ietf.org/html/rfc4648#section-5) (JSON) string with the following structure:  ``` {   \"username\": \"string\",   \"password\": \"string\",   \"email\": \"string\",   \"serveraddress\": \"string\" } ```  The `serveraddress` is a domain/IP without a protocol. Throughout this structure, double quotes are required.  If you have already got an identity token from the [`/auth` endpoint](#operation/SystemAuth), you can just pass this instead of credentials:  ``` {   \"identitytoken\": \"9cbaf023786cd7...\" } ``` 
 *
 * API version: 1.41
 * Generated by: api Generator (https://api-generator.tech)
 */

package api

// ContainerState - ContainerState stores container's running state. It's part of ContainerJSONBase and will be returned by the \"inspect\" command. 
type ContainerState struct {

	// String representation of the container state. Can be one of \"created\", \"running\", \"paused\", \"restarting\", \"removing\", \"exited\", or \"dead\". 
	Status string `json:"Status,omitempty"`

	// Whether this container is running.  Note that a running container can be _paused_. The `Running` and `Paused` booleans are not mutually exclusive:  When pausing a container (on Linux), the freezer cgroup is used to suspend all processes in the container. Freezing the process requires the process to be running. As a result, paused containers are both `Running` _and_ `Paused`.  Use the `Status` field instead to determine if a container's state is \"running\". 
	Running bool `json:"Running,omitempty"`

	// Whether this container is paused.
	Paused bool `json:"Paused,omitempty"`

	// Whether this container is restarting.
	Restarting bool `json:"Restarting,omitempty"`

	// Whether this container has been killed because it ran out of memory. 
	OOMKilled bool `json:"OOMKilled,omitempty"`

	Dead bool `json:"Dead,omitempty"`

	// The process ID of this container
	Pid int32 `json:"Pid,omitempty"`

	// The last exit code of this container
	ExitCode int32 `json:"ExitCode,omitempty"`

	Error string `json:"Error,omitempty"`

	// The time when this container was last started.
	StartedAt string `json:"StartedAt,omitempty"`

	// The time when this container last exited.
	FinishedAt string `json:"FinishedAt,omitempty"`

	Health Health `json:"Health,omitempty"`
}
