package storage

import (
	"context"
	"fmt"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/model"
	"github.com/stretchr/testify/suite"
)

type LogWriterStorageTestSuite struct {
	suite.Suite
	logWriterStorage logWriter
	seedData         []model.LogWriter
}

func (s *LogWriterStorageTestSuite) SetupSuite() {
	s.Suite.T().Log("Setting up LogWriterStorageTestSuite")

	s.logWriterStorage = logWriter{db: testDB}

	var seedData []model.LogWriter
	for i := 0; i < 100; i++ {
		seedData = append(seedData, model.LogWriter{
			Base: model.Base{
				ProjectID: 2,
			},
			Name:   fmt.Sprintf("log-writer-%d", i),
			Status: model.StatusDeployed,
			LogWriterSource: &model.LogWriterSource{
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
	}

	for _, seed := range seedData {
		c, err := s.logWriterStorage.Create(context.Background(), seed)
		s.Assert().NoError(err)
		s.seedData = append(s.seedData, c)
	}
}

func (s *LogWriterStorageTestSuite) TestGet() {
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
		Name: "log-writer-1",
	})

	s.Assert().NoError(err)
	s.Assert().Equal(s.seedData[0], got)
}

func (s *LogWriterStorageTestSuite) TestList() {
	ctx := context.Background()

	// Test list first 5
	got, err := s.logWriterStorage.List(ctx, ListInput{
		Offset: 0,
		Limit:  5,
	})

	s.Assert().NoError(err)
	s.Assert().Len(got, 5)
	s.Assert().ElementsMatch(got, s.seedData[:5])

	// Test list index 5 to 10
	got, err = s.logWriterStorage.List(ctx, ListInput{
		Offset: 5,
		Limit:  5,
	})

	s.Assert().NoError(err)
	s.Assert().Len(got, 5)
	s.Assert().ElementsMatch(got, s.seedData[5:10])
}

func (s *LogWriterStorageTestSuite) TestCreate() {
	ctx := context.Background()

	got, err := s.logWriterStorage.Create(ctx, model.LogWriter{
		Base: model.Base{
			ProjectID: 3,
		},
		Name:   fmt.Sprintf("log-writer-10"),
		Status: model.StatusDeployed,
		LogWriterSource: &model.LogWriterSource{
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
	s.Assert().Equal(got, exp)
}

func (s *LogWriterStorageTestSuite) TestUpdate() {
	ctx := context.Background()

	lw, err := s.logWriterStorage.Create(ctx, model.LogWriter{
		Base: model.Base{
			ProjectID: 3,
		},
		Name:   fmt.Sprintf("my-log-writer"),
		Status: model.StatusDeployed,
		LogWriterSource: &model.LogWriterSource{
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
	lw.LogWriterSource.RouterLogSource = &timberv1.RouterLogSource{
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
	s.Assert().Equal(got.Status, lw.Status)
	s.Assert().Equal(got.LogWriterSource, lw.LogWriterSource)
}
