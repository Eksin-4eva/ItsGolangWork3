package service

import (
	"errors"
	"todo_list/dao"
	"todo_list/model"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` // 0未完成 1已完成
}

type ListTaskService struct {
	PageNum  int    `json:"page_num" form:"page_num" query:"page_num"`    //页码
	PageSize int    `json:"page_size" form:"page_size" query:"page_size"` //每页条数
	Title    string `json:"title" form:"title" query:"title"`
	Status   *int   `json:"status" form:"status" query:"status"`
}

type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` // 0未完成 1已完成
}

type DeleteTaskService struct {
}

type BatchTaskService struct {
	TargetStatus int `json:"target_status" form:"targer_status"`
	DeleteType   int `json:"delete_type" form:"delete_type"`
}

// 创建任务逻辑
// 传入从middleware解析出来的userID

// 创建任务
func (service *CreateTaskService) Create(userID uint) error {
	task := model.Task{
		UserId:  userID,
		Title:   service.Title,
		Content: service.Content,
		Status:  service.Status,
	}
	return dao.DB.Create(&task).Error
}

// 查询任务列表
func (service *ListTaskService) List(userID uint) ([]model.Task, int64, error) {
	var tasks []model.Task
	var total int64

	// default value
	if service.PageNum == 0 {
		service.PageNum = 1
	}
	if service.PageSize == 0 {
		service.PageSize = 10
	}

	tx := dao.DB.Model(&model.Task{}).Where("user_id = ?", userID)
	if service.Title != "" {
		// 模糊查找标题
		tx = tx.Where("title LIKE ?", "%"+service.Title+"%")
	}

	if service.Status != nil {
		// 查找状态为 0 / 1, nil 为全查找
		tx = tx.Where("status = ?", *service.Status)
	}

	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	//分页查询
	//limit, offset, order, find
	err := tx.Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).
		Order("created_at DESC").
		Find(&tasks).Error
	return tasks, total, err
}

// 更新Task
func (service *UpdateTaskService) Update(userID uint, taskID string) error {
	var task model.Task

	//查找：满足ID匹配且属于当前用户

	if err := dao.DB.
		Where("id = ? AND user_id = ?", taskID, userID).
		First(&task).
		Error; err != nil {
		return errors.New("任务不存在/无权更改")
	}

	task.Title = service.Title
	task.Content = service.Content
	task.Status = service.Status

	return dao.DB.Save(&task).Error
}

// 删除单个Task
func (service *DeleteTaskService) Delete(userID uint, taskID string) error {
	var task model.Task

	//查找：满足ID匹配且属于当前用户
	err := dao.DB.
		Where("id = ? AND user_id = ?", taskID, userID).
		First(&task).
		Error
	if err != nil {
		return errors.New("任务不存在/无权限")
	}

	//执行软删除
	return dao.DB.Delete(&task).Error
}

// 批量更新状态
func (service *BatchTaskService) BatchUpdateStatus(userID uint) error {
	return dao.DB.Model(&model.Task{}).Where("user_id = ?", userID).Update("status", service.TargetStatus).Error
}

// 批量删除
func (service *BatchTaskService) BatchDelete(userID uint) error {
	tx := dao.DB.Model(&model.Task{}).Where("user_id = ?", userID)

	if service.DeleteType == 1 {
		//删除已完成事项
		tx = tx.Where("status = ?", 1)
	} else if service.DeleteType == 2 {
		//删除未完成事项
		tx = tx.Where("status = ?", 0)
	} else if service.DeleteType == 3 {
		//全删
		//不修改tx
	} else {
		return nil
	}
	//软删除
	return tx.Delete(&model.Task{}).Error
}
