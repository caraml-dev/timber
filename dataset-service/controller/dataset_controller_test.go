package controller

import (
	"context"
	"testing"

	mlp "github.com/gojek/mlp/api/client"
	"github.com/stretchr/testify/suite"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/appcontext"
	"github.com/caraml-dev/timber/dataset-service/services"
	"github.com/caraml-dev/timber/dataset-service/services/mocks"
)

type DataseServicetControllerTestSuite struct {
	suite.Suite
	ctrl *DatasetServiceController
}

func (s *DataseServicetControllerTestSuite) SetupSuite() {
	s.Suite.T().Log("Setting up DataseServicetControllerTestSuite")

	// Create mock MLP service and set up with test responses
	mlpSvc := &mocks.MLPService{}
	projectID := int64(0)
	expectedProject := &mlp.Project{Id: 0}
	mlpSvc.On("GetProject", projectID).Return(expectedProject, nil)

	s.ctrl = &DatasetServiceController{
		appCtx: &appcontext.AppContext{
			Services: services.Services{
				MLPService: mlpSvc,
			},
		},
	}
}

func TestDatasetServiceController(t *testing.T) {
	suite.Run(t, new(DataseServicetControllerTestSuite))
}

func (s *DataseServicetControllerTestSuite) TestListLogMetadata() {
	ctx := context.Background()
	req := &timberv1.ListLogMetadataRequest{}
	expected := &timberv1.ListLogMetadataResponse{}

	resp, err := s.ctrl.ListLogMetadata(ctx, req)
	s.Suite.Require().NoError(err)
	s.Suite.Assert().Equal(expected, resp)
}

func (s *DataseServicetControllerTestSuite) TestGetLogMetadata() {
	ctx := context.Background()
	req := &timberv1.GetLogMetadataRequest{}
	expected := &timberv1.GetLogMetadataResponse{}

	resp, err := s.ctrl.GetLogMetadata(ctx, req)
	s.Suite.Require().NoError(err)
	s.Suite.Assert().Equal(expected, resp)
}

func (s *DataseServicetControllerTestSuite) TestListLogWriters() {
	ctx := context.Background()
	req := &timberv1.ListLogWritersRequest{}
	expected := &timberv1.ListLogWritersResponse{}

	resp, err := s.ctrl.ListLogWriters(ctx, req)
	s.Suite.Require().NoError(err)
	s.Suite.Assert().Equal(expected, resp)
}

func (s *DataseServicetControllerTestSuite) TestGetLogWriter() {
	ctx := context.Background()
	req := &timberv1.GetLogWriterRequest{}
	expected := &timberv1.GetLogWriterResponse{}

	resp, err := s.ctrl.GetLogWriter(ctx, req)
	s.Suite.Require().NoError(err)
	s.Suite.Assert().Equal(expected, resp)
}

func (s *DataseServicetControllerTestSuite) TestCreateLogWriter() {
	ctx := context.Background()
	req := &timberv1.CreateLogWriterRequest{}
	expected := &timberv1.CreateLogWriterResponse{}

	resp, err := s.ctrl.CreateLogWriter(ctx, req)
	s.Suite.Require().NoError(err)
	s.Suite.Assert().Equal(expected, resp)
}

func (s *DataseServicetControllerTestSuite) TestUpdateLogWriter() {
	ctx := context.Background()
	req := &timberv1.UpdateLogWriterRequest{}
	expected := &timberv1.UpdateLogWriterResponse{}

	resp, err := s.ctrl.UpdateLogWriter(ctx, req)
	s.Suite.Require().NoError(err)
	s.Suite.Assert().Equal(expected, resp)
}

func (s *DataseServicetControllerTestSuite) TestListObservationServices() {
	ctx := context.Background()
	req := &timberv1.ListObservationServicesRequest{}
	expected := &timberv1.ListObservationServicesResponse{}

	resp, err := s.ctrl.ListObservationServices(ctx, req)
	s.Suite.Require().NoError(err)
	s.Suite.Assert().Equal(expected, resp)
}

func (s *DataseServicetControllerTestSuite) TestGetObservationService() {
	ctx := context.Background()
	req := &timberv1.GetObservationServiceRequest{}
	expected := &timberv1.GetObservationServiceResponse{}

	resp, err := s.ctrl.GetObservationService(ctx, req)
	s.Suite.Require().NoError(err)
	s.Suite.Assert().Equal(expected, resp)
}

func (s *DataseServicetControllerTestSuite) TestCreateObservationService() {
	ctx := context.Background()
	req := &timberv1.CreateObservationServiceRequest{}
	expected := &timberv1.CreateObservationServiceResponse{}

	resp, err := s.ctrl.CreateObservationService(ctx, req)
	s.Suite.Require().NoError(err)
	s.Suite.Assert().Equal(expected, resp)
}

func (s *DataseServicetControllerTestSuite) TestUpdateObservationService() {
	ctx := context.Background()
	req := &timberv1.UpdateObservationServiceRequest{}
	expected := &timberv1.UpdateObservationServiceResponse{}

	resp, err := s.ctrl.UpdateObservationService(ctx, req)
	s.Suite.Require().NoError(err)
	s.Suite.Assert().Equal(expected, resp)
}
