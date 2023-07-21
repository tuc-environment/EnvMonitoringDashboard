package app

import (
	"api/api"
	"api/config"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func NewEngine(
	c *config.Config,
	helloAPI *api.HelloAPI,
) (*gin.Engine, error) {
	// gin.SetMode(gin.ReleaseMode)
	e := gin.Default()
	e.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Authorization", "Origin", "User-Agent", "X-API-Key"},
		ExposeHeaders:    []string{"Content-Type", "Content-Length", "Authorization", "Origin", "User-Agent", "X-API-Key"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	e.GET("/", helloAPI.Ping)

	return e, nil
}
