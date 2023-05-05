package model

const TableNameCategory = "category"

// Category mapped from table <category>
type Category struct {
	CategoryID   int64  `gorm:"column:category_id;type:bigint(20);primaryKey;autoIncrement:true;comment:分类ID，自增ID" json:"category_id"`
	CategoryName string `gorm:"column:category_name;type:varchar(100);not null;comment:分类名称" json:"category_name"`
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}
