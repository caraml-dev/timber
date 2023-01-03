package controller

import (
	"context"

	"google.golang.org/grpc"

	"github.com/caraml-dev/timber/common/log"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/appcontext"
)

// DatasetServiceController implements controller logic for Dataset Service endpoints
type DatasetServiceController struct {
	timberv1.UnimplementedDatasetServiceServer

	appCtx *appcontext.AppContext
}

// NewDatasetServiceController instantiates DatasetServiceController
func NewDatasetServiceController(ctx *appcontext.AppContext) (*grpc.Server, *DatasetServiceController) {
	gsrv := grpc.NewServer()
	srv := &DatasetServiceController{appCtx: ctx}
	timberv1.RegisterDatasetServiceServer(gsrv, srv)

	return gsrv, srv
}

// ListLogMetadata definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) ListLogMetadata(
	context.Context,
	*timberv1.ListLogMetadataRequest,
) (*timberv1.ListLogMetadataResponse, error) {
	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/ListLogMetadata")
	response := &timberv1.ListLogMetadataResponse{}
	return response, nil
}

// GetLogMetadata definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) GetLogMetadata(
	context.Context,
	*timberv1.GetLogMetadataRequest,
) (*timberv1.GetLogMetadataResponse, error) {
	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/GetLogMetadata")
	response := &timberv1.GetLogMetadataResponse{}
	return response, nil
}

// ListLogWriters definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) ListLogWriters(
	context.Context,
	*timberv1.ListLogWritersRequest,
) (*timberv1.ListLogWritersResponse, error) {
	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/ListLogWriters")
	response := &timberv1.ListLogWritersResponse{}
	return response, nil
}

// GetLogWriter definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) GetLogWriter(
	context.Context,
	*timberv1.GetLogWriterRequest,
) (*timberv1.GetLogWriterResponse, error) {
	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/GetLogWriter")
	response := &timberv1.GetLogWriterResponse{}
	return response, nil
}

// CreateLogWriter definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) CreateLogWriter(
	context.Context,
	*timberv1.CreateLogWriterRequest,
) (*timberv1.CreateLogWriterResponse, error) {
	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/CreateLogWriter")
	response := &timberv1.CreateLogWriterResponse{}
	return response, nil
}

// UpdateLogWriter definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) UpdateLogWriter(
	context.Context,
	*timberv1.UpdateLogWriterRequest,
) (*timberv1.UpdateLogWriterResponse, error) {
	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/UpdateLogWriter")
	response := &timberv1.UpdateLogWriterResponse{}
	return response, nil
}

// ListObservationServices definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) ListObservationServices(
	context.Context,
	*timberv1.ListObservationServicesRequest,
) (*timberv1.ListObservationServicesResponse, error) {
	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/ListObservationServices")
	response := &timberv1.ListObservationServicesResponse{}
	return response, nil
}

// GetObservationService definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) GetObservationService(
	context.Context,
	*timberv1.GetObservationServiceRequest,
) (*timberv1.GetObservationServiceResponse, error) {
	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/GetObservationService")
	response := &timberv1.GetObservationServiceResponse{}
	return response, nil
}

// CreateObservationService definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) CreateObservationService(
	context.Context,
	*timberv1.CreateObservationServiceRequest,
) (*timberv1.CreateObservationServiceResponse, error) {
	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/CreateObservationService")
	response := &timberv1.CreateObservationServiceResponse{}
	return response, nil
}

// UpdateObservationService definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) UpdateObservationService(
	context.Context,
	*timberv1.UpdateObservationServiceRequest,
) (*timberv1.UpdateObservationServiceResponse, error) {
	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/UpdateObservationService")
	response := &timberv1.UpdateObservationServiceResponse{}
	return response, nil
}
