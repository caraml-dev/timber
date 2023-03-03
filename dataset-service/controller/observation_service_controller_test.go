package controller

import (
	"context"
	"fmt"
	"testing"

	mlp "github.com/gojek/mlp/api/client"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/caraml-dev/timber/dataset-service/model"
	"github.com/caraml-dev/timber/dataset-service/storage"

	"github.com/caraml-dev/timber/common/errors"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	mlpMock "github.com/caraml-dev/timber/dataset-service/mlp/mocks"
	"github.com/caraml-dev/timber/dataset-service/service/mocks"
	storageMock "github.com/caraml-dev/timber/dataset-service/storage/mocks"
)

type ObservationServiceControllerTestSuite struct {
	suite.Suite
	ctrl *ObservationServiceController
}

var observationServiceStub = &model.ObservationService{
	Base: model.Base{
		ID:        1,
		ProjectID: 1,
	},
	Source: &model.ObservationServiceSource{
		ObservationServiceSource: &timberv1.ObservationServiceSource{
			Type: timberv1.ObservationServiceSourceType_OBSERVATION_SERVICE_SOURCE_TYPE_KAFKA,
			Kafka: &timberv1.KafkaConfig{
				Brokers: "broker",
				Topic:   "topic",
			},
		},
	},
	Name:   "observation-svc",
	Status: model.StatusDeployed,
}

var pendingObservationServiceStub = &model.ObservationService{
	Base: model.Base{
		ID:        1,
		ProjectID: 1,
	},
	Source: &model.ObservationServiceSource{
		ObservationServiceSource: &timberv1.ObservationServiceSource{
			Type: timberv1.ObservationServiceSourceType_OBSERVATION_SERVICE_SOURCE_TYPE_KAFKA,
			Kafka: &timberv1.KafkaConfig{
				Brokers: "broker",
				Topic:   "topic",
			},
		},
	},
	Name:   "observation-svc",
	Status: model.StatusPending,
}

func (s *ObservationServiceControllerTestSuite) SetupSuite() {
	s.Suite.T().Log("Setting up ObservationServiceControllerTestSuite")

	// InstallOrUpgrade mock MLP service and set up with test responses
	mlpSvc := &mlpMock.Client{}
	projectID := int64(1)
	projectName := "test-project"
	expectedProject := &mlp.Project{ID: 0, Name: projectName}
	failedProjectID := int64(4)
	failedProjectName := "failed-test-project"
	expectedFailedProject := &mlp.Project{ID: 4, Name: failedProjectName}
	mlpSvc.On("GetProject", projectID).Return(expectedProject, nil)
	mlpSvc.On("GetProject", failedProjectID).Return(expectedFailedProject, nil)
	mlpSvc.On(
		"GetProject", int64(3),
	).Return(nil, errors.Newf(errors.NotFound, "MLP Project info for id %d not found in the cache", int64(3)))

	// InstallOrUpgrade mock Observation service and set up with test responses
	observationSvc := &mocks.ObservationService{}
	observationSvc.On("InstallOrUpgrade", projectName, mock.Anything).Return(observationServiceStub, nil)
	observationSvc.On("Update", projectName, mock.Anything).Return(observationServiceStub, nil)
	observationSvc.On("InstallOrUpgrade", failedProjectName, mock.Anything).Return(nil, fmt.Errorf("failed create"))
	observationSvc.On("Update", failedProjectName, mock.Anything).Return(nil, fmt.Errorf("failed update"))

	observationSvcStorage := &storageMock.ObservationService{}
	observationSvcStorage.On("Get", mock.Anything, storage.GetInput{ID: observationServiceStub.ID, ProjectID: observationServiceStub.ProjectID}).
		Return(observationServiceStub, nil)
	observationSvcStorage.On("List", mock.Anything, storage.ListInput{ProjectID: observationServiceStub.ProjectID, Offset: 0, Limit: 10}).
		Return([]*model.ObservationService{observationServiceStub}, nil)
	observationSvcStorage.On("Create", mock.Anything, pendingObservationServiceStub).
		Return(pendingObservationServiceStub, nil)
	observationSvcStorage.On("Update", mock.Anything, observationServiceStub).
		Return(observationServiceStub, nil)
	observationSvcStorage.On("Update", mock.Anything, pendingObservationServiceStub).
		Return(pendingObservationServiceStub, nil)

	s.ctrl = &ObservationServiceController{
		observationService: observationSvc,
		mlpClient:          mlpSvc,
		storage:            observationSvcStorage,
	}
}

func TestObservationServiceControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ObservationServiceControllerTestSuite))
}

