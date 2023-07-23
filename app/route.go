package app

import (
	"api/api"
	"api/config"
	"api/logger"
	"time"

	"github.com/gin-contrib/cors"

	_ "api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewEngine(
	c *config.Config,
	log *logger.Logger,
	accountAPI *api.AccountAPI,
	dataAPI *api.DataAPI,
) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Use(NewLoggerMiddleware(c, log), gin.Recovery())

	e.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Authorization", "Origin", "User-Agent", "X-API-Key"},
		ExposeHeaders:    []string{"Content-Type", "Content-Length", "Authorization", "Origin", "User-Agent", "X-API-Key"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	e.POST("/register", accountAPI.Register)
	e.POST("/login", accountAPI.Login)
	e.GET("/", accountAPI.AuthMiddleware, accountAPI.GetAccount)

	e.Group("/api", accountAPI.AuthMiddleware).
		GET("/", dataAPI.Ping)

	// append docs
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return e, nil
}
