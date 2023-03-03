package storage

import (
	"context"
	"fmt"

	"github.com/stretchr/testify/suite"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/model"
)

type LogWriterStorageTestSuite struct {
	suite.Suite
	logWriterStorage logWriter
	seedData         []*model.LogWriter
}

func (s *LogWriterStorageTestSuite) SetupSuite() {
	s.Suite.T().Log("Setting up LogWriterStorageTestSuite")
	s.logWriterStorage = logWriter{db: testDB}

	for i := 0; i < 100; i++ {
		c, err := s.logWriterStorage.Create(context.Background(), &model.LogWriter{
			Base: model.Base{
				ProjectID: 2,
			},
			Name:   fmt.Sprintf("log-writer-%d", i),
			Status: model.StatusDeployed,
			Error:  fmt.Sprintf("log-writer-%d", i),
			Source: &model.LogWriterSource{
				LogWriterSource: &timberv1.LogWriterSource{
					Type: timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG,
					RouterLogSource: &timberv1.RouterLogSource{
						RouterId:   1,
						RouterName: "router-1",
						Kafka: &timberv1.KafkaConfig{
							Brokers: "my-broker",
							Topic:   "my-topic",
						},
					},
				},
			},
		})

		s.Assert().NoError(err)
		s.seedData = append(s.seedData, c)
	}
}

func (s *LogWriterStorageTestSuite) TestGetLogWriter() {
	ctx := context.Background()

	// Test get with id
	got, err := s.logWriterStorage.Get(ctx, GetInput{
		ID:        1,
		ProjectID: 2,
	})

	s.Assert().NoError(err)
	s.Assert().Equal(s.seedData[0], got)

	// Test get with name
	got, err = s.logWriterStorage.Get(ctx, GetInput{
		Name: "log-writer-99",
	})

	s.Assert().NoError(err)
	s.Assert().Equal(s.seedData[99], got)
	s.Assert().Equal(int64(100), got.ID)
}

func (s *LogWriterStorageTestSuite) TestListLogWriter() {
	ctx := context.Background()

	// Test list first 5
	got, err := s.logWriterStorage.List(ctx, ListInput{
		ProjectID: 2,
		Offset:    0,
		Limit:     5,
	})

	s.Assert().NoError(err)
	s.Assert().Len(got, 5)
	s.Assert().ElementsMatch(got, s.seedData[:5])

	// Test list index 5 to 10
	got, err = s.logWriterStorage.List(ctx, ListInput{
		ProjectID: 2,
		Offset:    5,
		Limit:     5,
	})

	s.Assert().NoError(err)
	s.Assert().Len(got, 5)
	s.Assert().ElementsMatch(got, s.seedData[5:10])

	// Test list index of empty project
	got, err = s.logWriterStorage.List(ctx, ListInput{
		ProjectID: 10,
		Offset:    0,
		Limit:     10,
	})

	s.Assert().NoError(err)
	s.Assert().Len(got, 0)
}

func (s *LogWriterStorageTestSuite) TestCreateLogWriter() {
	ctx := context.Background()

	got, err := s.logWriterStorage.Create(ctx, &model.LogWriter{
		Base: model.Base{
			ProjectID: 3,
		},
		Name:   "log-writer-10",
		Status: model.StatusDeployed,
		Source: &model.LogWriterSource{
			LogWriterSource: &timberv1.LogWriterSource{
				Type: timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG,
				RouterLogSource: &timberv1.RouterLogSource{
					RouterId:   1,
					RouterName: "router-1",
					Kafka: &timberv1.KafkaConfig{
						Brokers: "my-broker",
						Topic:   "my-topic",
					},
				},
			},
		},
	})

	s.Assert().NoError(err)
	exp, err := s.logWriterStorage.Get(ctx, GetInput{
		ID:        got.ID,
		ProjectID: got.ProjectID,
	})

	s.Assert().NoError(err)
	s.Assert().Equal(exp, got)

	// test conflict
	_, err = s.logWriterStorage.Create(ctx, &model.LogWriter{
		Base: model.Base{
			ProjectID: 2,
		},
		Name:   "log-writer-0",
		Status: model.StatusDeployed,
		Source: &model.LogWriterSource{
			LogWriterSource: &timberv1.LogWriterSource{
				Type: timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG,
				RouterLogSource: &timberv1.RouterLogSource{
					RouterId:   1,
					RouterName: "router-1",
					Kafka: &timberv1.KafkaConfig{
						Brokers: "my-broker",
						Topic:   "my-topic",
					},
				},
			},
		},
	})

	s.Assert().ErrorContains(err, "log_writer exists")
}

func (s *LogWriterStorageTestSuite) TestUpdateLogWriter() {
	ctx := context.Background()

	lw, err := s.logWriterStorage.Create(ctx, &model.LogWriter{
		Base: model.Base{
			ProjectID: 3,
		},
		Name:   "my-log-writer",
		Status: model.StatusDeployed,
		Source: &model.LogWriterSource{
			LogWriterSource: &timberv1.LogWriterSource{
				Type: timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG,
				RouterLogSource: &timberv1.RouterLogSource{
					RouterId:   1,
					RouterName: "router-1",
					Kafka: &timberv1.KafkaConfig{
						Brokers: "my-broker",
						Topic:   "my-topic",
					},
				},
			},
		},
	})

	lw.Status = model.StatusUninstalled
	lw.Error = "updated error message"
	lw.Source.RouterLogSource = &timberv1.RouterLogSource{
		RouterId:   1,
		RouterName: "router-1",
		Kafka: &timberv1.KafkaConfig{
			Brokers: "my-broker-modified",
			Topic:   "my-topic-modified",
		},
	}

	s.Assert().NoError(err)
	got, err := s.logWriterStorage.Update(ctx, lw)

	s.Assert().NoError(err)
	s.Assert().Equal(lw.Status, got.Status)
	s.Assert().Equal(lw.Source, got.Source)
	s.Assert().Equal(lw.Error, got.Error)
}
