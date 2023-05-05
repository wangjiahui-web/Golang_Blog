package api

import (
	"blogProject/internal/pkg/resp"
	"blogProject/internal/service"
	"blogProject/model/dto"
	"blogProject/pkg/validate"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var UserApi userApi

type userApi struct{}

// UserRegister     Post 请求,用户注册
//
//	@Summary		Post 请求, 发送 Json 数据, 参数在消息体中 用户注册
//	@Description	Post 请求, 发送 Json 数据, 参数在消息体中 用户注册
//	@Tags			userApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			username	formData	string			false	"要注册的用户名"	minlength(1)	maxlength(32)
//	@Param			password	formData	string			false	"用户密码"		minlength(1)	maxlength(32)
//	@Param			tel			formData	string			false	"电话号码"		minlength(11)	maxlength(11)
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/user/UserRegister [post]
func (userApi) UserRegister(c *gin.Context) {
	var registerParam dto.UserRegisterParam
	if err := c.ShouldBind(&registerParam); err != nil {
		translateError := validate.TranslateError(err)
		zap.S().Errorf("绑定参数失败: %s\n", translateError)
		resp.InternalServerError(c, translateError)
		return
	}
	register, err := service.UserService.UserRegister(&registerParam)
	if err != nil {
		zap.S().Errorf("注册用户失败: %s\n", err)
		resp.InternalServerError(c, "注册用户失败")
		return
	}
	resp.Ok(c, "注册用户成功", register)
}


// UserLogin     Post 请求, 用户登陆
//
//	@Summary		Post 请求, 发送 Json 数据, 参数在消息体中 用户登陆
//	@Description	Post 请求, 发送 Json 数据, 参数在消息体中 用户登陆
//	@Tags			userApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			username	formData	string			false	"登陆的用户名"	minlength(1)	maxlength(32)
//	@Param			password	formData	string			false	"用户密码"		minlength(1)	maxlength(32)
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/user/UserLogin [post]
func (userApi) UserLogin(c *gin.Context) {
	var loginParam dto.UserLoginParam
	if err := c.ShouldBind(&loginParam); err != nil {
		translateError := validate.TranslateError(err)
		zap.S().Errorf("绑定参数失败: %s\n", translateError)
		resp.InternalServerError(c, translateError)
		return
	}
	register, err := service.UserService.UserLogin(&loginParam)
	if err != nil {
		zap.S().Errorf("用户登陆失败: %s\n", err)
		resp.InternalServerError(c, "用户登陆失败")
		return
	}
	resp.Ok(c, "用户登陆成功", register)
}

