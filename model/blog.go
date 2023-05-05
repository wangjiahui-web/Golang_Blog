package model

import (
	"blogProject/pkg/ctime"
)

const TableNameBlog = "blog"

// Blog mapped from table <blog>
type Blog struct {
	ID         int64      `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true;comment:主键" json:"id"`
	Title      string     `gorm:"column:title;type:varchar(255);not null;comment:标题" json:"title"`
	Summary    *string    `gorm:"column:summary;type:varchar(512);comment:博客摘要" json:"summary"`
	Content    *string    `gorm:"column:content;type:mediumtext;comment:内容" json:"content"`
	PostTime   *ctime.CTime `gorm:"column:post_time;type:datetime;default:CURRENT_TIMESTAMP;comment:发表时间" json:"post_time"`
	ReadCount  *int32     `gorm:"column:read_count;type:int(11);comment:阅读量" json:"read_count"`
	AdminID    *int64     `gorm:"column:admin_id;type:bigint(20);index:idx_admin_id,priority:1;comment:博客发布人 ID" json:"admin_id"`
	IsDelete   *int32     `gorm:"column:is_delete;type:int(11);comment:标识是否删除记录" json:"is_delete"`
	CategoryID *int64     `gorm:"column:category_id;type:bigint(20);index:idx_category_id,priority:1;comment:分类外键" json:"category_id"`
	FirstImage *string    `gorm:"column:first_image;type:varchar(255);comment:博客首图" json:"first_image"`
}

// TableName Blog's table name
func (*Blog) TableName() string {
	return TableNameBlog
}
