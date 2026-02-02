//	第一步：编写 API 控制层 (Controller)
//	负责：
//
//	接收 前端的 JSON 请求。
//
//	绑定 数据到 UserRegisterService 结构体。
//
//	调用 Service 层的方法。
//
//	返回 结果给前端。

package v1

import (
	"context"
	"net/http"
	"todo_list/pkg/utils"
	"todo_list/service"

	"github.com/cloudwego/hertz/pkg/app"
)

// 用户注册API
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var userRegisterService service.UserRegisterService
	//绑定参数
	if err := c.BindAndValidate(&userRegisterService); err != nil {
		utils.JSONError(c, http.StatusBadRequest, err)
		return
	}

	//调用service>注册逻辑
	if err := userRegisterService.Register(); err != nil {
		utils.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	//返回成功响应
	utils.JSONSuccess(c, nil, "注册成功")
}

func UserLogin(ctx context.Context, c *app.RequestContext) {
	var userLoginService service.UserLoginService

	// 绑定参数并校验
	if err := c.BindAndValidate(&userLoginService); err != nil {
		utils.JSONError(c, http.StatusBadRequest, nil)
		return
	}

	//调用service层登录逻辑
	token, err := userLoginService.Login()
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, nil)
		return
	}

	//返回token给前端
	utils.JSONSuccess(c, map[string]string{
		"token": token,
	}, "登录成功")
}
