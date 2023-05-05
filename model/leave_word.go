package model

import (
	"blogProject/pkg/ctime"
)

const TableNameLeaveWord = "leave_word"

// LeaveWord mapped from table <leave_word>
type LeaveWord struct {
	ID         int64      `gorm:"column:id;type:bigint(20);primaryKey" json:"id"`
	Name       *string    `gorm:"column:name;type:varchar(255);comment:姓名" json:"name"`
	Email      *string    `gorm:"column:email;type:varchar(255);comment:邮箱" json:"email"`
	Subject    *string    `gorm:"column:subject;type:varchar(255);comment:主题" json:"subject"`
	Content    *string    `gorm:"column:content;type:varchar(255);comment:留言内容" json:"content"`
	CreateTime *ctime.CTime `gorm:"column:create_time;type:datetime;comment:留言时间" json:"create_time"`
}

// TableName LeaveWord's table name
func (*LeaveWord) TableName() string {
	return TableNameLeaveWord
}
