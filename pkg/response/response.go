package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Data struct {
	Code string      `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func Msg(code string) string {
	return Message[code]
}

func BuildSuccessResponse(c *gin.Context, data interface{}) {
	responseData := &Data{
		Code: SUCCESS,
		Msg:  Msg(SUCCESS),
		Data: data,
	}
	c.JSON(http.StatusOK, responseData)
}

func BuildResponse(c *gin.Context, data interface{}, code string, httpCode int) {
	responseData := &Data{
		Code: code,
		Msg:  Msg(code),
		Data: data,
	}
	c.JSON(httpCode, responseData)
}

func BuildResponseWithMsg(c *gin.Context, data interface{}, msg interface{}, code string, httpCode int) {
	responseData := &Data{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(httpCode, responseData)
}
