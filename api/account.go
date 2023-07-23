package api

import (
	"api/api/args"
	"api/config"
	"api/logger"
	"api/service"

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
// @Summary      Register a new account
// @Description  Register a new account with username and password
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        account body args.AccountRegisterArgs true "username and password"
// @Success      200 "Account created, return account information"
// @Failure      400 "Username or password is missing"
// @Router       /register [post]
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
