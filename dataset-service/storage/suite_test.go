package storage

import (
	"context"
	"net"
	"os"
	"strconv"
	"testing"

	"github.com/caraml-dev/timber/common/log"
	"github.com/caraml-dev/timber/dataset-service/config"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"

	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		return
	}

	defer func() {
		log.Infof("Tear down")

		_ = pgC.Terminate(ctx)
		os.Exit(code)
	}()

	ep, err := pgC.Endpoint(ctx, "")
	if err != nil {
		log.Errorf("Unable to get endpoint : %v", err)
		return
	}

	host, portStr, err := net.SplitHostPort(ep)
	if err != nil {
		log.Errorf("Unable to get endpoint : %v", err)
		return
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Errorf("error conversion : %v", err)
		return
	}

	dbConfig := &config.DatabaseConfig{
		Host:     host,
		Port:     port,
		User:     "timber",
		Password: "timber",
		Database: "timber",
	}

	db, err := InitDB(dbConfig)
	if err != nil {
		log.Errorf("error initializing database : %v", err)
		return
	}

	testDB = db
	code = m.Run()
}

func TestStorage(t *testing.T) {
	suite.Run(t, new(LogWriterStorageTestSuite))
}
