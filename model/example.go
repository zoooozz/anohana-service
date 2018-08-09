package model

import (
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	RetErr = &Ecode{
		Code:    10000,
		Message: "请求信息失效",
	}
	RetLoginErr = &Ecode{
		Code:    10001,
		Message: "请重新登录",
	}
	ParamsPhoneErr = &Ecode{
		Code:    10002,
		Message: "手机号码错误",
	}
	ParamsPassErr = &Ecode{
		Code:    10003,
		Message: "密码错误",
	}
	RetLoginInfoErr = &Ecode{
		Code:    10004,
		Message: "登录信息异常",
	}
)

var (
	PassSalt = "anohana"
	JwtSalt  = "anohana"
)

type jwtCustomClaims struct {
	Phone string `json:"phone"`
	jwt.StandardClaims
}
