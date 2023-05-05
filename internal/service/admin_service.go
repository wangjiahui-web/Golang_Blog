package service

import (
	"blogProject/internal/repository/cache"
	"blogProject/internal/repository/dao"
	"blogProject/model"
	"blogProject/model/dto"
	"blogProject/pkg/util"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var AdminService adminService

type adminService struct{}





func (adminService) AdminRegister(adminRegisterParam *dto.AdminRegisterParam) (rowsAffected int64, err error) {
	zap.S().Infof("用户注册接收到的密码为:%s\n", adminRegisterParam.Password)
	// 对密码进行加密
	zap.S().Infof("开始对明文密码进行加密, 明文密码为: %s\n", adminRegisterParam.Password)
	hashedPassword, err := util.GenerateFromPassword(adminRegisterParam.Password)
	zap.S().Infof("加密之后的密码为: %s\n", hashedPassword)
	if err != nil {
		zap.S().Errorf("密码加密失败: %s\n", err)
		return 0, err
	}

	// 构造 User 对象
	user := &model.Admin{
		Username: &adminRegisterParam.Username,
		Nickname: &adminRegisterParam.Nickname,
		Password: &hashedPassword,
		Website:  &adminRegisterParam.Website,
		Birth:    &adminRegisterParam.Birth,
		Tel:      &adminRegisterParam.Tel,
		Email:    &adminRegisterParam.Email,
		City:     &adminRegisterParam.City,
		School:   &adminRegisterParam.School,
	}
	register, err := dao.AdminDao.AdminRegister(dao.DB, user)
	if err != nil {
		zap.S().Errorf("注册用户失败: %s", err)
		return 0, err
	}
	return register, nil
}

func (adminService) AdminLogin(adminLoginParam *dto.AdminLoginParam) (token string, err error) {
	zap.S().Infof("用户登录前端传递过来的密码为: %s", adminLoginParam.Password)
	// 判断传入参数是否为空
	if adminLoginParam == nil {
		zap.L().Info("登录信息为空 Login 直接返回")
		return "", errors.New("登录信息为空")
	}

	// 根据用户名从数据库中查询用户

	admin, err := dao.AdminDao.AdminLogin(dao.DB, adminLoginParam.Username)
	if err != nil {
		zap.S().Errorf("从数据库中查询用户名为: %s 的用户失败: %s", adminLoginParam.Username, err)
		return "", err
	}

	zap.S().Infof("从数据库中获取的用户密码为:%s", *admin.Password)
	zap.S().Infof("接收到的前端用户用户密码为:%s", adminLoginParam.Password)

	// 比对数据库中的密码(加密之后的)和传递过来的密码(明文密码)
	err = util.CompareHashAndPassword(*admin.Password, adminLoginParam.Password)
	if err != nil {
		zap.S().Errorf("用户名或者密码错误: %s", err)
		return "", err
	}

	fmt.Println(admin.Username)

	// 密码比对成功之后, 生成 token 字符串
	generateToken, err := util.GenerateToken(*admin.Username)
	fmt.Println(generateToken)

	if err != nil {
		zap.S().Errorf("根据用户名生成 token 失败: %s\n", err)
		return "", err
	}

	// 登录成功之后, 将登录用户信息保存到 redis 中
	if err = cache.SaveLoginAdmin(admin); err != nil {
		zap.S().Infof("将登录用户存放到 redis 中失败: %s", err)
	}

	return generateToken, nil
}
