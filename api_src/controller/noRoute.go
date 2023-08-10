package controller

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type NoRouteAPI struct {
}

func NewNoRouteAPI() *NoRouteAPI {
	return &NoRouteAPI{}
}

func (api *NoRouteAPI) ServeWebapp(webappFS fs.FS) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") || webappFS == nil {
			g := WrapContext(c)
			g.NotFound(nil)
			return
		}

		fp := strings.TrimPrefix(c.Request.URL.Path, "/")
		req := c.Request.Clone(c.Request.Context())
		_, err := webappFS.Open(fp)
		if err != nil {
			req.URL.Path = "/"
		}
		fs := http.FileServer(http.FS(webappFS))
		fs.ServeHTTP(c.Writer, req)
	}
}
