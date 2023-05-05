package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

//
// SaveUploadedFile 将前端上传的文件保存到 dst 参数指定目录
//	@parameter	file 前端上传的文件对象
//	@parameter	dst 目标文件, 包含路径和目标文件名
//

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		fmt.Println("SaveUploadedFile代码1",err)
		return err
	}
	defer src.Close()

	//这里的代码有问题需要解决
	out, err := os.Create(dst)
	if err != nil {
		fmt.Println("SaveUploadedFile代码2",err)
		return err
	}
	defer func() {
		if closer, ok := interface{}(out).(io.Closer); ok {
			closer.Close()
		}
	}()

	_, err = io.Copy(out, src)
	return err
}
