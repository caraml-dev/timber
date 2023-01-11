package services

// Services contain all instantiated Service layer interfaces
type Services struct {
	MLPService         MLPService
	ObservationService ObservationService
}

// NewServices instantiates Services
func NewServices(
	mlpSvc MLPService,
	obsSvc ObservationService,
) Services {
	return Services{
		MLPService:         mlpSvc,
		ObservationService: obsSvc,
	}
}
