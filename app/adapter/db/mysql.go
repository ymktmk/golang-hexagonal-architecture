package db

import (
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {

	var dsn = os.Getenv("MYSQL_DSN")

	exponentialBackoff := backoff.NewExponentialBackOff()
	exponentialBackoff.MaxElapsedTime = time.Second * 50

	err := backoff.Retry(func() error {
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		return err
	}, exponentialBackoff)

	sqlCon, err := db.DB()
	if err != nil {
		return nil, err
	}

	return db, sqlCon.Ping()
}
