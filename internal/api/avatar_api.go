package api

import (
	"blogProject/internal/pkg/resp"
	"blogProject/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

var AvatarApi avatarApi

type avatarApi struct{}


// UploadAvatar     Post 请求, 发送 表单 数据, 上传头像
//
//	@Summary		Post 请求, 发送 表单 数据, 参数在消息体中，上传头像
//	@Description	Post 请求, 发送 表单 数据, 参数在消息体中，上传头像
//	@Tags			avatarApi
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id	    formData  	int			false	"注册头像的id"
//  @Param			table	formData  	string		false	"注册头像的表"
//	@Param			avatar	formData	file			false	"上传头像"
//	@Success		200		{object}	swagger.HttpOk	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/avatar/UploadAvatar [post]
func (avatarApi)UploadAvatar(c *gin.Context)  {
	sid := c.PostForm("id")
	id, _ := strconv.Atoi(sid)
	table := c.PostForm("table")
	avatar, _ := c.FormFile("avatar")

	res, err := service.AvatarService.UploadAvatar(avatar,&id,&table)
	if err != nil {
		zap.S().Errorf("上传头像失败: %s\n", err)
		resp.InternalServerError(c, "上传头像失败")
		return
	}
	resp.Ok(c, "上传头像成功", res)
}

