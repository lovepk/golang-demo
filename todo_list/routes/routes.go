package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"todo_list/api"
)

func NewRouter() *gin.Engine  {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		//	用户操作
		v1.POST("user/register", api.UserRegister)
	}
	return r
}