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
	claims,_ := utils.ParseToken(c.GetHeader("authorization"))
	res := showTodoService.GetTodoById(claims.Id, c.Param("id"))
	c.JSON(200, res)
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

func UpdateTodo(c *gin.Context)  {
	var updateTodoService service.UpdateTodoService
	claims,_ := utils.ParseToken(c.GetHeader("authorization"))
	if err := c.ShouldBind(&updateTodoService); err != nil {
		c.JSON(400, err)
	} else {
		res := updateTodoService.UpdateTodo(claims.Id, c.Param("id"))
		c.JSON(200, res)
	}
}

func SearchListTodo(c *gin.Context)  {
	var searchTodoService service.SearchTodoService
	claims,_ := utils.ParseToken(c.GetHeader("authorization"))
	if err := c.ShouldBind(&searchTodoService); err != nil {
		c.JSON(400, err)
	} else {
		res := searchTodoService.Search(claims.Id)
		c.JSON(200, res)
	}
}

func DeleteTodo(c *gin.Context)  {
	var deleteTodoService service.DeleteTodoService
	claims, _ := utils.ParseToken(c.GetHeader("authorization"))
	res := deleteTodoService.Delete(claims.Id, c.Param("id"))
	c.JSON(200, res)
}
