package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"architechture/pkg/errcode"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

func Fail(c *gin.Context, err errcode.Error) {
	c.JSON(http.StatusOK, &Response{
		Code:    err.Code(),
		Message: err.Error(),
		Data:    nil,
	})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
		Data:    nil,
	})
}
