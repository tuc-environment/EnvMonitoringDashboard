package controller

import (
	"api/api_src/config"
	"api/api_src/controller/args"
	"api/api_src/logger"
	"api/api_src/service"

	"github.com/gin-gonic/gin"
)

type AccountAPI struct {
	config         *config.Config
	logger         *logger.Logger
	accountService *service.AccountService
}

func NewAccountAPI(c *config.Config, l *logger.Logger, s *service.AccountService) *AccountAPI {
	return &AccountAPI{c, l, s}
}

// Register godoc
//
//	@Summary		Register a new account
//	@Description	Register a new account with username and password
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			account	body	args.AccountRegisterArgs	true	"username and password"
//	@Success		200		"Account created, return account information"
//	@Failure		400		"Username or password is missing"
//	@Router			/register [post]
func (api *AccountAPI) Register(g *gin.Context) {
	c := WrapContext(g)

	body := args.AccountRegisterArgs{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.BadRequest(err)
		return
	}

	account, err := api.accountService.CreateAccount(body.Username, body.Password)
	if err != nil {
		c.BadRequest(err)
		return
	}

	c.OK(account)
}

// Login godoc
//
//	@Summary		Login to an account
//	@Description	Login to an account with username and password
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			account	body	args.AccountLoginArgs	true	"username and password"
//	@Success		200		"Login success, return account information"
//	@Failure		400		"Username or password is missing"
//	@Failure		401		"Username or password is incorrect"
//	@Router			/login [post]
func (api *AccountAPI) Login(g *gin.Context) {
	c := WrapContext(g)

	body := args.AccountLoginArgs{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.BadRequest(err)
		return
	}

	account, err := api.accountService.GetAccount(body.Username, body.Password)
	if err != nil {
		c.BadRequest(err)
		return
	}

	c.OK(account)
}

// Account Infomation godoc
//
//	@Summary		Get account information
//	@Description	Get account information with token
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	"Return account information"
//	@Failure		401	"Token is incorrect"
//	@Router			/ [get]
func (api *AccountAPI) GetAccount(g *gin.Context) {
	c := WrapContext(g)

	apiToken := c.GetHeader("Authorization")
	account, err := api.accountService.GetAccountWithToken(apiToken)
	if err != nil {
		c.Unauthorized(err)
		return
	}

	c.OK(account)
}

func (api *AccountAPI) AuthMiddleware(g *gin.Context) {
	c := WrapContext(g)

	apiToken := c.GetHeader("Authorization")
	_, err := api.accountService.GetAccountWithToken(apiToken)
	if err != nil {
		c.Unauthorized(err)
		return
	}

	c.Next()
}
