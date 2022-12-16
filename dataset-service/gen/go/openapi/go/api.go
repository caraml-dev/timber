/*
 * caraml/timber/v1/dataset_service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)



// DatasetServiceApiRouter defines the required methods for binding the api requests to a responses for the DatasetServiceApi
// The DatasetServiceApiRouter implementation should parse necessary information from the http request,
// pass the data to a DatasetServiceApiServicer to perform the required actions, then write the service results to the http response.
type DatasetServiceApiRouter interface { 
	DatasetServiceCreateLogWriter(http.ResponseWriter, *http.Request)
	DatasetServiceCreateObservationService(http.ResponseWriter, *http.Request)
	DatasetServiceGetLogMetadata(http.ResponseWriter, *http.Request)
	DatasetServiceGetLogWriter(http.ResponseWriter, *http.Request)
	DatasetServiceGetObservationService(http.ResponseWriter, *http.Request)
	DatasetServiceListLogMetadata(http.ResponseWriter, *http.Request)
	DatasetServiceListLogWriters(http.ResponseWriter, *http.Request)
	DatasetServiceListObservationServices(http.ResponseWriter, *http.Request)
	DatasetServiceUpdateLogWriter(http.ResponseWriter, *http.Request)
	DatasetServiceUpdateObservationService(http.ResponseWriter, *http.Request)
}


// DatasetServiceApiServicer defines the api actions for the DatasetServiceApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DatasetServiceApiServicer interface { 
	DatasetServiceCreateLogWriter(context.Context, string, DatasetServiceCreateLogWriterRequest) (ImplResponse, error)
	DatasetServiceCreateObservationService(context.Context, string, DatasetServiceCreateObservationServiceRequest) (ImplResponse, error)
	DatasetServiceGetLogMetadata(context.Context, string, string) (ImplResponse, error)
	DatasetServiceGetLogWriter(context.Context, string, string) (ImplResponse, error)
	DatasetServiceGetObservationService(context.Context, string, string) (ImplResponse, error)
	DatasetServiceListLogMetadata(context.Context, string) (ImplResponse, error)
	DatasetServiceListLogWriters(context.Context, string) (ImplResponse, error)
	DatasetServiceListObservationServices(context.Context, string) (ImplResponse, error)
	DatasetServiceUpdateLogWriter(context.Context, string, string, DatasetServiceUpdateLogWriterRequest) (ImplResponse, error)
	DatasetServiceUpdateObservationService(context.Context, string, string, DatasetServiceUpdateObservationServiceRequest) (ImplResponse, error)
}
