package service

import (
	"errors"
	"todo_list/dao"
	"todo_list/model"
	"todo_list/pkg/utils"
)

// 用户注册服务的参数模型
type UserRegisterService struct {
	UserName string `json:"user_name" form:"user_name"`
	PassWord string `json:"password" form:"password"`
}

type UserLoginService struct {
	UserName string `json:"user_name" form:"user_name"`
	PassWord string `json:"password" form:"password"`
}

// Register 注册逻辑
func (service *UserRegisterService) Register() error {
	var count int64

	//验证用户名是否存在
	//Model(&model.User{})指定查询哪个表
	//where查询条件
	//count统计数量
	dao.DB.Model(&model.User{}).Where("user_name= ?", service.UserName).Count(&count)
	if count > 0 {
		return errors.New("用户名已存在")
	}

	//密码加密
	encryptedPassword, err := utils.CryptPassword(service.PassWord)
	if err != nil {
		return errors.New("密码加密失败")
	}

	user := model.User{
		UserName:       service.UserName,
		PasswordDigest: encryptedPassword,
	}

	if err := dao.DB.Create(&user).Error; err != nil {
		return errors.New("数据库保存失败")
	}
	return nil
}

// Login逻辑
func (service *UserLoginService) Login() (string, error) {
	var user model.User
	// Where("user_name = ?", ...) 相当于 SQL 里的 WHERE 子句
	// First(&user) 相当于 LIMIT 1，并将结果填充到 user 变量中
	if err := dao.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return "", errors.New("用户不存在或数据库错误")
	}

	//验证密码
	if !utils.CheckPassword(service.PassWord, user.PasswordDigest) {
		return "", errors.New("密码错误")
	}

	//分发Token
	token, err := utils.GenerateToken(user.ID, service.UserName)
	if err != nil {
		return "", errors.New("Token签发失败")
	}

	return token, nil
}
