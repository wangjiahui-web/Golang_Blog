package service

import (
	"blogProject/internal/repository/dao"
	"blogProject/pkg/config"
	"blogProject/pkg/util"
	"fmt"
	"go.uber.org/zap"
	"mime/multipart"
	"path/filepath"
	"time"
)

var AvatarService avatarService

type avatarService struct {}

func (avatarService)UploadAvatar(file *multipart.FileHeader,id *int ,table *string) (path string, err error) {
	filename := fmt.Sprintf("%s%s", time.Now().Format(config.DigitDateTimeLayout), filepath.Ext(file.Filename))
	fmt.Println(filename)
	// 头像保存路径
	avatarPath := config.UploadsPath + filename
	fmt.Println(avatarPath)
	// 保存文件
	err = util.SaveUploadedFile(file, avatarPath)
	if err != nil {
		fmt.Printf("保存头像失败: %s\n", err)
		zap.S().Errorf("保存头像失败: %s\n", err)
		return "", err
	}
	_ , err1 := dao.AvatarDao.InsertAvatar(dao.DB,*id,*table,avatarPath)
	if err1 != nil {
		return
	}
	return avatarPath ,nil
}
