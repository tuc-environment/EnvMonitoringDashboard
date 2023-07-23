package service

import (
	"api/config"
	"api/logger"
	"api/store"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Token    string `json:"token" gorm:"unique"`
}

type AccountService struct {
	config *config.Config
	db     *store.DBClient
	logger *logger.Logger
}

func NewAccountService(c *config.Config, db *store.DBClient, logger *logger.Logger) *AccountService {
	log := logger.Sugar()
	defer log.Sync()

	db.AutoMigrate(&Account{})
	log.Infoln("database table 'accounts' migrated")
	return &AccountService{c, db, logger}
}

func (s *AccountService) HashedString(password string) string {
	hash := md5.Sum([]byte(fmt.Sprintf("%s%s", password, s.config.HASH_SALT)))
	return hex.EncodeToString(hash[:])
}

func (s *AccountService) HashedPassword(password string) string {
	return s.HashedString(password)
}

func (s *AccountService) NewToken() string {
	return s.HashedString(fmt.Sprintf("%s%s", time.Now().UTC().Format(time.RFC3339), s.config.HASH_SALT))
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
		Token:    s.NewToken(),
	}
	err := s.db.Create(&account).Error
	if err != nil {
		log.Error(err)
		return Account{}, err
	}
	log.Infoln("account %s created", username)
	return account, nil
}

func (s *AccountService) GetAccount(username, password string) (Account, error) {
	log := s.logger.Sugar()
	defer log.Sync()

	var account Account
	err := s.db.Where("username = ? AND password = ?", username, s.HashedPassword(password)).First(&account).Error
	if err != nil {
		log.Error(err)
		return Account{}, err
	}
	log.Infoln("account %s retrieved", username)
	return account, nil
}

func (s *AccountService) GetAccountWithToken(token string) (Account, error) {
	log := s.logger.Sugar()
	defer log.Sync()

	var account Account
	err := s.db.Where("token = ?", token).First(&account).Error
	if err != nil {
		log.Error(err)
		return Account{}, err
	}
	log.Infoln("account %s retrieved", account.Username)
	return account, nil
}
