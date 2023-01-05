package controller

import (
	"context"

	"github.com/caraml-dev/timber/common/log"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/appcontext"
)

// LogWriterController implements controller logic for Dataset Service log writer endpoints
type LogWriterController struct {
	appCtx *appcontext.AppContext
}

// NewLogWriterController instantiates LogWriterController
func NewLogWriterController(ctx *appcontext.AppContext) *LogWriterController {
	return &LogWriterController{appCtx: ctx}
}

// ListLogWriters definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (l LogWriterController) ListLogWriters(
	c context.Context,
	r *timberv1.ListLogWritersRequest,
) (*timberv1.ListLogWritersResponse, error) {
	// Check if the projectId is valid
	err := l.checkProject(r.GetProjectId())
	if err != nil {
		return nil, err
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/ListLogWriters")
	response := &timberv1.ListLogWritersResponse{}
	return response, nil
}

// GetLogWriter definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (l LogWriterController) GetLogWriter(
	c context.Context,
	r *timberv1.GetLogWriterRequest,
) (*timberv1.GetLogWriterResponse, error) {
	// Check if the projectId is valid
	err := l.checkProject(r.GetProjectId())
	if err != nil {
		return nil, err
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/GetLogWriter")
	response := &timberv1.GetLogWriterResponse{}
	return response, nil
}

// CreateLogWriter definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (l LogWriterController) CreateLogWriter(
	c context.Context,
	r *timberv1.CreateLogWriterRequest,
) (*timberv1.CreateLogWriterResponse, error) {
	// Check if the projectId is valid
	err := l.checkProject(r.GetProjectId())
	if err != nil {
		return nil, err
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/CreateLogWriter")
	response := &timberv1.CreateLogWriterResponse{}
	return response, nil
}

// UpdateLogWriter definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (l LogWriterController) UpdateLogWriter(
	c context.Context,
	r *timberv1.UpdateLogWriterRequest,
) (*timberv1.UpdateLogWriterResponse, error) {
	// Check if the projectId is valid
	err := l.checkProject(r.GetProjectId())
	if err != nil {
		return nil, err
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/UpdateLogWriter")
	response := &timberv1.UpdateLogWriterResponse{}
	return response, nil
}

func (l LogWriterController) checkProject(projectId int64) error {
	// Check if the projectId is valid
	if _, err := l.appCtx.Services.MLPService.GetProject(projectId); err != nil {
		return err
	}

	return nil
}
