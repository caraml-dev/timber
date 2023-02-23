package controller

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/caraml-dev/timber/common/log"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	dserrors "github.com/caraml-dev/timber/dataset-service/errors"
	"github.com/caraml-dev/timber/dataset-service/mlp"
	"github.com/caraml-dev/timber/dataset-service/model"
	"github.com/caraml-dev/timber/dataset-service/service"
	"github.com/caraml-dev/timber/dataset-service/storage"
)

// ObservationServiceController implements controller logic for Dataset Service observation service endpoints
type ObservationServiceController struct {
	mlpClient          mlp.Client
	storage            storage.ObservationService
	observationService service.ObservationService
}

// NewObservationServiceController instantiates ObservationServiceController
func NewObservationServiceController(observationServiceStorage storage.ObservationService,
	observationService service.ObservationService,
	mlpClient mlp.Client) *ObservationServiceController {
	return &ObservationServiceController{
		mlpClient:          mlpClient,
		storage:            observationServiceStorage,
		observationService: observationService,
	}
}

// ListObservationServices definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (o *ObservationServiceController) ListObservationServices(
	ctx context.Context,
	r *timberv1.ListObservationServicesRequest,
) (*timberv1.ListObservationServicesResponse, error) {
	// Check if the projectId is valid
	err := o.checkProject(r.GetProjectId())
	if err != nil {
		return nil, err
	}

	obsSvcs, err := o.storage.List(ctx, storage.ListInputFromOption(r.ProjectId, r.List))
	if err != nil {
		return nil, err
	}

	obsSvcProtos := make([]*timberv1.ObservationService, len(obsSvcs))
	for i, s := range obsSvcs {
		obsSvcProtos[i] = s.ToObservationServiceProto()
	}

	return &timberv1.ListObservationServicesResponse{
		ObservationServices: obsSvcProtos,
	}, nil
}

// GetObservationService definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (o *ObservationServiceController) GetObservationService(
	ctx context.Context,
	r *timberv1.GetObservationServiceRequest,
) (*timberv1.GetObservationServiceResponse, error) {
	// Check if the projectId is valid
	err := o.checkProject(r.GetProjectId())
	if err != nil {
		return nil, err
	}

	obsSvc, err := o.storage.Get(ctx, storage.GetInput{
		ID:        r.Id,
		ProjectID: r.ProjectId,
	})

	if err != nil {
		return nil, err
	}

	response := &timberv1.GetObservationServiceResponse{
		ObservationService: obsSvc.ToObservationServiceProto(),
	}
	return response, nil
}

// CreateObservationService definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (o *ObservationServiceController) CreateObservationService(
	ctx context.Context,
	r *timberv1.CreateObservationServiceRequest,
) (*timberv1.CreateObservationServiceResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	project, err := o.mlpClient.GetProject(projectID)
	if err != nil {
		log.Errorf("error finding project: %v", err)
		return nil, err
	}

	result, err := o.createObservationService(ctx, project.Name, model.ObservationServiceFromProto(r.ObservationService))
	if err != nil {
		log.Errorf("error creating observation service: %v", err)
		return nil, err
	}

	resp := &timberv1.CreateObservationServiceResponse{ObservationService: result.ToObservationServiceProto()}
	return resp, nil
}

// UpdateObservationService definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (o *ObservationServiceController) UpdateObservationService(
	ctx context.Context,
	r *timberv1.UpdateObservationServiceRequest,
) (*timberv1.UpdateObservationServiceResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	project, err := o.mlpClient.GetProject(projectID)
	if err != nil {
		log.Errorf("error finding project: %v", err)
		return nil, err
	}

	targetStatus := r.ObservationService.Status
	if targetStatus != timberv1.Status_STATUS_UNINSTALLED && targetStatus != timberv1.Status_STATUS_DEPLOYED {
		return nil, dserrors.NewInvalidInputErrorf("invalid expected status: %s", targetStatus)
	}

	result, err := o.updateOrDeleteObservationService(ctx, project.Name, model.ObservationServiceFromProto(r.ObservationService))
	if err != nil {
		log.Errorf("error updating observation service: %v", err)
		return nil, err
	}

	resp := &timberv1.UpdateObservationServiceResponse{ObservationService: result.ToObservationServiceProto()}
	return resp, nil
}

