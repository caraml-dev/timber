package controller

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/caraml-dev/timber/common/errors"
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
	c context.Context,
	r *timberv1.ListLogMetadataRequest,
) (*timberv1.ListLogMetadataResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	if _, err := d.appCtx.Services.MLPService.GetProject(projectID); err != nil {
		return nil, errors.Newf(errors.GetType(err), fmt.Sprintf("Failed getting projectID (%d) from MLP: %v", projectID, err))
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/ListLogMetadata")
	response := &timberv1.ListLogMetadataResponse{}
	return response, nil
}

// GetLogMetadata definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) GetLogMetadata(
	c context.Context,
	r *timberv1.GetLogMetadataRequest,
) (*timberv1.GetLogMetadataResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	if _, err := d.appCtx.Services.MLPService.GetProject(projectID); err != nil {
		return nil, errors.Newf(errors.GetType(err), fmt.Sprintf("Failed getting projectID (%d) from MLP: %v", projectID, err))
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/GetLogMetadata")
	response := &timberv1.GetLogMetadataResponse{}
	return response, nil
}

// ListLogWriters definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) ListLogWriters(
	c context.Context,
	r *timberv1.ListLogWritersRequest,
) (*timberv1.ListLogWritersResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	if _, err := d.appCtx.Services.MLPService.GetProject(projectID); err != nil {
		return nil, errors.Newf(errors.GetType(err), fmt.Sprintf("Failed getting projectID (%d) from MLP: %v", projectID, err))
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/ListLogWriters")
	response := &timberv1.ListLogWritersResponse{}
	return response, nil
}

// GetLogWriter definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) GetLogWriter(
	c context.Context,
	r *timberv1.GetLogWriterRequest,
) (*timberv1.GetLogWriterResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	if _, err := d.appCtx.Services.MLPService.GetProject(projectID); err != nil {
		return nil, errors.Newf(errors.GetType(err), fmt.Sprintf("Failed getting projectID (%d) from MLP: %v", projectID, err))
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/GetLogWriter")
	response := &timberv1.GetLogWriterResponse{}
	return response, nil
}

// CreateLogWriter definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) CreateLogWriter(
	c context.Context,
	r *timberv1.CreateLogWriterRequest,
) (*timberv1.CreateLogWriterResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	if _, err := d.appCtx.Services.MLPService.GetProject(projectID); err != nil {
		return nil, errors.Newf(errors.GetType(err), fmt.Sprintf("Failed getting projectID (%d) from MLP: %v", projectID, err))
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/CreateLogWriter")
	response := &timberv1.CreateLogWriterResponse{}
	return response, nil
}

// UpdateLogWriter definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) UpdateLogWriter(
	c context.Context,
	r *timberv1.UpdateLogWriterRequest,
) (*timberv1.UpdateLogWriterResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	if _, err := d.appCtx.Services.MLPService.GetProject(projectID); err != nil {
		return nil, errors.Newf(errors.GetType(err), fmt.Sprintf("Failed getting projectID (%d) from MLP: %v", projectID, err))
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/UpdateLogWriter")
	response := &timberv1.UpdateLogWriterResponse{}
	return response, nil
}

// ListObservationServices definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) ListObservationServices(
	c context.Context,
	r *timberv1.ListObservationServicesRequest,
) (*timberv1.ListObservationServicesResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	if _, err := d.appCtx.Services.MLPService.GetProject(projectID); err != nil {
		return nil, errors.Newf(errors.GetType(err), fmt.Sprintf("Failed getting projectID (%d) from MLP: %v", projectID, err))
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/ListObservationServices")
	response := &timberv1.ListObservationServicesResponse{}
	return response, nil
}

// GetObservationService definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) GetObservationService(
	c context.Context,
	r *timberv1.GetObservationServiceRequest,
) (*timberv1.GetObservationServiceResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	if _, err := d.appCtx.Services.MLPService.GetProject(projectID); err != nil {
		return nil, errors.Newf(errors.GetType(err), fmt.Sprintf("Failed getting projectID (%d) from MLP: %v", projectID, err))
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/GetObservationService")
	response := &timberv1.GetObservationServiceResponse{}
	return response, nil
}

// CreateObservationService definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) CreateObservationService(
	c context.Context,
	r *timberv1.CreateObservationServiceRequest,
) (*timberv1.CreateObservationServiceResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	if _, err := d.appCtx.Services.MLPService.GetProject(projectID); err != nil {
		return nil, errors.Newf(errors.GetType(err), fmt.Sprintf("Failed getting projectID (%d) from MLP: %v", projectID, err))
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/CreateObservationService")
	response := &timberv1.CreateObservationServiceResponse{}
	return response, nil
}

// UpdateObservationService definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (d DatasetServiceController) UpdateObservationService(
	c context.Context,
	r *timberv1.UpdateObservationServiceRequest,
) (*timberv1.UpdateObservationServiceResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	if _, err := d.appCtx.Services.MLPService.GetProject(projectID); err != nil {
		return nil, errors.Newf(errors.GetType(err), fmt.Sprintf("Failed getting projectID (%d) from MLP: %v", projectID, err))
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/UpdateObservationService")
	response := &timberv1.UpdateObservationServiceResponse{}
	return response, nil
}
