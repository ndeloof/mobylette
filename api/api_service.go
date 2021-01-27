/*
 * Docker Engine API
 *
 * The Engine API is an HTTP API served by Docker Engine. It is the API the Docker client uses to communicate with the Engine, so everything the Docker client can do can be done with the API.  Most of the client's commands map directly to API endpoints (e.g. `docker ps` is `GET /containers/json`). The notable exception is running containers, which consists of several API calls.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```  # Versioning  The API is usually changed in each release, so API calls are versioned to ensure that clients don't break. To lock to a specific version of the API, you prefix the URL with its version, for example, call `/v1.30/info` to use the v1.30 version of the `/info` endpoint. If the API version specified in the URL is not supported by the daemon, a HTTP `400 Bad Request` error message is returned.  If you omit the version-prefix, the current version of the API (v1.41) is used. For example, calling `/info` is the same as calling `/v1.41/info`. Using the API without a version-prefix is deprecated and will be removed in a future release.  Engine releases in the near future should support this version of the API, so your client will continue to work even if it is talking to a newer Engine.  The API uses an open schema model, which means server may add extra properties to responses. Likewise, the server will ignore any extra query parameters and request body properties. When you write clients, you need to ignore additional properties in responses to ensure they do not break when talking to newer daemons.   # Authentication  Authentication for registries is handled client side. The client has to send authentication details to various endpoints that need to communicate with registries, such as `POST /images/(name)/push`. These are sent as `X-Registry-Auth` header as a [base64url encoded](https://tools.ietf.org/html/rfc4648#section-5) (JSON) string with the following structure:  ``` {   \"username\": \"string\",   \"password\": \"string\",   \"email\": \"string\",   \"serveraddress\": \"string\" } ```  The `serveraddress` is a domain/IP without a protocol. Throughout this structure, double quotes are required.  If you have already got an identity token from the [`/auth` endpoint](#operation/SystemAuth), you can just pass this instead of credentials:  ``` {   \"identitytoken\": \"9cbaf023786cd7...\" } ``` 
 *
 * API version: 1.41
 * Generated by: api Generator (https://api-generator.tech)
 */

package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A ServiceApiController binds http requests to an api service and writes the service results to the http response
type ServiceApiController struct {
	service ServiceApiServicer
}

// NewServiceApiController creates a default api controller
func NewServiceApiController(s ServiceApiServicer) Router {
	return &ServiceApiController{ service: s }
}

// Routes returns all of the api route for the ServiceApiController
func (c *ServiceApiController) Routes() Routes {
	return Routes{ 
		{
			"ServiceCreate",
			strings.ToUpper("Post"),
			"/v1.41/services/create",
			c.ServiceCreate,
		},
		{
			"ServiceDelete",
			strings.ToUpper("Delete"),
			"/v1.41/services/{id}",
			c.ServiceDelete,
		},
		{
			"ServiceInspect",
			strings.ToUpper("Get"),
			"/v1.41/services/{id}",
			c.ServiceInspect,
		},
		{
			"ServiceList",
			strings.ToUpper("Get"),
			"/v1.41/services",
			c.ServiceList,
		},
		{
			"ServiceLogs",
			strings.ToUpper("Get"),
			"/v1.41/services/{id}/logs",
			c.ServiceLogs,
		},
		{
			"ServiceUpdate",
			strings.ToUpper("Post"),
			"/v1.41/services/{id}/update",
			c.ServiceUpdate,
		},
	}
}

// ServiceCreate - Create a service
func (c *ServiceApiController) ServiceCreate(w http.ResponseWriter, r *http.Request) { 
	body := &UNKNOWN_BASE_TYPE{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	xRegistryAuth := r.Header.Get("X-Registry-Auth")
	result, err := c.service.ServiceCreate(r.Context(), *body, xRegistryAuth)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ServiceDelete - Delete a service
func (c *ServiceApiController) ServiceDelete(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	id := params["id"]
	result, err := c.service.ServiceDelete(r.Context(), id)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ServiceInspect - Inspect a service
func (c *ServiceApiController) ServiceInspect(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	query := r.URL.Query()
	id := params["id"]
	insertDefaults := query.Get("insertDefaults")
	result, err := c.service.ServiceInspect(r.Context(), id, insertDefaults)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ServiceList - List services
func (c *ServiceApiController) ServiceList(w http.ResponseWriter, r *http.Request) { 
	query := r.URL.Query()
	filters := query.Get("filters")
	status := query.Get("status")
	result, err := c.service.ServiceList(r.Context(), filters, status)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ServiceLogs - Get service logs
func (c *ServiceApiController) ServiceLogs(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	query := r.URL.Query()
	id := params["id"]
	details := query.Get("details")
	follow := query.Get("follow")
	stdout := query.Get("stdout")
	stderr := query.Get("stderr")
	since, err := parseInt32Parameter(query.Get("since"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	timestamps := query.Get("timestamps")
	tail := query.Get("tail")
	result, err := c.service.ServiceLogs(r.Context(), id, details, follow, stdout, stderr, since, timestamps, tail)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// ServiceUpdate - Update a service
func (c *ServiceApiController) ServiceUpdate(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	query := r.URL.Query()
	id := params["id"]
	version, err := parseInt32Parameter(query.Get("version"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	body := &UNKNOWN_BASE_TYPE{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	registryAuthFrom := query.Get("registryAuthFrom")
	rollback := query.Get("rollback")
	xRegistryAuth := r.Header.Get("X-Registry-Auth")
	result, err := c.service.ServiceUpdate(r.Context(), id, version, *body, registryAuthFrom, rollback, xRegistryAuth)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}
