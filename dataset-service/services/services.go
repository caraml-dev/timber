package services

// Services contain all instantiated Service layer interfaces
type Services struct {
	MLPService MLPService
}

// NewServices instantiates Services
func NewServices(
	mlpSvc MLPService,
) Services {
	return Services{
		MLPService: mlpSvc,
	}
}
