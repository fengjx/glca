package errno

import "github.com/fengjx/luchen"

// common

var (
	SystemErr = &luchen.Errno{Code: 500, Msg: "系统错误"}
	ArgsErr   = &luchen.Errno{Code: 400, Msg: "参数错误"}
)

// user

var (
	PasswordErr     = &luchen.Errno{Code: 10000, Msg: "密码错误"}
	UserNotExistErr = &luchen.Errno{Code: 10001, Msg: "用户不存在"}
)
