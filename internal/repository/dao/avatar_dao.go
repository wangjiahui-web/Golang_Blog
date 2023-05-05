package dao

import (
	"fmt"
	"gorm.io/gorm"
)

var AvatarDao avatarDao

type avatarDao struct {}



func (avatarDao) InsertAvatar(db *gorm.DB, id int ,table string ,path string) (rowsAffected int64, err error) {
	fmt.Println(id)
	tx := db.Table(table).Where("id = ?" ,id).Update("avatar", path)

	if tx.Error != nil {
		return 0, err
	}
	return tx.RowsAffected, nil
}

