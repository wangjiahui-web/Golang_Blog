package dao

import (
	"blogProject/model"
	"fmt"
	"gorm.io/gorm"
)

var UserDao userDao

type userDao struct {}

func (userDao) UserRegister(db *gorm.DB,user *model.User) (rowsAffected int64, err error) {
	tx := db.Create(&user)
	if tx.Error != nil {
		return 0, err
	}
	return tx.RowsAffected, nil
}

func (userDao) UserLogin(db *gorm.DB,username string) (*model.User, error) {
	var u model.User
	tx := db.Table("User").Where("username=?", username).First(&u)
	if tx.Error != nil {
		fmt.Println("查询全部学生信息失败:", tx.Error)
		return nil, tx.Error
	}
	return &u, nil
}

