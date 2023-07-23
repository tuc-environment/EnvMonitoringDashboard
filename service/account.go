package service

import "api/store"

type AccountService struct {
	db *store.DBClient
}

func NewAccountService(db *store.DBClient) *AccountService {
	return &AccountService{db}
}
