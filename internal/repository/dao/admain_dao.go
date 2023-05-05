package dao

import (
	"blogProject/model"
	"fmt"
	"gorm.io/gorm"
)

var AdminDao adminDao

type adminDao struct {}

func (adminDao) AdminRegister(db *gorm.DB,admin *model.Admin) (rowsAffected int64, err error) {
	tx := db.Create(&admin)
	if tx.Error != nil {
		return 0, err
	}
	return tx.RowsAffected, nil
}

func (adminDao) AdminLogin(db *gorm.DB,username string) (*model.Admin, error) {
	var u model.Admin
	tx := db.Table("admin").Where("username=?", username).First(&u)
	if tx.Error != nil {
		fmt.Println("查询全部学生信息失败:", tx.Error)
		return nil, tx.Error
	}

	return &u, nil
}
