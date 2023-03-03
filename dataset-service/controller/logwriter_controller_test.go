package controller

import (
	"context"
	"testing"

	mlp "github.com/gojek/mlp/api/client"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/caraml-dev/timber/dataset-service/model"
	"github.com/caraml-dev/timber/dataset-service/storage"

	"github.com/caraml-dev/timber/common/errors"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	mlpMock "github.com/caraml-dev/timber/dataset-service/mlp/mocks"
	svcMock "github.com/caraml-dev/timber/dataset-service/service/mocks"
	storageMock "github.com/caraml-dev/timber/dataset-service/storage/mocks"
)

type LogWriterControllerTestSuite struct {
	suite.Suite
	ctrl *LogWriterController
}

var logWriterStub = &model.LogWriter{
	Base: model.Base{
		ID:        1,
		ProjectID: 1,
	},
	Source: &model.LogWriterSource{LogWriterSource: &timberv1.LogWriterSource{
		Type: timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG,
		RouterLogSource: &timberv1.RouterLogSource{
			RouterId:   1,
			RouterName: "router-1",
		},
	}},
	Name:   "log-writer-1",
	Status: model.StatusDeployed,
}

var pendingLogWriter = &model.LogWriter{
	Base: model.Base{
		ID:        1,
		ProjectID: 1,
	},
	Source: &model.LogWriterSource{LogWriterSource: &timberv1.LogWriterSource{
		Type: timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG,
		RouterLogSource: &timberv1.RouterLogSource{
			RouterId:   1,
			RouterName: "router-1",
		},
	}},
	Name:   "log-writer-1",
	Status: model.StatusPending,
}

func (s *LogWriterControllerTestSuite) SetupSuite() {
	s.Suite.T().Log("Setting up LogWriterControllerTestSuite")

	// mock MLP service and set up with test responses
	mlpSvc := &mlpMock.Client{}
	projectID := int64(1)
	projectName := "my-project"
	expectedProject := &mlp.Project{ID: 1, Name: projectName}
	mlpSvc.On("GetProject", projectID).Return(expectedProject, nil)
	mlpSvc.On("GetProject", int64(3)).
		Return(nil,
			errors.Newf(errors.NotFound, "MLP Project info for id %d not found in the cache", int64(3)))

	logWriterSvcMock := &svcMock.LogWriterService{}
	logWriterSvcMock.On("InstallOrUpgrade", projectName, mock.Anything).Return(logWriterStub, nil)
	logWriterSvcMock.On("Update", projectName, mock.Anything).Return(logWriterStub, nil)

	logWriterStorageMock := &storageMock.LogWriter{}
	logWriterStorageMock.On("Get", mock.Anything, storage.GetInput{ID: logWriterStub.ID, ProjectID: logWriterStub.ProjectID}).
		Return(logWriterStub, nil)
	logWriterStorageMock.On("List", mock.Anything, storage.ListInput{ProjectID: logWriterStub.ProjectID, Offset: 0, Limit: 10}).
		Return([]*model.LogWriter{logWriterStub}, nil)
	logWriterStorageMock.On("Create", mock.Anything, pendingLogWriter).
		Return(pendingLogWriter, nil)
	logWriterStorageMock.On("Update", mock.Anything, logWriterStub).
		Return(logWriterStub, nil)
	logWriterStorageMock.On("Update", mock.Anything, pendingLogWriter).
		Return(pendingLogWriter, nil)

	s.ctrl = &LogWriterController{
		mlpClient:        mlpSvc,
		logWriterService: logWriterSvcMock,
		storage:          logWriterStorageMock,
	}
}

func TestLogWriterControllerTestSuite(t *testing.T) {
	suite.Run(t, new(LogWriterControllerTestSuite))
}

