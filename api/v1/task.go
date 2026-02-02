package v1

import (
	"context"
	"net/http"
	"todo_list/pkg/utils"
	"todo_list/service"

	"github.com/cloudwego/hertz/pkg/app"
)

// CreateTask 创建任务
// @Summary 创建新任务
// @Description 创建一条新的待办事项
// @Tags 任务管理
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param request body service.CreateTaskService true "任务信息"
// @Success 200 {object} utils.Response "{"status": 200, "data": null, "msg": "创建成功"}"
// @Router /task/create [post]
func CreateTask(ctx context.Context, c *app.RequestContext) {
	var createTaskService service.CreateTaskService
	// 绑定参数
	if err := c.BindAndValidate(&createTaskService); err != nil {
		utils.JSONError(c, http.StatusBadRequest, err)
		return
	}

	//从ctx拿出中间件存入的userid
	res, _ := c.Get("user_id")
	userID := res.(uint)

	if err := createTaskService.Create(userID); err != nil {
		utils.JSONError(c, http.StatusInternalServerError, nil)
		return
	}

	utils.JSONSuccess(c, nil, "任务创建成功")
}

// GetTaskList 获取任务列表
// @Summary 获取任务列表
// @Description 支持分页、模糊搜索、状态筛选
// @Tags 任务管理
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param page_num query int false "页码 (默认1)"
// @Param page_size query int false "每页条数 (默认10)"
// @Param title query string false "搜索关键词"
// @Param status query int false "状态筛选 (0未完成, 1已完成)"
// @Success 200 {object} utils.Response{data=map[string]interface{}} "返回列表和总数"
// @Router /task/list [get]
func GetTaskList(ctx context.Context, c *app.RequestContext) {
	var listTaskService service.ListTaskService
	//绑定参数
	if err := c.BindAndValidate(&listTaskService); err != nil {
		utils.JSONError(c, http.StatusBadRequest, nil)
		return
	}

	res, _ := c.Get("user_id")
	userID := res.(uint)

	tasks, total, err := listTaskService.List(userID)
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, nil)
		return
	}

	utils.JSONSuccess(c, map[string]interface{}{
		"items": tasks,
		"total": total,
	}, "获取成功")
}

// UpdateTask 更新任务
// @Summary 更新任务
// @Description 修改任务的标题、内容或状态
// @Tags 任务管理
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param id path int true "任务ID"
// @Param request body service.UpdateTaskService true "更新信息"
// @Success 200 {object} utils.Response "修改成功"
// @Router /task/{id} [put]
func UpdateTask(ctx context.Context, c *app.RequestContext) {
	//获得路径参数
	taskID := c.Param("id")

	var updateTaskService service.UpdateTaskService

	//绑定参数
	if err := c.BindAndValidate(&updateTaskService); err != nil {
		utils.JSONError(c, http.StatusBadRequest, nil)
		return
	}

	res, _ := c.Get("user_id")
	userID := res.(uint)

	if err := updateTaskService.Update(userID, taskID); err != nil {
		utils.JSONError(c, http.StatusInternalServerError, nil)
		return
	}

	utils.JSONSuccess(c, nil, "修改成功")
}

// DeleteTask 删除任务
// @Summary 删除任务
// @Description 根据ID删除单条任务(软删除)
// @Tags 任务管理
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param id path int true "任务ID"
// @Success 200 {object} utils.Response "删除成功"
// @Router /task/{id} [delete]
func DeleteTask(ctx context.Context, c *app.RequestContext) {
	taskID := c.Param("id")

	var deleteTaskService service.DeleteTaskService

	//绑定参数
	if err := c.BindAndValidate(&deleteTaskService); err != nil {
		utils.JSONError(c, http.StatusBadRequest, nil)
		return
	}

	res, _ := c.Get("user_id")
	userID := res.(uint)

	if err := deleteTaskService.Delete(userID, taskID); err != nil {
		utils.JSONError(c, http.StatusInternalServerError, nil)
		return
	}

	utils.JSONSuccess(c, nil, "删除成功")
}

// BatchUpdateStatus 批量更新状态
// @Summary 批量更新状态
// @Description 一键将所有任务标记为完成或未完成
// @Tags 任务管理 (批量)
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param request body service.BatchTaskService true "目标状态 (target_status: 0或1)"
// @Success 200 {object} utils.Response "批量更新成功"
// @Router /task/batch_update [post]
func BatchUpdateStatus(ctx context.Context, c *app.RequestContext) {
	var batchService service.BatchTaskService
	if err := c.BindAndValidate(&batchService); err != nil {
		utils.JSONError(c, http.StatusBadRequest, nil)
		return
	}

	res, _ := c.Get("user_id")
	userID := res.(uint)

	if err := batchService.BatchUpdateStatus(userID); err != nil {
		utils.JSONError(c, http.StatusInternalServerError, nil)
		return
	}

	utils.JSONSuccess(c, nil, "批量更新状态成功")
}

// BatchDelete 批量删除
// @Summary 批量删除
// @Description 一键删除任务 (1:已完成, 2:未完成, 3:全部)
// @Tags 任务管理 (批量)
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param request body service.BatchTaskService true "删除类型 (delete_type)"
// @Success 200 {object} utils.Response "批量删除成功"
// @Router /task/batch_delete [delete]
func BatchDelete(ctx context.Context, c *app.RequestContext) {
	var batchService service.BatchTaskService
	if err := c.BindAndValidate(&batchService); err != nil {
		utils.JSONError(c, http.StatusBadRequest, nil)
		return
	}

	res, _ := c.Get("user_id")
	userID := res.(uint)

	if err := batchService.BatchDelete(userID); err != nil {
		utils.JSONError(c, http.StatusInternalServerError, nil)
		return
	}

	utils.JSONSuccess(c, nil, "批量删除成功")
}
