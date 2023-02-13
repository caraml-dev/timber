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
func (l *LogWriterController) ListLogWriters(
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
func (l *LogWriterController) GetLogWriter(
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
func (l *LogWriterController) CreateLogWriter(
	c context.Context,
	r *timberv1.CreateLogWriterRequest,
) (*timberv1.CreateLogWriterResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	project, err := l.appCtx.Services.MLPService.GetProject(projectID)
	if err != nil {
		log.Errorf("error finding project: %v", err)
		return nil, err
	}

	logWriter, err := l.appCtx.Services.LogWriterService.Create(project.Name, r.LogWriter)
	if err != nil {
		log.Errorf("error creating logwriter: %v", err)
		return nil, err
	}

	response := &timberv1.CreateLogWriterResponse{
		LogWriter: logWriter,
	}
	return response, nil
}

// UpdateLogWriter definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (l *LogWriterController) UpdateLogWriter(
	c context.Context,
	r *timberv1.UpdateLogWriterRequest,
) (*timberv1.UpdateLogWriterResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	project, err := l.appCtx.Services.MLPService.GetProject(projectID)
	if err != nil {
		log.Errorf("error finding project: %v", err)
		return nil, err
	}

	logWriter, err := l.appCtx.Services.LogWriterService.Update(project.Name, r.LogWriter)
	if err != nil {
		log.Errorf("error updating logwriter: %v", err)
		return nil, err
	}

	response := &timberv1.UpdateLogWriterResponse{LogWriter: logWriter}
	return response, nil
}

func (l *LogWriterController) checkProject(projectId int64) error {
	// Check if the projectId is valid
	if _, err := l.appCtx.Services.MLPService.GetProject(projectId); err != nil {
		return err
	}

	return nil
}
