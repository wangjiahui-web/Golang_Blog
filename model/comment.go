package model

import (
	"blogProject/pkg/ctime"
)

const TableNameComment = "comment"

// Comment mapped from table <comment>
type Comment struct {
	ID         int64      `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true;comment:评论ID，自增ID" json:"id"`
	Content    *string    `gorm:"column:content;type:text;comment:评论内容" json:"content"`
	CreateTime *ctime.CTime `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;comment:评论时间" json:"create_time"`
	UserID     *int64     `gorm:"column:user_id;type:bigint(20);comment:评论人ID" json:"user_id"`
	ToUserID   *int64     `gorm:"column:to_user_id;type:bigint(20);comment:评论人向谁回复" json:"to_user_id"`
	BlogID     *int64     `gorm:"column:blog_id;type:bigint(20);comment:被评论的博客ID" json:"blog_id"`
	ParentID   *int64     `gorm:"column:parent_id;type:bigint(20);comment:父节点 id" json:"parent_id"`
	Level      *int32     `gorm:"column:level;type:int(11);comment:层级" json:"level"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
