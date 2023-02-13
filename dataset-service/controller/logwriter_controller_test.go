package controller

import (
	"context"
	"testing"

	"github.com/caraml-dev/timber/dataset-service/service"
	mlp "github.com/gojek/mlp/api/client"
	"github.com/stretchr/testify/suite"

	"github.com/caraml-dev/timber/common/errors"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/appcontext"
	"github.com/caraml-dev/timber/dataset-service/mlp/mocks"
)

type LogWriterControllerTestSuite struct {
	suite.Suite
	ctrl *LogWriterController
}

func (s *LogWriterControllerTestSuite) SetupSuite() {
	s.Suite.T().Log("Setting up LogWriterControllerTestSuite")

	// Create mock MLP service and set up with test responses
	mlpSvc := &mocks.Client{}
	projectID := int64(0)
	expectedProject := &mlp.Project{ID: 0}
	mlpSvc.On("GetProject", projectID).Return(expectedProject, nil)
	mlpSvc.On(
		"GetProject", int64(3),
	).Return(nil, errors.Newf(errors.NotFound, "MLP Project info for id %d not found in the cache", int64(3)))

	s.ctrl = &LogWriterController{
		appCtx: &appcontext.AppContext{
			Services: service.Services{
				MLPService: mlpSvc,
			},
		},
	}
}

func TestLogWriterControllerTestSuite(t *testing.T) {
	suite.Run(t, new(LogWriterControllerTestSuite))
}

func (s *LogWriterControllerTestSuite) TestListLogWriters() {
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
			err:       "MLP Project info for id 3 not found in the cache",
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

func (s *LogWriterControllerTestSuite) TestGetLogWriter() {
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
			err:       "MLP Project info for id 3 not found in the cache",
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

func (s *LogWriterControllerTestSuite) TestCreateLogWriter() {
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
			err:       "MLP Project info for id 3 not found in the cache",
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

func (s *LogWriterControllerTestSuite) TestUpdateLogWriter() {
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
			err:       "MLP Project info for id 3 not found in the cache",
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
