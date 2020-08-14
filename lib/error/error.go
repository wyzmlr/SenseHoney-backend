package error

import (
	"fmt"
)

const (
	SuccessCode                  = 200
	SuccessMsg                   = "success"
	ErrCode                      = 100
	ErrMsg                       = "error"
	ErrJsonParseCode             = 10000
	ErrJsonParseMsg              = "JSON解析失败"
	ErrAuthCheckTokenFailCode    = 10001
	ErrAuthCheckTokenFailMsg     = "Token鉴权失败"
	ErrAuthCheckTokenTimeoutCode = 10002
	ErrAuthCheckTokenTimeoutMsg  = "Token已过期"
	ErrAuthGenTokenFailCode      = 10003
	ErrAuthGenTokenFailMsg       = "Token生成失败"
	ErrTokenCode                 = 10004
	ErrTokenMsg                  = "Token错误"
)

func Check(e error, tips string) {
	if e != nil {
		fmt.Println(tips)
	}
}
