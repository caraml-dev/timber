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

// LogWriterController implements controller logic for Dataset Service log writer endpoints
type LogWriterController struct {
	mlpClient        mlp.Client
	storage          storage.LogWriter
	logWriterService service.LogWriterService
}

// NewLogWriterController instantiates LogWriterController
func NewLogWriterController(logWriterStorage storage.LogWriter,
	logWriterService service.LogWriterService,
	mlpClient mlp.Client,
) *LogWriterController {
	return &LogWriterController{
		mlpClient:        mlpClient,
		logWriterService: logWriterService,
		storage:          logWriterStorage,
	}
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

	logWriters, err := l.storage.List(c, storage.ListInputFromOption(r.ProjectId, r.List))
	if err != nil {
		return nil, err
	}

	// convert to list of protos
	logWriterProtos := make([]*timberv1.LogWriter, len(logWriters))
	for i, l := range logWriters {
		logWriterProtos[i] = l.ToLogWriterProto()
	}

	return &timberv1.ListLogWritersResponse{
		LogWriters: logWriterProtos,
	}, nil
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

	lw, err := l.storage.Get(c, storage.GetInput{
		ID:        r.Id,
		ProjectID: r.ProjectId,
	})
	if err != nil {
		return nil, err
	}

	response := &timberv1.GetLogWriterResponse{
		LogWriter: lw.ToLogWriterProto(),
	}
	return response, nil
}

// CreateLogWriter definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (l *LogWriterController) CreateLogWriter(
	c context.Context,
	r *timberv1.CreateLogWriterRequest,
) (*timberv1.CreateLogWriterResponse, error) {
	// Check if the projectId is valid
	projectID := r.GetProjectId()
	project, err := l.mlpClient.GetProject(projectID)
	if err != nil {
		return nil, fmt.Errorf("error finding project %d: %w", projectID, err)
	}

	logWriter, err := l.createLogWriter(c, project.Name, model.LogWriterFromProto(r.LogWriter))
	if err != nil {
		log.Errorf("error creating log writer: %v", err)
		return nil, err
	}

	response := &timberv1.CreateLogWriterResponse{
		LogWriter: logWriter.ToLogWriterProto(),
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
	project, err := l.mlpClient.GetProject(projectID)
	if err != nil {
		return nil, fmt.Errorf("error finding project %d: %w", projectID, err)
	}

	targetStatus := r.LogWriter.Status
	if targetStatus != timberv1.Status_STATUS_UNINSTALLED && targetStatus != timberv1.Status_STATUS_DEPLOYED {
		return nil, dserrors.NewInvalidInputErrorf("invalid expected status: %s", targetStatus)
	}

	_, err = l.storage.Get(c, storage.GetInput{ID: r.Id, ProjectID: r.ProjectId})
	if err != nil {
		return nil, err
	}

	logWriter, err := l.updateOrDeleteLogWriter(c, project.Name, model.LogWriterFromProto(r.LogWriter))
	if err != nil {
		return nil, fmt.Errorf("error updating logwriter: %w", err)
	}

	response := &timberv1.UpdateLogWriterResponse{LogWriter: logWriter.ToLogWriterProto()}
	return response, nil
}

// createLogWriter install log writer deployment and update the storage.
// The long-running operation is performed in the background and the function will return with log writer having pending status
func (l *LogWriterController) createLogWriter(ctx context.Context, project string, logWriter *model.LogWriter) (*model.LogWriter, error) {
	// InstallOrUpgrade new log writer entry in DB with pending state
	logWriter.Status = model.StatusPending
	logWriter, err := l.storage.Create(ctx, logWriter)
	if err != nil {
		return nil, err
	}

	// copy request to avoid data race
	var logWriterCopy model.LogWriter
	err = copier.CopyWithOption(&logWriterCopy, logWriter, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, err
	}

	// Deploy new log writer async
	go func(project string, logWriter *model.LogWriter) {
		deployedLogWriter, err := l.logWriterService.InstallOrUpgrade(project, logWriter)
		// use separate context as the original one must have been completed
		bgCtx := context.Background()
		if err != nil {
			// If deployment failed, we'll update the status to fail and populate the error message
			logWriter.Status = model.StatusFailed
			logWriter.Error = err.Error()

			_, err = l.storage.Update(bgCtx, logWriter)
			if err != nil {
				log.Errorf("error updating log writer status to failed: %v", err)
			}
			return
		}

		// Update log writer with values returned by the log writer service creation
		_, err = l.storage.Update(bgCtx, deployedLogWriter)
		if err != nil {
			log.Errorf("error updating log writer status: %v", err)
		}
	}(project, &logWriterCopy)

	// Immediately return
	return logWriter, nil
}

// updateOrDeleteLogWriter update or uninstall log writer and update the storage accordingly.
// The long-running operation is performed in the background and the function will return with log writer having pending status
func (l *LogWriterController) updateOrDeleteLogWriter(ctx context.Context, project string, logWriter *model.LogWriter) (*model.LogWriter, error) {
	targetStatus := logWriter.Status
	// Update log writer entry in DB with pending state
	logWriter.Status = model.StatusPending
	logWriter, err := l.storage.Update(ctx, logWriter)
	if err != nil {
		return nil, err
	}

	// copy request to avoid data race
	var logWriterCopy model.LogWriter
	err = copier.CopyWithOption(&logWriterCopy, logWriter, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, err
	}

	// Update or uninstall logwriter
	go func(project string, logWriter *model.LogWriter) {
		var updatedLogWriter *model.LogWriter

		if targetStatus == model.StatusDeployed {
			updatedLogWriter, err = l.logWriterService.InstallOrUpgrade(project, logWriter)
		} else {
			updatedLogWriter, err = l.logWriterService.Uninstall(project, logWriter)
		}

		// use separate context as the original one must have been completed
		bgCtx := context.Background()
		if err != nil {
			// If deployment failed, we'll update the status to fail and populate the error message
			logWriter.Status = model.StatusFailed
			logWriter.Error = fmt.Sprintf("failed setting log writer to state %s: %s", targetStatus, err.Error())

			_, err = l.storage.Update(bgCtx, logWriter)
			if err != nil {
				log.Errorf("error updating log writer status to failed: %v", err)
			}
			return
		}

		// Update log writer with values returned by the log writer service update operation
		_, err = l.storage.Update(bgCtx, updatedLogWriter)
		if err != nil {
			log.Errorf("error updating log writer status: %v", err)
		}
	}(project, &logWriterCopy)

	// Immediately return
	return logWriter, nil
}

func (l *LogWriterController) checkProject(projectId int64) error {
	// Check if the projectId is valid
	if _, err := l.mlpClient.GetProject(projectId); err != nil {
		return err
	}

	return nil
}
