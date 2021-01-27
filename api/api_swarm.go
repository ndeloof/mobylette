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

// A SwarmApiController binds http requests to an api service and writes the service results to the http response
type SwarmApiController struct {
	service SwarmApiServicer
}

// NewSwarmApiController creates a default api controller
func NewSwarmApiController(s SwarmApiServicer) Router {
	return &SwarmApiController{ service: s }
}

// Routes returns all of the api route for the SwarmApiController
func (c *SwarmApiController) Routes() Routes {
	return Routes{ 
		{
			"SwarmInit",
			strings.ToUpper("Post"),
			"/v1.41/swarm/init",
			c.SwarmInit,
		},
		{
			"SwarmInspect",
			strings.ToUpper("Get"),
			"/v1.41/swarm",
			c.SwarmInspect,
		},
		{
			"SwarmJoin",
			strings.ToUpper("Post"),
			"/v1.41/swarm/join",
			c.SwarmJoin,
		},
		{
			"SwarmLeave",
			strings.ToUpper("Post"),
			"/v1.41/swarm/leave",
			c.SwarmLeave,
		},
		{
			"SwarmUnlock",
			strings.ToUpper("Post"),
			"/v1.41/swarm/unlock",
			c.SwarmUnlock,
		},
		{
			"SwarmUnlockkey",
			strings.ToUpper("Get"),
			"/v1.41/swarm/unlockkey",
			c.SwarmUnlockkey,
		},
		{
			"SwarmUpdate",
			strings.ToUpper("Post"),
			"/v1.41/swarm/update",
			c.SwarmUpdate,
		},
	}
}

// SwarmInit - Initialize a new swarm
func (c *SwarmApiController) SwarmInit(w http.ResponseWriter, r *http.Request) { 
	body := &InlineObject5{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.SwarmInit(r.Context(), *body)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// SwarmInspect - Inspect swarm
func (c *SwarmApiController) SwarmInspect(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.SwarmInspect(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// SwarmJoin - Join an existing swarm
func (c *SwarmApiController) SwarmJoin(w http.ResponseWriter, r *http.Request) { 
	body := &InlineObject6{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.SwarmJoin(r.Context(), *body)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// SwarmLeave - Leave a swarm
func (c *SwarmApiController) SwarmLeave(w http.ResponseWriter, r *http.Request) { 
	query := r.URL.Query()
	force := query.Get("force")
	result, err := c.service.SwarmLeave(r.Context(), force)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// SwarmUnlock - Unlock a locked manager
func (c *SwarmApiController) SwarmUnlock(w http.ResponseWriter, r *http.Request) { 
	body := &InlineObject7{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.SwarmUnlock(r.Context(), *body)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// SwarmUnlockkey - Get the unlock key
func (c *SwarmApiController) SwarmUnlockkey(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.SwarmUnlockkey(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// SwarmUpdate - Update a swarm
func (c *SwarmApiController) SwarmUpdate(w http.ResponseWriter, r *http.Request) { 
	query := r.URL.Query()
	version, err := parseInt64Parameter(query.Get("version"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	body := &SwarmSpec{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	rotateWorkerToken := query.Get("rotateWorkerToken")
	rotateManagerToken := query.Get("rotateManagerToken")
	rotateManagerUnlockKey := query.Get("rotateManagerUnlockKey")
	result, err := c.service.SwarmUpdate(r.Context(), version, *body, rotateWorkerToken, rotateManagerToken, rotateManagerUnlockKey)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}
