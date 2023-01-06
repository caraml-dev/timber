package controller

import (
	"context"

	"github.com/caraml-dev/timber/common/log"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/appcontext"
)

// MetadataController implements controller logic for Dataset Service metadata endpoints
type MetadataController struct {
	appCtx *appcontext.AppContext
}

// NewMetadataController instantiates MetadataController
func NewMetadataController(ctx *appcontext.AppContext) *MetadataController {
	return &MetadataController{appCtx: ctx}
}

// ListLogMetadata definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (m MetadataController) ListLogMetadata(
	c context.Context,
	r *timberv1.ListLogMetadataRequest,
) (*timberv1.ListLogMetadataResponse, error) {
	// Check if the projectId is valid
	err := m.checkProject(r.GetProjectId())
	if err != nil {
		return nil, err
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/ListLogMetadata")
	response := &timberv1.ListLogMetadataResponse{}
	return response, nil
}

// GetLogMetadata definition: See dataset-service/api/caraml/timber/v1/dataset_service.proto
func (m MetadataController) GetLogMetadata(
	c context.Context,
	r *timberv1.GetLogMetadataRequest,
) (*timberv1.GetLogMetadataResponse, error) {
	// Check if the projectId is valid
	err := m.checkProject(r.GetProjectId())
	if err != nil {
		return nil, err
	}

	// TODO: Implement method
	log.Info("Called caraml.upi.v1.DatasetService/GetLogMetadata")
	response := &timberv1.GetLogMetadataResponse{}
	return response, nil
}

func (m MetadataController) checkProject(projectId int64) error {
	// Check if the projectId is valid
	if _, err := m.appCtx.Services.MLPService.GetProject(projectId); err != nil {
		return err
	}

	return nil
}
