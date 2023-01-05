package controller

import (
	"google.golang.org/grpc"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/appcontext"
)

// DatasetServiceController implements controller logic for Dataset Service endpoints
type DatasetServiceController struct {
	*MetadataController
	*LogWriterController
	*ObservationServiceController

	appCtx *appcontext.AppContext
}

// NewDatasetServiceController instantiates DatasetServiceController
func NewDatasetServiceController(
	ctx *appcontext.AppContext,
) (*grpc.Server, *DatasetServiceController) {
	gsrv := grpc.NewServer()
	srv := &DatasetServiceController{
		appCtx:                       ctx,
		MetadataController:           &MetadataController{appCtx: ctx},
		LogWriterController:          &LogWriterController{appCtx: ctx},
		ObservationServiceController: &ObservationServiceController{appCtx: ctx},
	}
	timberv1.RegisterDatasetServiceServer(gsrv, srv)

	return gsrv, srv
}
