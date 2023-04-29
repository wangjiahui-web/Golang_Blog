package model

import (
	"blogProject/pkg/ctime"
)

type User struct {
	UserID        int         `json:"user_id" gorm:"column:user_id"`                 // 用户 ID
	Username      string      `json:"username" gorm:"column:username"`               // 用户名
	Password      string      `json:"password" gorm:"column:password"`               // 密码
	Tel           string      `json:"tel" gorm:"column:tel"`                         // 电话号码
	CreateTime    ctime.CTime `json:"create_time" gorm:"column:create_time"`         // 记录创建时间
	IsDelete      int         `json:"is_delete" gorm:"column:is_delete"`             // 删除
}
