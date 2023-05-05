package model

import (
	"blogProject/pkg/ctime"
)

const TableNameEducationExperience = "education_experience"

// EducationExperience mapped from table <education_experience>
type EducationExperience struct {
	ID         int64      `gorm:"column:id;type:bigint(20);primaryKey;comment:教育经历 ID" json:"id"`
	AdminID    *int64     `gorm:"column:admin_id;type:bigint(20);comment:管理员 ID" json:"admin_id"`
	SchoolName *string    `gorm:"column:school_name;type:varchar(255);comment:学校名称" json:"school_name"`
	StartDate  *ctime.CTime `gorm:"column:start_date;type:datetime;comment:入学日期，日期形式，便于排序" json:"start_date"`
	Major      *string    `gorm:"column:major;type:varchar(255);comment:在校专业" json:"major"`
	EndDate    *string    `gorm:"column:end_date;type:varchar(20);comment:毕业日期，字符串形式，容易处理 "至今"" json:"end_date"`
	Courses    *string    `gorm:"column:courses;type:varchar(2048);comment:在校期间学习的专业课程" json:"courses"`
}

// TableName EducationExperience's table name
func (*EducationExperience) TableName() string {
	return TableNameEducationExperience
}