package db

type Service struct {
	db *DB
}

func NewService(db *DB) *Service {
	return &Service{db}
}
