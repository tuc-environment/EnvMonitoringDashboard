package store

import (
	"EnvMonitoringDashboard/api_src/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DBClient struct {
	*gorm.DB
}

type RawDBClient struct {
	*gorm.DB
}

func NewDBClient(c *config.Config) (*DBClient, error) {
	db, err := createClient(c.DB_DRIVER, c.DB_CONNECTION_STRING)
	if err == nil {
		return &DBClient{DB: db}, nil
	} else {
		return nil, err
	}
}

func NewRawDBClient(c *config.Config) (*RawDBClient, error) {
	db, err := createClient(c.RAW_DB_DRIVER, c.RAW_DB_CONNECTION_STRING)
	if err == nil {
		return &RawDBClient{DB: db}, nil
	} else {
		return nil, err
	}
}

func createClient(driver string, connectionString string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if driver == "sqlite" {
		db, err = gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	} else if driver == "mysql" {
		db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	} else if driver == "postgres" {
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	} else if driver == "mssql" {
		db, err = gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	} else {
		return nil, fmt.Errorf("DB_DRIVER %s not supported", driver)
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}
