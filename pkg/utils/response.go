package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// JSON 成功响应
func JSONSuccess(c *app.RequestContext, data interface{}, msg string) {
	c.JSON(consts.StatusOK, Response{
		Status: 200,
		Data:   data,
		Msg:    msg,
	})
}

// JSON 错误响应 (业务错误)
func JSONError(c *app.RequestContext, status int, err error) {
	errString := ""
	if err != nil {
		errString = err.Error()
	}
	c.JSON(consts.StatusOK, Response{
		Status: status,
		Data:   nil,
		Msg:    "Error",
		Error:  errString,
	})
}
