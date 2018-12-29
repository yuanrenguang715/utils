package errors

import (
	"encoding/json"
)

var CodeInfo = map[int64]string{
	1001: "接口认证失败",
	1002: "授权过期",
	1003: "未登陆,token为空",
	1004: "登陆次数被限制",
	1005: "获取token失败",
	1006: "操作过于频繁",
	1007: "用户名或者密码不对",
	1009: "账号被禁用",
	1008: "ip受限",
	2001: "短信发送失败,请重试",
	2002: "此号码不存在",
	2003: "发送失败，同一个手机号码，每天只能发送10条短信",
	2004: "手机号码不合法",
	2005: "此号码已被注册",
	2006: "密码格式不正确",
	2007: "验证失败,请重试",
	2008: "验证码不正确,请重试",
	2009: "参数缺失",
	2010: "此手机未被注册",
	2011: "此账号被封禁,请联系管理员",
	2012: "密码错误,请重试",
	2013: "登录失败",
	2014: "登出失败",
	2015: "修改密码失败",
	2016: "新密码错误",
	2017: "旧密码错误",
	2018: "重置密码失败",
	2019: "暂无记录",
	2020: "修改失败",
	2021: "支付密码错误,请重新输入",
	2022: "邮箱地址格式不正确",
	2023: "密码错误",
	2024: "获取失败,请重试",
	2025: "添加失败",
}

// Error ...
type Error struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// New 常规业务逻辑错误
func New(code int, detail string) error {
	return &Error{
		Code:   code,
		Detail: detail,
	}
}

// Unauthorized 未经授权错误
func Unauthorized(detail string) error {
	return &Error{
		Code:   401,
		Detail: detail,
	}
}

// Internal 内部错误
func Internal(err error) error {
	return &Error{
		Code:   500,
		Detail: err.Error(),
	}
}

//returns error code on info
func QueryReturnCode(code int64, info interface{}) map[string]interface{} {
	if code == 0 {
		return map[string]interface{}{
			"status": code,
			"info":   info,
		}
	}
	if CodeInfo[code] == "" {
		return map[string]interface{}{
			"status": code,
			"info":   "未知错误代码",
		}
	}

	return map[string]interface{}{
		"status": code,
		"info":   CodeInfo[code],
	}
}
