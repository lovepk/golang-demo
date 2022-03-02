package api

import (
	"github.com/gin-gonic/gin"
	"todo_list/pkg/utils"
	"todo_list/service"
)

func CreatTodo(c *gin.Context)  {
	var createTodoService service.CreateTodoService
	claims,_ := utils.ParseToken(c.GetHeader("authorization"))
	if err := c.ShouldBind(&createTodoService); err != nil {
		c.JSON(400, err)
	} else {
		res := createTodoService.Create(claims.Id)
		c.JSON(200, res)
	}
}

func ShowTodo(c *gin.Context)  {
	var showTodoService service.ShowTodoService
	if err := c.ShouldBind(&showTodoService); err != nil {
		c.JSON(400, err)
	} else {
		res := showTodoService.GetTodoById(c.Param("id"))
		c.JSON(200, res)
	}
}

func ListTodo(c *gin.Context)  {
	var listTodoService service.ListTodoService
	claims,_ := utils.ParseToken(c.GetHeader("authorization"))
	if err := c.ShouldBind(&listTodoService); err != nil {
		c.JSON(400, err)
	} else {
		res := listTodoService.GetListTodo(claims.Id)
		c.JSON(200, res)
	}
}
