package db

import (
	"time"

	"github.com/okatu-loli/TikTokLite/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

func Init() {
	var err error
	gormlogrus := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)
	DB, err = gorm.Open(mysql.Open(config.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:                              true,
			Logger:                                   gormlogrus,
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
	if err != nil {
		//panic(err)
	}

	if err := DB.Use(tracing.NewPlugin()); err != nil {
		//panic(err)
	}
}
