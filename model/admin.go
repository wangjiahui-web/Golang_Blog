package model

import (
	"blogProject/pkg/ctime"
)

const TableNameAdmin = "admin"

// Admin mapped from table <admin>
type Admin struct {
	ID       int64      `gorm:"column:id;type:bigint(20);primaryKey" json:"id"`
	Username *string    `gorm:"column:username;type:varchar(255);comment:用户名" json:"username"`
	Nickname *string    `gorm:"column:nickname;type:varchar(255);comment:昵称" json:"nickname"`
	Password *string    `gorm:"column:password;type:varchar(255);comment:密码" json:"password"`
	Website  *string    `gorm:"column:website;type:varchar(255);comment:博主个人首页" json:"website"`
	Birth    *ctime.CTime `gorm:"column:birth;type:datetime;comment:博主生日" json:"birth"`
	Tel      *string    `gorm:"column:tel;type:varchar(255);comment:电话号码" json:"tel"`
	Email    *string    `gorm:"column:email;type:varchar(255);comment:电子邮箱" json:"email"`
	Avatar   *string    `gorm:"column:avatar;type:varchar(255);comment:头像地址" json:"avatar"`
	City     *string    `gorm:"column:city;type:varchar(255);comment:所在城市" json:"city"`
	School   *string    `gorm:"column:school;type:varchar(255);comment:毕业学校" json:"school"`
}

// TableName Admin's table name
func (*Admin) TableName() string {
	return TableNameAdmin
}
