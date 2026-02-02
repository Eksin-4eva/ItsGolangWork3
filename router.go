package main

import (
	v1 "todo_list/api/v1"
	"todo_list/middleware"

	_ "todo_list/docs"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

// init router
func InitRouter(h *server.Hertz) {
	//swagger
	url := swagger.URL("http://localhost:8888/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
	// 👆👆👆
	// 创建一个路由组 /api/v1
	apiV1 := h.Group("/api/v1")
	{
		// 公开接口
		userGroup := apiV1.Group("/user")
		{
			// POST /api/v1/user/register
			userGroup.POST("/register", v1.UserRegister)
			userGroup.POST("/login", v1.UserLogin)
		}
		// 私有接口
		taskGroup := apiV1.Group("/task")
		taskGroup.Use(middleware.JWTAuth())
		{
			// POST /api/v1/task/create
			taskGroup.POST("/create", v1.CreateTask)
			//GET  /api/v1/task/list
			taskGroup.GET("/list", v1.GetTaskList)
			// PUT /api/v1/task/
			taskGroup.PUT("/:id", v1.UpdateTask)
			// DELETE /api/v1/task/1
			taskGroup.DELETE("/:id", v1.DeleteTask)
			//批量操作路由
			taskGroup.POST("/batch_update", v1.BatchUpdateStatus)
			taskGroup.DELETE("/batch_delete", v1.BatchDelete)
		}
	}
}
