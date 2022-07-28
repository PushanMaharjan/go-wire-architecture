package db

import (
	"fmt"
	"go-wire-architecture/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, dbErr := gorm.Open(mysql.Open(url), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	return db, dbErr
}
