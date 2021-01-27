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
	"context"
	"net/http"
	"errors"
	"os"
)

// ContainerApiService is a service that implents the logic for the ContainerApiServicer
// This service should implement the business logic for every endpoint for the ContainerApi API. 
// Include any external packages or services that will be required by this service.
type ContainerApiService struct {
}

// NewContainerApiService creates a default api service
func NewContainerApiService() ContainerApiServicer {
	return &ContainerApiService{}
}

// ContainerArchive - Get an archive of a filesystem resource in a container
func (s *ContainerApiService) ContainerArchive(ctx context.Context, id string, path string) (ImplResponse, error) {
	// TODO - update ContainerArchive with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, InlineResponse400{}) or use other options such as http.Ok ...
	//return Response(400, InlineResponse400{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerArchive method not implemented")
}

// ContainerArchiveInfo - Get information about files in a container
func (s *ContainerApiService) ContainerArchiveInfo(ctx context.Context, id string, path string) (ImplResponse, error) {
	// TODO - update ContainerArchiveInfo with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, InlineResponse400{}) or use other options such as http.Ok ...
	//return Response(400, InlineResponse400{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerArchiveInfo method not implemented")
}

// ContainerAttach - Attach to a container
func (s *ContainerApiService) ContainerAttach(ctx context.Context, id string, detachKeys string, logs bool, stream bool, stdin bool, stdout bool, stderr bool) (ImplResponse, error) {
	// TODO - update ContainerAttach with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(101, {}) or use other options such as http.Ok ...
	//return Response(101, nil),nil

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerAttach method not implemented")
}

// ContainerAttachWebsocket - Attach to a container via a websocket
func (s *ContainerApiService) ContainerAttachWebsocket(ctx context.Context, id string, detachKeys string, logs bool, stream bool, stdin bool, stdout bool, stderr bool) (ImplResponse, error) {
	// TODO - update ContainerAttachWebsocket with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(101, {}) or use other options such as http.Ok ...
	//return Response(101, nil),nil

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerAttachWebsocket method not implemented")
}

// ContainerChanges - Get changes on a container’s filesystem
func (s *ContainerApiService) ContainerChanges(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update ContainerChanges with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []ContainerChangeResponseItem{}) or use other options such as http.Ok ...
	//return Response(200, []ContainerChangeResponseItem{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerChanges method not implemented")
}

// ContainerCreate - Create a container
func (s *ContainerApiService) ContainerCreate(ctx context.Context, body UNKNOWN_BASE_TYPE, name string) (ImplResponse, error) {
	// TODO - update ContainerCreate with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, ContainerCreateResponse{}) or use other options such as http.Ok ...
	//return Response(201, ContainerCreateResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(409, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(409, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerCreate method not implemented")
}

// ContainerDelete - Remove a container
func (s *ContainerApiService) ContainerDelete(ctx context.Context, id string, v bool, force bool, link bool) (ImplResponse, error) {
	// TODO - update ContainerDelete with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(409, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(409, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerDelete method not implemented")
}

// ContainerExport - Export a container
func (s *ContainerApiService) ContainerExport(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update ContainerExport with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerExport method not implemented")
}

// ContainerInspect - Inspect a container
func (s *ContainerApiService) ContainerInspect(ctx context.Context, id string, size bool) (ImplResponse, error) {
	// TODO - update ContainerInspect with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ContainerInspectResponse{}) or use other options such as http.Ok ...
	//return Response(200, ContainerInspectResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerInspect method not implemented")
}

// ContainerKill - Kill a container
func (s *ContainerApiService) ContainerKill(ctx context.Context, id string, signal string) (ImplResponse, error) {
	// TODO - update ContainerKill with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(409, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(409, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerKill method not implemented")
}

// ContainerList - List containers
func (s *ContainerApiService) ContainerList(ctx context.Context, all bool, limit int32, size bool, filters string) (ImplResponse, error) {
	// TODO - update ContainerList with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(200, []map[string]interface{}{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerList method not implemented")
}

