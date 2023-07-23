package store

import "api/config"

type DB struct {
}

func NewDB(c *config.Config) (*DB, error) {
	return &DB{}, nil
}
