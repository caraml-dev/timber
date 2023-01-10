package controller

import (
	"context"
	"testing"

	mlp "github.com/gojek/mlp/api/client"
	"github.com/stretchr/testify/suite"

	"github.com/caraml-dev/timber/common/errors"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/services"
	"github.com/caraml-dev/timber/dataset-service/services/mocks"
)

type MetadataControllerTestSuite struct {
	suite.Suite
	ctrl *MetadataController
}

func (s *MetadataControllerTestSuite) SetupSuite() {
	s.Suite.T().Log("Setting up MetadataControllerTestSuite")

	// Create mock MLP service and set up with test responses
	mlpSvc := &mocks.MLPService{}
	projectID := int64(0)
	expectedProject := &mlp.Project{Id: 0}
	mlpSvc.On("GetProject", projectID).Return(expectedProject, nil)
	mlpSvc.On(
		"GetProject", int64(3),
	).Return(nil, errors.Newf(errors.NotFound, "MLP Project info for id %d not found in the cache", int64(3)))

	s.ctrl = &MetadataController{
		services: &services.Services{
			MLPService: mlpSvc,
		},
	}
}

func TestMetadataControllerTestSuite(t *testing.T) {
	suite.Run(t, new(MetadataControllerTestSuite))
}

func (s *MetadataControllerTestSuite) TestListLogMetadata() {
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
			err:       "MLP Project info for id 3 not found in the cache",
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

func (s *MetadataControllerTestSuite) TestGetLogMetadata() {
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
			err:       "MLP Project info for id 3 not found in the cache",
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
