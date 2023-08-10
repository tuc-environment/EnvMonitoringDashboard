package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

type ResponseBody struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Error   *string     `json:"error,omitempty"`
	Payload interface{} `json:"payload"`
}

func WrapContext(c *gin.Context) *Context {
	return &Context{c}
}

func (c *Context) makeResponse(code int, err error, payload interface{}) ResponseBody {
	var errMsg *string = nil
	if err != nil {
		msg := err.Error()
		errMsg = &msg
	}
	return ResponseBody{code, http.StatusText(code), errMsg, payload}
}

func (c *Context) InternalServerError(err error) {
	c.JSON(http.StatusInternalServerError, c.makeResponse(http.StatusInternalServerError, err, nil))
}

func (c *Context) NotFound(err error) {
	c.JSON(http.StatusNotFound, c.makeResponse(http.StatusNotFound, err, nil))
}

func (c *Context) Unauthorized(err error) {
	c.JSON(http.StatusUnauthorized, c.makeResponse(http.StatusUnauthorized, err, nil))
}

func (c *Context) BadRequest(err error) {
	c.JSON(http.StatusBadRequest, c.makeResponse(http.StatusBadRequest, err, nil))
}

func (c *Context) OK(payload interface{}) {
	c.JSON(http.StatusOK, c.makeResponse(http.StatusOK, nil, payload))
}
