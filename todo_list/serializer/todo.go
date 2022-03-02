package serializer

import "todo_list/model"

type Todo struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Status uint `json:"status"` // 0 未完成 1 已完成
	Content string `json:"content"`
	StartTime int64 `json:"start_time"`
	EndTime int64 `json:"end_time"`
}

func BuildTodo(todo model.Todo) Todo {
	return Todo{
		ID: todo.ID,
		Title: todo.Title,
		Status: todo.Status,
		Content: todo.Content,
		StartTime: todo.StartTime,
		EndTime: todo.EndTime,
	}
}

func BuildTodos(todos []model.Todo) (results []Todo) {
	for _, item := range todos {
		todo := BuildTodo(item)
		results = append(results, todo)
	}
	return results
}
