package dto

import "blogProject/pkg/ctime"

// AdminRegisterParam 用户注册的时候需要提供的信息
// 注册的时候适用的参数校验规则:
// Username: 参数必须传递, 用户名字符串最小长度为 1, 最大长度为 32
// Password: 参数必须传递, 密码字符串最小长度为 1, 最大长度为 32
// Tel: 电话号码必须为数字, 长度必须是 11 位
type AdminRegisterParam struct {
	Username string      `json:"username" form:"username" binding:"required,min=1,max=32"` // 用户名
	Nickname string      `json:"nickname" form:"nickname"`                                 // 昵称
	Password string      `json:"password" form:"password" binding:"required"`              //密码
	Website  string      `json:"website" form:"website"`                                   //博主个人首页
	Birth    ctime.CTime `json:"birth" form:"birth"`                                       //生日
	Tel      string      `json:"tel" form:"tel" binding:"alphanum,len=11"`                 // 电话
	Email    string      `json:"email" form:"email"`                                       //邮箱地址
	City     string      `json:"city" form:"city"`                                         //所在城市
	School   string      `json:"school" form:"school"`                                     //毕业院校
}

type AdminLoginParam struct {
	Username string `json:"username" form:"username" binding:"required,min=1,max=32"` // 用户名
	Password string `json:"password" form:"password"`                                 // 密码
}