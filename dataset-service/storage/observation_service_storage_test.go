package storage

import (
	"context"
	"fmt"

	"github.com/stretchr/testify/suite"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/model"
)

type ObservationServiceStorageTestSuite struct {
	suite.Suite
	ObservationServiceStorage observationService
	seedData                  []*model.ObservationService
}

func (s *ObservationServiceStorageTestSuite) SetupSuite() {
	s.Suite.T().Log("Setting up ObservationServiceStorageTestSuite")

	s.ObservationServiceStorage = observationService{db: testDB}

	for i := 0; i < 100; i++ {
		c, err := s.ObservationServiceStorage.Create(context.Background(), &model.ObservationService{
			Base: model.Base{
				ProjectID: 2,
			},
			Name:   fmt.Sprintf("observation-service-%d", i),
			Status: model.StatusDeployed,
			Error:  fmt.Sprintf("observation-service-%d", i),
			Source: &model.ObservationServiceSource{
				ObservationServiceSource: &timberv1.ObservationServiceSource{
					Type: timberv1.ObservationServiceSourceType_OBSERVATION_SERVICE_SOURCE_TYPE_KAFKA,
					Kafka: &timberv1.KafkaConfig{
						Brokers: "my-broker",
						Topic:   "my-topic",
					},
				},
			},
		})

		s.Assert().NoError(err)
		s.seedData = append(s.seedData, c)
	}
}

func (s *ObservationServiceStorageTestSuite) TestGetObservationService() {
	ctx := context.Background()

	// Test get with id
	got, err := s.ObservationServiceStorage.Get(ctx, GetInput{
		ID:        1,
		ProjectID: 2,
	})

	s.Assert().NoError(err)
	s.Assert().Equal(s.seedData[0], got)

	// Test get with name
	got, err = s.ObservationServiceStorage.Get(ctx, GetInput{
		Name: "observation-service-99",
	})

	s.Assert().NoError(err)
	s.Assert().Equal(s.seedData[99], got)
	s.Assert().Equal(int64(100), got.ID)
}

func (s *ObservationServiceStorageTestSuite) TestListObservationService() {
	ctx := context.Background()

	// Test list first 5
	got, err := s.ObservationServiceStorage.List(ctx, ListInput{
		ProjectID: 2,
		Offset:    0,
		Limit:     5,
	})

	s.Assert().NoError(err)
	s.Assert().Len(got, 5)
	s.Assert().ElementsMatch(got, s.seedData[:5])

	// Test list index 5 to 10
	got, err = s.ObservationServiceStorage.List(ctx, ListInput{
		ProjectID: 2,
		Offset:    5,
		Limit:     5,
	})

	s.Assert().NoError(err)
	s.Assert().Len(got, 5)
	s.Assert().ElementsMatch(got, s.seedData[5:10])

	// Test list index of empty project
	got, err = s.ObservationServiceStorage.List(ctx, ListInput{
		ProjectID: 10,
		Offset:    0,
		Limit:     10,
	})

	s.Assert().NoError(err)
	s.Assert().Len(got, 0)
}

func (s *ObservationServiceStorageTestSuite) TestCreateObservationService() {
	ctx := context.Background()

	got, err := s.ObservationServiceStorage.Create(ctx, &model.ObservationService{
		Base: model.Base{
			ProjectID: 3,
		},
		Name:   "observation-service-10",
		Status: model.StatusDeployed,
		Source: &model.ObservationServiceSource{
			ObservationServiceSource: &timberv1.ObservationServiceSource{
				Type: timberv1.ObservationServiceSourceType_OBSERVATION_SERVICE_SOURCE_TYPE_KAFKA,
				Kafka: &timberv1.KafkaConfig{
					Brokers: "my-broker",
					Topic:   "my-topic",
				},
			},
		},
	})

	s.Assert().NoError(err)
	exp, err := s.ObservationServiceStorage.Get(ctx, GetInput{
		ID:        got.ID,
		ProjectID: got.ProjectID,
	})

	s.Assert().NoError(err)
	s.Assert().Equal(exp, got)

	// test conflict
	_, err = s.ObservationServiceStorage.Create(ctx, &model.ObservationService{
		Base: model.Base{
			ProjectID: 2,
		},
		Name:   "observation-service-0",
		Status: model.StatusDeployed,
		Source: &model.ObservationServiceSource{
			ObservationServiceSource: &timberv1.ObservationServiceSource{
				Type: timberv1.ObservationServiceSourceType_OBSERVATION_SERVICE_SOURCE_TYPE_KAFKA,
				Kafka: &timberv1.KafkaConfig{
					Brokers: "my-broker",
					Topic:   "my-topic",
				},
			},
		},
	})

	s.Assert().ErrorContains(err, "observation_service exists")
}

func (s *ObservationServiceStorageTestSuite) TestUpdateObservationService() {
	ctx := context.Background()

	lw, err := s.ObservationServiceStorage.Create(ctx, &model.ObservationService{
		Base: model.Base{
			ProjectID: 3,
		},
		Name:   "my-observation-service",
		Status: model.StatusDeployed,
		Source: &model.ObservationServiceSource{
			ObservationServiceSource: &timberv1.ObservationServiceSource{
				Type: timberv1.ObservationServiceSourceType_OBSERVATION_SERVICE_SOURCE_TYPE_KAFKA,
				Kafka: &timberv1.KafkaConfig{
					Brokers: "my-broker",
					Topic:   "my-topic",
				},
			},
		},
	})

	lw.Status = model.StatusUninstalled
	lw.Error = "updated error message"
	lw.Source.Kafka = &timberv1.KafkaConfig{
		Brokers: "my-broker-modified",
		Topic:   "my-topic-modified",
	}

	s.Assert().NoError(err)
	got, err := s.ObservationServiceStorage.Update(ctx, lw)

	s.Assert().NoError(err)
	s.Assert().Equal(lw.Status, got.Status)
	s.Assert().Equal(lw.Source, got.Source)
	s.Assert().Equal(lw.Error, got.Error)
}
