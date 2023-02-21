package storage

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/caraml-dev/timber/common/log"
	"github.com/caraml-dev/timber/dataset-service/storage/migration"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	user := "timber"
	password := "timber"
	dbName := "timber"
	code := 0

	// Initialize DB for testing
	req := testcontainers.ContainerRequest{
		Image:        "postgres:12.14-alpine",
		ExposedPorts: []string{"5432"},
		WaitingFor:   wait.ForExposedPort(),
		Env: map[string]string{
			"POSTGRES_USER":     user,
			"POSTGRES_PASSWORD": password,
			"POSTGRES_DB":       dbName,
		},
		SkipReaper: true,
	}

	ctx := context.Background()
	pgC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		log.Errorf("Unable to start postgresql container : %v", err)
	}

	defer func() {
		log.Infof("Tear down")

		_ = pgC.Terminate(ctx)
		os.Exit(code)
	}()

	ep, err := pgC.Endpoint(ctx, "")
	if err != nil {
		log.Errorf("Unable to get endpoint : %v", err)
	}

	host, port, err := net.SplitHostPort(ep)
	if err != nil {
		log.Errorf("Unable to get endpoint : %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("Unable to connect to postgresql : %v", err)
	}

	testDB = db

	mig := gormigrate.New(db, gormigrate.DefaultOptions, migration.Migrations)
	if err = mig.Migrate(); err != nil {
		log.Errorf("Could not migrate: %v", err)
	}
	log.Infof("Migration ran successfully")

	code = m.Run()
}

func TestStorage(t *testing.T) {
	suite.Run(t, new(LogWriterStorageTestSuite))
}
