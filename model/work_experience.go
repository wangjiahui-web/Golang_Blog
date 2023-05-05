package model

import (
	"blogProject/pkg/ctime"
)

const TableNameWorkExperience = "work_experience"

// WorkExperience mapped from table <work_experience>
type WorkExperience struct {
	ID          int64      `gorm:"column:id;type:bigint(20);primaryKey" json:"id"`
	AdminID     *int64     `gorm:"column:admin_id;type:bigint(20);comment:管理员 ID" json:"admin_id"`
	ComanyName  *string    `gorm:"column:comany_name;type:varchar(255);comment:所在公司名称" json:"comany_name"`
	Post        *string    `gorm:"column:post;type:varchar(255);comment:岗位" json:"post"`
	WorkContent *string    `gorm:"column:work_content;type:varchar(2048);comment:工作内容" json:"work_content"`
	StartDate   *ctime.CTime `gorm:"column:start_date;type:datetime;comment:在该公司的入职时间" json:"start_date"`
	EndDate     *string    `gorm:"column:end_date;type:varchar(255);comment:在该公司的离职时间，便于使用“至今”" json:"end_date"`
}

// TableName WorkExperience's table name
func (*WorkExperience) TableName() string {
	return TableNameWorkExperience
}
