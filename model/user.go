package model

import (
	"blogProject/pkg/ctime"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID         int64      `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true;comment:主键" json:"id"`
	Username   string     `gorm:"column:username;type:varchar(20);not null;comment:用户名" json:"username"`
	Password   string     `gorm:"column:password;type:varchar(128);not null;comment:密码" json:"password"`
	Tel        *string    `gorm:"column:tel;type:varchar(30);comment:手机号" json:"tel"`
	Avatar     *string    `gorm:"column:avatar;type:varchar(255);comment:用户头像地址" json:"avatar"`
	CreateTime *ctime.CTime `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
