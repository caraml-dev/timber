package service

import (
	"github.com/caraml-dev/timber/dataset-service/mlp"
)

// Services contain all instantiated Service layer interfaces
type Services struct {
	MLPService         mlp.Client
	ObservationService ObservationService
	LogWriterService   LogWriterService
}

// NewServices instantiates Services
func NewServices(
	mlpSvc mlp.Client,
	obsSvc ObservationService,
	logWriterSvc LogWriterService,
) Services {
	return Services{
		MLPService:         mlpSvc,
		ObservationService: obsSvc,
		LogWriterService:   logWriterSvc,
	}
}
