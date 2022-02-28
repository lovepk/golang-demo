package api

import (
	"github.com/gin-gonic/gin"
	"todo_list/service"
)

func UserRegister(c *gin.Context)  {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err != nil {
		c.JSON(400, err)
	} else {
		res := userService.Register()
		c.JSON(200, res)
	}
}