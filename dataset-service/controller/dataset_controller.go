package controller

import (
	"google.golang.org/grpc"

	"github.com/caraml-dev/timber/dataset-service/storage"

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

	logWriterStorage := storage.NewLogWriter(ctx.DB)
	logWriterController := NewLogWriterController(logWriterStorage, ctx.Services.LogWriterService, ctx.Services.MLPService)

	observationServiceStorage := storage.NewObservationService(ctx.DB)
	observationServiceController := NewObservationServiceController(observationServiceStorage, ctx.Services.ObservationService, ctx.Services.MLPService)

	srv := &DatasetServiceController{
		appCtx:                       ctx,
		MetadataController:           &MetadataController{appCtx: ctx},
		LogWriterController:          logWriterController,
		ObservationServiceController: observationServiceController,
	}
	timberv1.RegisterDatasetServiceServer(gsrv, srv)

	return gsrv, srv
}
