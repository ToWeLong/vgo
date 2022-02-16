package errors

import (
	"fmt"
	"net/http"
)

// Error 该结构体用于定义错误的返回
// code的设计采用http status code 进行设计
// 为什么：1. 繁多的业务code会导致错误难以维护。
//        2. 前端只需要通过判断http status code就能判断接口是否异常。
//        3. 可通过具体的错误信息来区分具体的错误。
type Error struct {
	// Code http状态码
	Code int `json:"code"`
	// Message 错误提示消息
	Message interface{} `json:"message"`
}

func NewError(code int, message interface{}) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e *Error) Error() string {
	switch m := e.Message.(type) {
	case string:
		return m
	case map[string]string:
		var msg string
		for k, v := range m {
			msg += fmt.Sprintf("%s : %s", k, v)
		}
		return msg
	default:
		return ""
	}
}

var (
	Unknown   = NewError(http.StatusInternalServerError, "服务器未知异常")
	ParamsErr = NewError(http.StatusBadRequest, "参数校验异常")
	NotFound  = NewError(http.StatusNotFound, "资源未找到")
)