func (s *ObservationServiceControllerTestSuite) TestListObservationServices() {
	ctx := context.Background()
	tests := []struct {
		name string
		req  *timberv1.ListObservationServicesRequest
		resp *timberv1.ListObservationServicesResponse
		err  string
	}{
		{
			name: "success",
			req: &timberv1.ListObservationServicesRequest{
				ProjectId: 1,
				List: &timberv1.ListOption{
					Offset: 0,
					Limit:  10,
				},
			},
			resp: &timberv1.ListObservationServicesResponse{
				ObservationServices: []*timberv1.ObservationService{
					observationServiceStub.ToObservationServiceProto(),
				},
			},
		},
		{
			name: "failure: project not found",
			req:  &timberv1.ListObservationServicesRequest{ProjectId: int64(3)},
			err:  "MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		s.Run(data.name, func() {
			resp, err := s.ctrl.ListObservationServices(ctx, data.req)
			if data.err == "" {
				s.Suite.Assert().NoError(err)
				s.Suite.Assert().Equal(resp, data.resp)
			} else {
				s.Suite.Assert().EqualError(err, data.err)
			}
		})
	}
}

func (s *ObservationServiceControllerTestSuite) TestGetObservationService() {
	ctx := context.Background()
	tests := []struct {
		name string
		req  *timberv1.GetObservationServiceRequest
		resp *timberv1.GetObservationServiceResponse
		err  string
	}{
		{
			name: "success",
			req: &timberv1.GetObservationServiceRequest{
				ProjectId: 1,
				Id:        1,
			},
			resp: &timberv1.GetObservationServiceResponse{
				ObservationService: observationServiceStub.ToObservationServiceProto(),
			},
		},
		{
			name: "failure: project not found",
			req:  &timberv1.GetObservationServiceRequest{ProjectId: int64(3)},
			err:  "MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		s.Run(data.name, func() {
			resp, err := s.ctrl.GetObservationService(ctx, data.req)
			if data.err == "" {
				s.Suite.Assert().NoError(err)
				s.Suite.Assert().Equal(resp, data.resp)
			} else {
				s.Suite.Assert().EqualError(err, data.err)
			}
		})
	}
}

func (s *ObservationServiceControllerTestSuite) TestCreateObservationService() {
	ctx := context.Background()
	tests := []struct {
		name string
		req  *timberv1.CreateObservationServiceRequest
		resp *timberv1.CreateObservationServiceResponse
		err  string
	}{
		{
			name: "success",
			req: &timberv1.CreateObservationServiceRequest{
				ProjectId: 1,
				ObservationService: &timberv1.ObservationService{
					ProjectId: 1,
					Id:        1,
					Name:      "observation-svc",
					Source:    observationServiceStub.Source.ObservationServiceSource,
				},
			},
			resp: &timberv1.CreateObservationServiceResponse{
				ObservationService: pendingObservationServiceStub.ToObservationServiceProto(),
			},
		},
		{
			name: "failure | project not found",
			req:  &timberv1.CreateObservationServiceRequest{ProjectId: int64(3)},
			err:  "MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		s.Run(data.name, func() {
			resp, err := s.ctrl.CreateObservationService(ctx, data.req)
			if data.err == "" {
				s.Suite.Assert().NoError(err)
				s.Suite.Assert().Equal(data.resp, resp)
			} else {
				s.Suite.Assert().EqualError(err, data.err)
			}
		})
	}
}

func (s *ObservationServiceControllerTestSuite) TestUpdateObservationService() {
	ctx := context.Background()
	tests := []struct {
		name string
		req  *timberv1.UpdateObservationServiceRequest
		resp *timberv1.UpdateObservationServiceResponse
		err  string
	}{
		{
			name: "success",
			req: &timberv1.UpdateObservationServiceRequest{
				ProjectId: 1,
				Id:        1,
				ObservationService: &timberv1.ObservationService{
					ProjectId: 1,
					Id:        1,
					Name:      "observation-svc",
					Source:    observationServiceStub.Source.ObservationServiceSource,
					Status:    timberv1.Status_STATUS_DEPLOYED,
				},
			},
			resp: &timberv1.UpdateObservationServiceResponse{
				ObservationService: pendingObservationServiceStub.ToObservationServiceProto(),
			},
		},
		{
			name: "failure | project not found",
			req:  &timberv1.UpdateObservationServiceRequest{Id: int64(3), ProjectId: int64(3)},
			err:  "MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		s.Run(data.name, func() {
			resp, err := s.ctrl.UpdateObservationService(ctx, data.req)
			if data.err == "" {
				s.Suite.Assert().NoError(err)
				s.Suite.Assert().Equal(data.resp, resp)
			} else {
				s.Suite.Assert().EqualError(err, data.err)
			}
		})
	}
}
