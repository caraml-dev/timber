package migration

import (
	"github.com/caraml-dev/timber/dataset-service/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var Migrations = []*gormigrate.Migration{
	{
		ID: "2023-03-20-log_writers",
		Migrate: func(db *gorm.DB) error {
			return db.AutoMigrate(&model.LogWriter{})
		},
		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable("log_writers")
		},
	},
}
