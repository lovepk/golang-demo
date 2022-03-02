package service

import (
	"time"
	"todo_list/model"
	"todo_list/serializer"
)

type CreateTodoService struct {
	Title string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
	Status int `form:"status" json:"status"`
}

type ShowTodoService struct {
	
}

type ListTodoService struct {
	PageNum int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type UpdateTodoService struct {
	Title string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
	Status int `form:"status" json:"status"`
}

type SearchTodoService struct {
	Info string `json:"info" form:"info"`
	PageNum int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type DeleteTodoService struct {
	
}

func (service *CreateTodoService) Create(id uint) serializer.Response  {
	var user model.User
	model.DB.First(&user, id)
	todo := model.Todo{
		User: user,
		Title: service.Title,
		Content: service.Content,
		Status: 0,
		UserID: user.ID,
		StartTime: time.Now().Unix(),
		EndTime: 0,
	}
	if err := model.DB.Create(&todo).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg: "代办事项创建失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg: "代办事项创建成功",
	}
}

func (service *ShowTodoService) GetTodoById(uid uint, tid string) serializer.Response {
	var todo model.Todo
	if err := model.DB.Model(&model.Todo{}).Where("user_id=?", uid).First(&todo, tid).Error; err != nil {
		return serializer.Response {
			Status: 500,
			Msg: "查询失败",
		}
	}
	return serializer.Response {
		Status: 200,
		Data: serializer.BuildTodo(todo),
		Msg: "查询成功",
	}
}

func (service *ListTodoService) GetListTodo(uid uint) serializer.Response  {
	var todos []model.Todo
	var count int64 = 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	if err := model.DB.Model(&model.Todo{}).Where("user_id=?", uid).Count(&count).Preload("User").
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&todos).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "查询失败"+err.Error(),
		}
	}
	return serializer.BuildListResponse(todos, uint(count))
}

func (service UpdateTodoService) UpdateTodo(uid uint, tid string) serializer.Response {
	var todo model.Todo
	if err := model.DB.Model(&model.Todo{}).Where("user_id=?", uid).First(&todo, tid).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg: "没有查到这条记录",
		}
	}
	todo.Title = service.Title
	todo.Content = service.Content
	todo.Status = uint(service.Status)
	model.DB.Save(&todo)
	return serializer.Response{
		Status: 200,
		Msg: "更新成功",
	}
}

func (service *SearchTodoService) Search(uid uint) serializer.Response {
	var todos []model.Todo
	var count int64 = 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	println(service.Info)
	if err := model.DB.Model(&model.Todo{}).Where("user_id=?", uid).
		Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%").Count(&count).Preload("User").
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&todos).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "查询失败"+err.Error(),
		}
	}
	return serializer.BuildListResponse(todos, uint(count))
}

func (service *DeleteTodoService) Delete(uid uint, tid string) serializer.Response {
	var todo model.Todo
	if err := model.DB.Model(&model.Todo{}).Where("user_id=?", uid).First(&todo, tid).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg: "没有查到这条记录",
		}
	}
	if err := model.DB.Delete(&todo).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg: "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}