package model

import (
	"blogProject/pkg/ctime"
)

const TableNameProjectExperience = "project_experience"

// ProjectExperience mapped from table <project_experience>
type ProjectExperience struct {
	ID                 int32      `gorm:"column:id;type:int(11);primaryKey;comment:项目 ID，自增 ID" json:"id"`
	AdminID            *int64     `gorm:"column:admin_id;type:bigint(20);comment:管理员ID" json:"admin_id"`
	ProjectName        *string    `gorm:"column:project_name;type:varchar(255);comment:项目名称" json:"project_name"`
	ProjectDescription *string    `gorm:"column:project_description;type:varchar(255);comment:项目描述" json:"project_description"`
	ProjectDate        *ctime.CTime `gorm:"column:project_date;type:datetime;comment:项目日期" json:"project_date"`
}

// TableName ProjectExperience's table name
func (*ProjectExperience) TableName() string {
	return TableNameProjectExperience
}
