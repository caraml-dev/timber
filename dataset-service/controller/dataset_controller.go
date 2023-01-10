package controller

import (
	"google.golang.org/grpc"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/services"
)

// DatasetServiceController implements controller logic for Dataset Service endpoints
type DatasetServiceController struct {
	*MetadataController
	*LogWriterController
	*ObservationServiceController
}

// NewDatasetServiceController instantiates DatasetServiceController
func NewDatasetServiceController(
	services *services.Services,
) (*grpc.Server, *DatasetServiceController) {
	gsrv := grpc.NewServer()
	srv := &DatasetServiceController{
		MetadataController:           &MetadataController{services: services},
		LogWriterController:          &LogWriterController{services: services},
		ObservationServiceController: &ObservationServiceController{services: services},
	}
	timberv1.RegisterDatasetServiceServer(gsrv, srv)

	return gsrv, srv
}