func (s *LogWriterControllerTestSuite) TestListLogWriters() {
	ctx := context.Background()
	tests := []struct {
		name string
		req  *timberv1.ListLogWritersRequest
		resp *timberv1.ListLogWritersResponse
		err  string
	}{
		{
			name: "success",
			req: &timberv1.ListLogWritersRequest{
				ProjectId: 1,
				List: &timberv1.ListOption{
					Offset: 0,
					Limit:  10,
				},
			},
			resp: &timberv1.ListLogWritersResponse{
				LogWriters: []*timberv1.LogWriter{
					logWriterStub.ToLogWriterProto(),
				},
			},
		},
		{
			name: "failure | project not found",
			req:  &timberv1.ListLogWritersRequest{ProjectId: int64(3)},
			err:  "MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		s.Run(data.name, func() {
			resp, err := s.ctrl.ListLogWriters(ctx, data.req)
			if data.err == "" {
				s.Suite.Assert().NoError(err)
				s.Suite.Assert().Equal(data.resp, resp)
			} else {
				s.Suite.Assert().EqualError(err, data.err)
			}
		})
	}
}

func (s *LogWriterControllerTestSuite) TestGetLogWriter() {
	ctx := context.Background()
	tests := []struct {
		name string
		req  *timberv1.GetLogWriterRequest
		resp *timberv1.GetLogWriterResponse
		err  string
	}{
		{
			name: "success",
			req: &timberv1.GetLogWriterRequest{
				ProjectId: 1,
				Id:        1,
			},
			resp: &timberv1.GetLogWriterResponse{
				LogWriter: logWriterStub.ToLogWriterProto(),
			},
		},
		{
			name: "failure | project not found",
			req:  &timberv1.GetLogWriterRequest{ProjectId: int64(3)},
			err:  "MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		s.Run(data.name, func() {
			resp, err := s.ctrl.GetLogWriter(ctx, data.req)
			if data.err == "" {
				s.Suite.Assert().NoError(err)
				s.Suite.Assert().Equal(data.resp, resp)
			} else {
				s.Suite.Assert().EqualError(err, data.err)
			}
		})
	}
}

func (s *LogWriterControllerTestSuite) TestCreateLogWriter() {
	ctx := context.Background()
	tests := []struct {
		name string
		req  *timberv1.CreateLogWriterRequest
		resp *timberv1.CreateLogWriterResponse
		err  string
	}{
		{
			name: "success",
			req: &timberv1.CreateLogWriterRequest{
				ProjectId: 1,
				LogWriter: &timberv1.LogWriter{
					Id:        1,
					ProjectId: 1,
					Name:      "log-writer-1",
					Source:    logWriterStub.Source.LogWriterSource,
				},
			},
			resp: &timberv1.CreateLogWriterResponse{
				LogWriter: pendingLogWriter.ToLogWriterProto(),
			},
		},
		{
			name: "failure: project not found",
			req: &timberv1.CreateLogWriterRequest{
				ProjectId: int64(3),
				LogWriter: &timberv1.LogWriter{
					ProjectId: 3,
					Name:      "log_writer",
				}},
			err: "error finding project 3: MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		s.Run(data.name, func() {
			resp, err := s.ctrl.CreateLogWriter(ctx, data.req)
			if data.err == "" {
				s.Suite.Assert().NoError(err)
				s.Suite.Assert().Equal(data.resp, resp)
			} else {
				s.Suite.Assert().EqualError(err, data.err)
			}
		})
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
			projectID: 1,
			req: &timberv1.UpdateLogWriterRequest{
				ProjectId: 1,
				Id:        1,
				LogWriter: &timberv1.LogWriter{
					Id:        1,
					ProjectId: 1,
					Name:      "log-writer-1",
					Source:    logWriterStub.Source.LogWriterSource,
					Status:    timberv1.Status_STATUS_DEPLOYED,
				},
			},
			resp: &timberv1.UpdateLogWriterResponse{
				LogWriter: pendingLogWriter.ToLogWriterProto(),
			},
		},
		{
			name:      "failure: project not found",
			projectID: 3,
			req:       &timberv1.UpdateLogWriterRequest{ProjectId: int64(3)},
			err:       "error finding project 3: MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		s.Run(data.name, func() {
			resp, err := s.ctrl.UpdateLogWriter(ctx, data.req)
			if data.err == "" {
				s.Suite.Assert().NoError(err)
				s.Suite.Assert().Equal(data.resp, resp)
			} else {
				s.Suite.Assert().EqualError(err, data.err)
			}
		})
	}
}
