package models

import timberv1 "github.com/caraml-dev/timber/dataset-service/api"

// ObservationServiceResponse describes the data returned by Observation Service layer Create/Update methods
type ObservationServiceResponse struct {
	Id string
}

// ToApiSchema converts structure to expected Proto response
func (s *ObservationServiceResponse) ToApiSchema() *timberv1.ObservationServiceResponse {
	return &timberv1.ObservationServiceResponse{
		Id: s.Id,
	}
}
