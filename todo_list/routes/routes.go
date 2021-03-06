package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"todo_list/api"
	"todo_list/middleware"
)

func NewRouter() *gin.Engine  {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		//	用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("todo", api.CreatTodo)
			authed.GET("todo/:id", api.ShowTodo)
			authed.GET("todos", api.ListTodo)
			authed.PUT("todo/:id", api.UpdateTodo)
			authed.POST("todo/search", api.SearchListTodo)
			authed.DELETE("todo/:id", api.DeleteTodo)
		}
	}
	return r
}