// createObservationService install observation service deployment and update the storage.
// The long-running operation is performed in the background and the function will return with observation service having pending status
func (o *ObservationServiceController) createObservationService(ctx context.Context,
	project string,
	observationService *model.ObservationService,
) (*model.ObservationService, error) {
	// InstallOrUpgrade new observation service entry in DB with pending state
	observationService.Status = model.StatusPending
	observationService, err := o.storage.Create(ctx, observationService)
	if err != nil {
		return nil, err
	}

	// copy request to avoid data race
	var observationServiceCopy model.ObservationService
	err = copier.CopyWithOption(&observationServiceCopy, observationService, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, err
	}

	// Deploy new observation service async
	go func(project string, observationService *model.ObservationService) {
		deployedObservationService, err := o.observationService.InstallOrUpgrade(project, observationService)
		// use separate context as the original one must have been completed
		bgCtx := context.Background()
		if err != nil {
			// If deployment failed, we'll update the status to fail and populate the error message
			observationService.Status = model.StatusFailed
			observationService.Error = err.Error()

			_, err = o.storage.Update(bgCtx, observationService)
			if err != nil {
				log.Errorf("error updating observation service status to failed: %v", err)
			}
			return
		}

		// Update observation service with values returned by the observation service service creation
		_, err = o.storage.Update(bgCtx, deployedObservationService)
		if err != nil {
			log.Errorf("error updating observation service status: %v", err)
		}
	}(project, &observationServiceCopy)

	// Immediately return
	return observationService, nil
}

// updateOrDeleteObservationService update or uninstall observation service and update the storage accordingly.
// The long-running operation is performed in the background and the function will return with observation service having pending status
func (o *ObservationServiceController) updateOrDeleteObservationService(
	ctx context.Context,
	project string,
	observationService *model.ObservationService,
) (*model.ObservationService, error) {
	targetStatus := observationService.Status
	// Update observation service entry in DB with pending state
	observationService.Status = model.StatusPending
	observationService, err := o.storage.Update(ctx, observationService)
	if err != nil {
		return nil, err
	}

	// copy request to avoid data race
	var observationServiceCopy model.ObservationService
	err = copier.CopyWithOption(&observationServiceCopy, observationService, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, err
	}

	// Update or uninstall observation service
	go func(project string, observationService *model.ObservationService) {
		var updatedObservationService *model.ObservationService

		if targetStatus == model.StatusDeployed {
			updatedObservationService, err = o.observationService.InstallOrUpgrade(project, observationService)
		} else {
			updatedObservationService, err = o.observationService.Uninstall(project, observationService)
		}

		// use separate context as the original one must have been completed
		bgCtx := context.Background()
		if err != nil {
			// If deployment failed, we'll update the status to fail and populate the error message
			observationService.Status = model.StatusFailed
			observationService.Error = fmt.Sprintf("failed setting observation service to state %s: %s", targetStatus, err.Error())

			_, err = o.storage.Update(bgCtx, observationService)
			if err != nil {
				log.Errorf("error updating observation service status to failed: %v", err)
			}
			return
		}

		// Update observation service with values returned by the observation service service update operation
		_, err = o.storage.Update(bgCtx, updatedObservationService)
		if err != nil {
			log.Errorf("error updating observation service status: %v", err)
		}
	}(project, &observationServiceCopy)

	// Immediately return
	return observationService, nil
}

func (o *ObservationServiceController) checkProject(projectId int64) error {
	// Check if the projectId is valid
	if _, err := o.mlpClient.GetProject(projectId); err != nil {
		return err
	}

	return nil
}
