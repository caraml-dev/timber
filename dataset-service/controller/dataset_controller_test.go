package controller

import (
	"context"
	"testing"

	mlp "github.com/gojek/mlp/api/client"
	"github.com/stretchr/testify/suite"

	"github.com/caraml-dev/timber/common/errors"
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
	mlpSvc.On(
		"GetProject", int64(3),
	).Return(nil, errors.Newf(errors.NotFound, "MLP Project info for id %d not found in the cache", int64(3)))

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
	tests := []struct {
		name      string
		projectID int64
		req       *timberv1.ListLogMetadataRequest
		resp      *timberv1.ListLogMetadataResponse
		err       string
	}{
		{
			name:      "success",
			projectID: 0,
			req:       &timberv1.ListLogMetadataRequest{},
			resp:      &timberv1.ListLogMetadataResponse{},
		},
		{
			name:      "failure | project not found",
			projectID: 3,
			req:       &timberv1.ListLogMetadataRequest{ProjectId: int64(3)},
			err:       "Failed getting projectID (3) from MLP: MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		resp, err := s.ctrl.ListLogMetadata(ctx, data.req)
		if data.err == "" {
			s.Suite.Assert().NoError(err)
			s.Suite.Assert().Equal(data.resp, resp)
		} else {
			s.Suite.Assert().EqualError(err, data.err)
		}
	}
}

func (s *DataseServicetControllerTestSuite) TestGetLogMetadata() {
	ctx := context.Background()
	tests := []struct {
		name      string
		projectID int64
		req       *timberv1.GetLogMetadataRequest
		resp      *timberv1.GetLogMetadataResponse
		err       string
	}{
		{
			name:      "success",
			projectID: 0,
			req:       &timberv1.GetLogMetadataRequest{},
			resp:      &timberv1.GetLogMetadataResponse{},
		},
		{
			name:      "failure | project not found",
			projectID: 3,
			req:       &timberv1.GetLogMetadataRequest{ProjectId: int64(3)},
			err:       "Failed getting projectID (3) from MLP: MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		resp, err := s.ctrl.GetLogMetadata(ctx, data.req)
		if data.err == "" {
			s.Suite.Assert().NoError(err)
			s.Suite.Assert().Equal(data.resp, resp)
		} else {
			s.Suite.Assert().EqualError(err, data.err)
		}
	}
}

func (s *DataseServicetControllerTestSuite) TestListLogWriters() {
	ctx := context.Background()
	tests := []struct {
		name      string
		projectID int64
		req       *timberv1.ListLogWritersRequest
		resp      *timberv1.ListLogWritersResponse
		err       string
	}{
		{
			name:      "success",
			projectID: 0,
			req:       &timberv1.ListLogWritersRequest{},
			resp:      &timberv1.ListLogWritersResponse{},
		},
		{
			name:      "failure | project not found",
			projectID: 3,
			req:       &timberv1.ListLogWritersRequest{ProjectId: int64(3)},
			err:       "Failed getting projectID (3) from MLP: MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		resp, err := s.ctrl.ListLogWriters(ctx, data.req)
		if data.err == "" {
			s.Suite.Assert().NoError(err)
			s.Suite.Assert().Equal(data.resp, resp)
		} else {
			s.Suite.Assert().EqualError(err, data.err)
		}
	}
}

func (s *DataseServicetControllerTestSuite) TestGetLogWriter() {
	ctx := context.Background()
	tests := []struct {
		name      string
		projectID int64
		req       *timberv1.GetLogWriterRequest
		resp      *timberv1.GetLogWriterResponse
		err       string
	}{
		{
			name:      "success",
			projectID: 0,
			req:       &timberv1.GetLogWriterRequest{},
			resp:      &timberv1.GetLogWriterResponse{},
		},
		{
			name:      "failure | project not found",
			projectID: 3,
			req:       &timberv1.GetLogWriterRequest{ProjectId: int64(3)},
			err:       "Failed getting projectID (3) from MLP: MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		resp, err := s.ctrl.GetLogWriter(ctx, data.req)
		if data.err == "" {
			s.Suite.Assert().NoError(err)
			s.Suite.Assert().Equal(data.resp, resp)
		} else {
			s.Suite.Assert().EqualError(err, data.err)
		}
	}
}

func (s *DataseServicetControllerTestSuite) TestCreateLogWriter() {
	ctx := context.Background()
	tests := []struct {
		name      string
		projectID int64
		req       *timberv1.CreateLogWriterRequest
		resp      *timberv1.CreateLogWriterResponse
		err       string
	}{
		{
			name:      "success",
			projectID: 0,
			req:       &timberv1.CreateLogWriterRequest{},
			resp:      &timberv1.CreateLogWriterResponse{},
		},
		{
			name:      "failure | project not found",
			projectID: 3,
			req:       &timberv1.CreateLogWriterRequest{ProjectId: int64(3)},
			err:       "Failed getting projectID (3) from MLP: MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		resp, err := s.ctrl.CreateLogWriter(ctx, data.req)
		if data.err == "" {
			s.Suite.Assert().NoError(err)
			s.Suite.Assert().Equal(data.resp, resp)
		} else {
			s.Suite.Assert().EqualError(err, data.err)
		}
	}
}

func (s *DataseServicetControllerTestSuite) TestUpdateLogWriter() {
	ctx := context.Background()
	tests := []struct {
		name      string
		projectID int64
		req       *timberv1.UpdateLogWriterRequest
		resp      *timberv1.UpdateLogWriterResponse
		err       string
	}{
		{
			name:      "success",
			projectID: 0,
			req:       &timberv1.UpdateLogWriterRequest{},
			resp:      &timberv1.UpdateLogWriterResponse{},
		},
		{
			name:      "failure | project not found",
			projectID: 3,
			req:       &timberv1.UpdateLogWriterRequest{ProjectId: int64(3)},
			err:       "Failed getting projectID (3) from MLP: MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		resp, err := s.ctrl.UpdateLogWriter(ctx, data.req)
		if data.err == "" {
			s.Suite.Assert().NoError(err)
			s.Suite.Assert().Equal(data.resp, resp)
		} else {
			s.Suite.Assert().EqualError(err, data.err)
		}
	}
}

func (s *DataseServicetControllerTestSuite) TestListObservationServices() {
	ctx := context.Background()
	tests := []struct {
		name      string
		projectID int64
		req       *timberv1.ListObservationServicesRequest
		resp      *timberv1.ListObservationServicesResponse
		err       string
	}{
		{
			name:      "success",
			projectID: 0,
			req:       &timberv1.ListObservationServicesRequest{},
			resp:      &timberv1.ListObservationServicesResponse{},
		},
		{
			name:      "failure | project not found",
			projectID: 3,
			req:       &timberv1.ListObservationServicesRequest{ProjectId: int64(3)},
			err:       "Failed getting projectID (3) from MLP: MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		resp, err := s.ctrl.ListObservationServices(ctx, data.req)
		if data.err == "" {
			s.Suite.Assert().NoError(err)
			s.Suite.Assert().Equal(data.resp, resp)
		} else {
			s.Suite.Assert().EqualError(err, data.err)
		}
	}
}

func (s *DataseServicetControllerTestSuite) TestGetObservationService() {
	ctx := context.Background()
	tests := []struct {
		name      string
		projectID int64
		req       *timberv1.GetObservationServiceRequest
		resp      *timberv1.GetObservationServiceResponse
		err       string
	}{
		{
			name:      "success",
			projectID: 0,
			req:       &timberv1.GetObservationServiceRequest{},
			resp:      &timberv1.GetObservationServiceResponse{},
		},
		{
			name:      "failure | project not found",
			projectID: 3,
			req:       &timberv1.GetObservationServiceRequest{ProjectId: int64(3)},
			err:       "Failed getting projectID (3) from MLP: MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		resp, err := s.ctrl.GetObservationService(ctx, data.req)
		if data.err == "" {
			s.Suite.Assert().NoError(err)
			s.Suite.Assert().Equal(data.resp, resp)
		} else {
			s.Suite.Assert().EqualError(err, data.err)
		}
	}
}

func (s *DataseServicetControllerTestSuite) TestCreateObservationService() {
	ctx := context.Background()
	tests := []struct {
		name      string
		projectID int64
		req       *timberv1.CreateObservationServiceRequest
		resp      *timberv1.CreateObservationServiceResponse
		err       string
	}{
		{
			name:      "success",
			projectID: 0,
			req:       &timberv1.CreateObservationServiceRequest{},
			resp:      &timberv1.CreateObservationServiceResponse{},
		},
		{
			name:      "failure | project not found",
			projectID: 3,
			req:       &timberv1.CreateObservationServiceRequest{ProjectId: int64(3)},
			err:       "Failed getting projectID (3) from MLP: MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		resp, err := s.ctrl.CreateObservationService(ctx, data.req)
		if data.err == "" {
			s.Suite.Assert().NoError(err)
			s.Suite.Assert().Equal(data.resp, resp)
		} else {
			s.Suite.Assert().EqualError(err, data.err)
		}
	}
}

func (s *DataseServicetControllerTestSuite) TestUpdateObservationService() {
	ctx := context.Background()
	tests := []struct {
		name      string
		projectID int64
		req       *timberv1.UpdateObservationServiceRequest
		resp      *timberv1.UpdateObservationServiceResponse
		err       string
	}{
		{
			name:      "success",
			projectID: 0,
			req:       &timberv1.UpdateObservationServiceRequest{},
			resp:      &timberv1.UpdateObservationServiceResponse{},
		},
		{
			name:      "failure | project not found",
			projectID: 3,
			req:       &timberv1.UpdateObservationServiceRequest{ProjectId: int64(3)},
			err:       "Failed getting projectID (3) from MLP: MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		resp, err := s.ctrl.UpdateObservationService(ctx, data.req)
		if data.err == "" {
			s.Suite.Assert().NoError(err)
			s.Suite.Assert().Equal(data.resp, resp)
		} else {
			s.Suite.Assert().EqualError(err, data.err)
		}
	}
}
