package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 标准返回格式
type Response struct {
	Code int         `json:"code"` // 自定义业务码
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// Success 成功返回的包装函数
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Data: data,
		Msg:  "success",
	})
}

// Fail 失败返回
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{ // 即使业务失败，HTTP 状态码通常也给 200，由内部 code 区分
		Code: 1,
		Data: nil,
		Msg:  msg,
	})
}
