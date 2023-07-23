package api

import (
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
// @Router       /register [post]
func (api *AccountAPI) Register(g *gin.Context) {
	c := WrapContext(g)

	body := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
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

func (api *AccountAPI) Ping(g *gin.Context) {
	c := WrapContext(g)
	c.OK("pong")
}
