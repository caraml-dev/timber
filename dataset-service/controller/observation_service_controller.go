package controller

import (
	"context"

	"github.com/caraml-dev/timber/common/log"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/appcontext"
)

// ObservationServiceController implements controller logic for Dataset Service observation service endpoints
type ObservationServiceController struct {
	appCtx *appcontext.AppContext
}

// NewObservationServiceController instantiates ObservationServiceController
func NewObservationServiceController(ctx *appcontext.AppContext) *ObservationServiceController {
	return &ObservationServiceController{appCtx: ctx}
}

// ListObservationServices definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (o ObservationServiceController) ListObservationServices(
	c context.Context,
	r *timberv1.ListObservationServicesRequest,
) (*timberv1.ListObservationServicesResponse, error) {
	// Check if the projectId is valid
	err := o.checkProject(r.GetProjectId())
	if err != nil {
		return nil, err
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/ListObservationServices")
	response := &timberv1.ListObservationServicesResponse{}
	return response, nil
}

// GetObservationService definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (o ObservationServiceController) GetObservationService(
	c context.Context,
	r *timberv1.GetObservationServiceRequest,
) (*timberv1.GetObservationServiceResponse, error) {
	// Check if the projectId is valid
	err := o.checkProject(r.GetProjectId())
	if err != nil {
		return nil, err
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/GetObservationService")
	response := &timberv1.GetObservationServiceResponse{}
	return response, nil
}

// CreateObservationService definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (o ObservationServiceController) CreateObservationService(
	c context.Context,
	r *timberv1.CreateObservationServiceRequest,
) (*timberv1.CreateObservationServiceResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	project, err := o.appCtx.Services.MLPService.GetProject(projectID)
	if err != nil {
		return nil, err
	}

	result, err := o.appCtx.Services.ObservationService.CreateService(project.Name, r.GetObservationService())
	if err != nil {
		return nil, err
	}

	resp := &timberv1.CreateObservationServiceResponse{ObservationService: result.ToApiSchema()}
	return resp, nil
}

// UpdateObservationService definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (o ObservationServiceController) UpdateObservationService(
	c context.Context,
	r *timberv1.UpdateObservationServiceRequest,
) (*timberv1.UpdateObservationServiceResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	project, err := o.appCtx.Services.MLPService.GetProject(projectID)
	if err != nil {
		return nil, err
	}

	result, err := o.appCtx.Services.ObservationService.UpdateService(project.Name, int(r.GetId()), r.GetObservationService())
	if err != nil {
		return nil, err
	}

	resp := &timberv1.UpdateObservationServiceResponse{ObservationService: result.ToApiSchema()}
	return resp, nil
}

func (o ObservationServiceController) checkProject(projectId int64) error {
	// Check if the projectId is valid
	if _, err := o.appCtx.Services.MLPService.GetProject(projectId); err != nil {
		return err
	}

	return nil
}
