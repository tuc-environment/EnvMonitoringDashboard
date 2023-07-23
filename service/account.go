package service

import "api/store"

type AccountService struct {
	db *store.DB
}

func NewAccountService(db *store.DB) *AccountService {
	return &AccountService{db}
}
