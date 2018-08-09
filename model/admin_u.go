package model

import (
	"time"
)

type AdminUModel struct {
	Id       int64     `json:"id"`
	Ctime    time.Time `json:"ctime"`
	Mtime    time.Time `json:"mtime"`
	Username string    `json:"username"` //用户名
	Email    string    `json:"email"`    //邮箱
	Type     int8      `json:"type"`     //用户类型
	Password string    `json:"password"` //密码
	Photo    string    `json:"photo"`    //头像
	Ip       string    `json:"ip"`       //ip
	Phone    string    `json:"phone"`    //手机号
	State    int8      `json:"state"`    //状态
}

var (
	AdminUstateDisable = int8(4) //用户被禁用
	AdminUstateNormal  = int8(2) //正常
	DefaultPhoto       = "http://pd4tbi799.bkt.clouddn.com/photo/default.jpeg"
	PhotoUrl           = "http://pd4tbi799.bkt.clouddn.com"
)
