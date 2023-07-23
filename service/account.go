package service

import (
	"api/logger"
	"api/store"
	"fmt"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string
}

type AccountService struct {
	db     *store.DBClient
	logger *logger.Logger
}

func NewAccountService(db *store.DBClient, logger *logger.Logger) *AccountService {
	log := logger.Sugar()
	defer log.Sync()

	db.AutoMigrate(&Account{})
	log.Info("database table 'accounts' migrated")
	return &AccountService{db, logger}
}

func (s *AccountService) HashedPassword(password string) string {
	return password
}

func (s *AccountService) CreateAccount(username, password string) (Account, error) {
	log := s.logger.Sugar()
	defer log.Sync()

	var count int64 = 0
	s.db.Model(&Account{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		log.Warnf("username %s already exists", username)
		return Account{}, fmt.Errorf("username %s already exists", username)
	}

	account := Account{
		Username: username,
		Password: s.HashedPassword(password),
	}
	err := s.db.Create(&account).Error
	if err != nil {
		log.Error(err)
		return Account{}, err
	}
	log.Infof("account %s created", username)
	return account, nil
}
