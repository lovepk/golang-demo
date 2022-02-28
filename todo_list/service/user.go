package service

import (
	"todo_list/model"
	"todo_list/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json: "user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding: "required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response  {
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("user_name", service.UserName).
		Find(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg: "已经有这个人了，无需再注册",
		}
	}
	user.UserName = service.UserName
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg: err.Error(),
		}
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg: err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg: "用户注册成功",
	}
}