// ContainerLogs - Get container logs
func (s *ContainerApiService) ContainerLogs(ctx context.Context, id string, follow bool, stdout bool, stderr bool, since int32, until int32, timestamps bool, tail string) (ImplResponse, error) {
	// TODO - update ContainerLogs with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, *os.File{}) or use other options such as http.Ok ...
	//return Response(200, *os.File{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerLogs method not implemented")
}

// ContainerPause - Pause a container
func (s *ContainerApiService) ContainerPause(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update ContainerPause with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerPause method not implemented")
}

// ContainerPrune - Delete stopped containers
func (s *ContainerApiService) ContainerPrune(ctx context.Context, filters string) (ImplResponse, error) {
	// TODO - update ContainerPrune with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ContainerPruneResponse{}) or use other options such as http.Ok ...
	//return Response(200, ContainerPruneResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerPrune method not implemented")
}

// ContainerRename - Rename a container
func (s *ContainerApiService) ContainerRename(ctx context.Context, id string, name string) (ImplResponse, error) {
	// TODO - update ContainerRename with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(409, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(409, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerRename method not implemented")
}

// ContainerResize - Resize a container TTY
func (s *ContainerApiService) ContainerResize(ctx context.Context, id string, h int32, w int32) (ImplResponse, error) {
	// TODO - update ContainerResize with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerResize method not implemented")
}

// ContainerRestart - Restart a container
func (s *ContainerApiService) ContainerRestart(ctx context.Context, id string, t int32) (ImplResponse, error) {
	// TODO - update ContainerRestart with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerRestart method not implemented")
}

// ContainerStart - Start a container
func (s *ContainerApiService) ContainerStart(ctx context.Context, id string, detachKeys string) (ImplResponse, error) {
	// TODO - update ContainerStart with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(304, {}) or use other options such as http.Ok ...
	//return Response(304, nil),nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerStart method not implemented")
}

// ContainerStats - Get container stats based on resource usage
func (s *ContainerApiService) ContainerStats(ctx context.Context, id string, stream bool, oneShot bool) (ImplResponse, error) {
	// TODO - update ContainerStats with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(200, map[string]interface{}{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerStats method not implemented")
}

// ContainerStop - Stop a container
func (s *ContainerApiService) ContainerStop(ctx context.Context, id string, t int32) (ImplResponse, error) {
	// TODO - update ContainerStop with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(304, {}) or use other options such as http.Ok ...
	//return Response(304, nil),nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerStop method not implemented")
}

// ContainerTop - List processes running inside a container
func (s *ContainerApiService) ContainerTop(ctx context.Context, id string, psArgs string) (ImplResponse, error) {
	// TODO - update ContainerTop with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ContainerTopResponse{}) or use other options such as http.Ok ...
	//return Response(200, ContainerTopResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerTop method not implemented")
}

// ContainerUnpause - Unpause a container
func (s *ContainerApiService) ContainerUnpause(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update ContainerUnpause with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerUnpause method not implemented")
}

// ContainerUpdate - Update a container
func (s *ContainerApiService) ContainerUpdate(ctx context.Context, id string, update UNKNOWN_BASE_TYPE) (ImplResponse, error) {
	// TODO - update ContainerUpdate with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ContainerUpdateResponse{}) or use other options such as http.Ok ...
	//return Response(200, ContainerUpdateResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerUpdate method not implemented")
}

// ContainerWait - Wait for a container
func (s *ContainerApiService) ContainerWait(ctx context.Context, id string, condition string) (ImplResponse, error) {
	// TODO - update ContainerWait with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ContainerWaitResponse{}) or use other options such as http.Ok ...
	//return Response(200, ContainerWaitResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ContainerWait method not implemented")
}

// PutContainerArchive - Extract an archive of files or folders to a directory in a container
func (s *ContainerApiService) PutContainerArchive(ctx context.Context, id string, path string, inputStream *os.File, noOverwriteDirNonDir string, copyUIDGID string) (ImplResponse, error) {
	// TODO - update PutContainerArchive with the required logic for this service method.
	// Add api_container_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(403, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(403, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PutContainerArchive method not implemented")
}

