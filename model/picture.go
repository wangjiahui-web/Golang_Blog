package model

import (
	"blogProject/pkg/ctime"
)

const TableNamePicture = "picture"

// Picture mapped from table <picture>
type Picture struct {
	ID         int64      `gorm:"column:id;type:bigint(20);primaryKey;comment:照片 ID" json:"id"`
	Title      *string    `gorm:"column:title;type:varchar(255);comment:照片标题" json:"title"`
	ImagePath  *string    `gorm:"column:image_path;type:varchar(255);comment:照片路径" json:"image_path"`
	Address    *string    `gorm:"column:address;type:varchar(255);comment:照片拍摄地点" json:"address"`
	CreateTime *ctime.CTime `gorm:"column:create_time;type:datetime;comment:照片上传日期" json:"create_time"`
}

// TableName Picture's table name
func (*Picture) TableName() string {
	return TableNamePicture
}
