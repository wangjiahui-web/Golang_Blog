package api

import (
	"blogProject/internal/pkg/resp"
	"blogProject/internal/service"
	"blogProject/model/dto"
	"blogProject/pkg/validate"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var AdminApi adminApi

type adminApi struct{}

// AdminRegister     Post 请求,管理员注册
//
//	@Summary		Post 请求, 发送 Json 数据, 参数在消息体中 管理员注册
//	@Description	Post 请求, 发送 Json 数据, 参数在消息体中 管理员注册
//	@Tags			adminApi
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			username	formData	string			false	"要注册的用户名"	minlength(1)	maxlength(32)
//	@Param			nickname	formData	string			false	"要注册的昵称"	minlength(1)	maxlength(32)
//	@Param			password	formData	string			false	"用户密码"		minlength(1)	maxlength(32)
//	@Param			website		formData	string			false	"要注册的个人主页"	minlength(1)	maxlength(32)
//	@Param			birth		formData	string			false	"要注册的生日"
//	@Param			tel			formData	string			false	"电话号码"	minlength(11)	maxlength(11)
//	@Param			email		formData	string			false	"邮件地址"
//	@Param			city		formData	string			false	"所在城市"
//	@Param			school		formData	string			false	"毕业院校"
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/admin/AdminRegister [post]
func (adminApi) AdminRegister(c *gin.Context) {
	var registerParam dto.AdminRegisterParam
	if err := c.ShouldBind(&registerParam); err != nil {
		translateError := validate.TranslateError(err)
		zap.S().Errorf("绑定参数失败: %s\n", translateError)
		resp.InternalServerError(c, translateError)
		return
	}
	register, err := service.AdminService.AdminRegister(&registerParam)
	if err != nil {
		zap.S().Errorf("注册用户失败: %s\n", err)
		resp.InternalServerError(c, "注册用户失败")
		return
	}
	resp.Ok(c, "注册用户成功", register)
}

// AdminLogin     Post 请求, 管理员登陆
//
//	@Summary		Post 请求, 发送 Json 数据, 参数在消息体中 管理员登陆
//	@Description	Post 请求, 发送 Json 数据, 参数在消息体中 管理员登陆
//	@Tags			adminApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			username	formData	string			false	"登陆的用户名"	minlength(1)	maxlength(32)
//	@Param			password	formData	string			false	"用户密码"		minlength(1)	maxlength(32)
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/admin/AdminLogin [post]
func (adminApi) AdminLogin(c *gin.Context) {
	var loginParam dto.AdminLoginParam
	if err := c.ShouldBind(&loginParam); err != nil {
		translateError := validate.TranslateError(err)
		zap.S().Errorf("绑定参数失败: %s\n", translateError)
		resp.InternalServerError(c, translateError)
		return
	}
	register, err := service.AdminService.AdminLogin(&loginParam)
	if err != nil {
		zap.S().Errorf("用户登陆失败: %s\n", err)
		resp.InternalServerError(c, "用户登陆失败")
		return
	}
	resp.Ok(c, "用户登陆成功", register)
}
