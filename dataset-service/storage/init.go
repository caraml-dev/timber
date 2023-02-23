package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"path"
	"runtime"
	"time"

	"github.com/golang-migrate/migrate/v4"
	migratePg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/caraml-dev/timber/common/log"
	"github.com/caraml-dev/timber/dataset-service/config"
)

// InitDB initialize database instance based on application configuration
func InitDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	databaseURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Database,
		cfg.Password)

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		// https://github.com/go-gorm/gorm/issues/4834
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("UTC")
			return time.Now().Round(time.Second).In(ti)
		},
	})

	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting the underlying database connection: %w", err)
	}

	// Update connection-related configuration
	sqlDB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	// migrateDB
	err = migrateDB(sqlDB, cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("error migrating database: %w", err)
	}

	return db, nil
}

func migrateDB(sqlDB *sql.DB, dbName string) error {
	log.Infof("running database migration")

	driver, err := migratePg.WithInstance(sqlDB, &migratePg.Config{})
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithDatabaseInstance(
		getFileURL("../storage/migration"),
		dbName,
		driver)
	if err != nil {
		return err
	}

	err = migration.Up()
	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	return err
}

func getFileURL(filePath string) string {
	_, filename, _, _ := runtime.Caller(0)
	return fmt.Sprintf("file://%s/%s", path.Dir(filename), filePath)
}
