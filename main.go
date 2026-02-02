package main

import (
	"todo_list/dao"
	"todo_list/model"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// @title           TodoList API
// @version         1.0
// @description     这是一个基于 Hertz + GORM 的待办事项管理系统
// @host            localhost:8888
// @BasePath        /api/v1

func main() {
	// 初始化数据库
	dao.InitDB()

	// 绑定模型 (自动建表/同步表结构)
	model.Migration(dao.DB)

	// 初始化 Hertz
	h := server.Default()

	// 加载路由
	InitRouter(h)

	//启动
	h.Spin()
}
