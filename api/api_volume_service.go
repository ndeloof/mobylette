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
)

// VolumeApiService is a service that implents the logic for the VolumeApiServicer
// This service should implement the business logic for every endpoint for the VolumeApi API. 
// Include any external packages or services that will be required by this service.
type VolumeApiService struct {
}

// NewVolumeApiService creates a default api service
func NewVolumeApiService() VolumeApiServicer {
	return &VolumeApiService{}
}

// VolumeCreate - Create a volume
func (s *VolumeApiService) VolumeCreate(ctx context.Context, volumeConfig VolumeConfig) (ImplResponse, error) {
	// TODO - update VolumeCreate with the required logic for this service method.
	// Add api_volume_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, Volume{}) or use other options such as http.Ok ...
	//return Response(201, Volume{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("VolumeCreate method not implemented")
}

// VolumeDelete - Remove a volume
func (s *VolumeApiService) VolumeDelete(ctx context.Context, name string, force bool) (ImplResponse, error) {
	// TODO - update VolumeDelete with the required logic for this service method.
	// Add api_volume_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(409, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(409, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("VolumeDelete method not implemented")
}

// VolumeInspect - Inspect a volume
func (s *VolumeApiService) VolumeInspect(ctx context.Context, name string) (ImplResponse, error) {
	// TODO - update VolumeInspect with the required logic for this service method.
	// Add api_volume_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Volume{}) or use other options such as http.Ok ...
	//return Response(200, Volume{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("VolumeInspect method not implemented")
}

// VolumeList - List volumes
func (s *VolumeApiService) VolumeList(ctx context.Context, filters string) (ImplResponse, error) {
	// TODO - update VolumeList with the required logic for this service method.
	// Add api_volume_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, VolumeListResponse{}) or use other options such as http.Ok ...
	//return Response(200, VolumeListResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("VolumeList method not implemented")
}

// VolumePrune - Delete unused volumes
func (s *VolumeApiService) VolumePrune(ctx context.Context, filters string) (ImplResponse, error) {
	// TODO - update VolumePrune with the required logic for this service method.
	// Add api_volume_service.api to the .api-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, VolumePruneResponse{}) or use other options such as http.Ok ...
	//return Response(200, VolumePruneResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("VolumePrune method not implemented")
}

