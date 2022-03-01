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

func (service *ShowTodoService) GetTodoById(tid string) serializer.Response {
	var todo model.Todo
	if err := model.DB.First(&todo, tid).Error; err != nil {
		return serializer.Response{
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