package store

import (
	"EnvMonitoringDashboard/api_src/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBClient struct {
	*gorm.DB
}

func NewDBClient(c *config.Config) (*DBClient, error) {
	var db *gorm.DB
	var err error
	if c.DB_DRIVER == "sqlite" {
		db, err = gorm.Open(sqlite.Open(c.DB_CONNECTION_STRING), &gorm.Config{})
	} else if c.DB_DRIVER == "mysql" {
		db, err = gorm.Open(mysql.Open(c.DB_CONNECTION_STRING), &gorm.Config{})
	} else if c.DB_DRIVER == "postgres" {
		db, err = gorm.Open(postgres.Open(c.DB_CONNECTION_STRING), &gorm.Config{})
	} else {
		return nil, fmt.Errorf("DB_DRIVER %s not supported", c.DB_DRIVER)
	}
	if err != nil {
		return nil, err
	}
	return &DBClient{DB: db}, nil
}
