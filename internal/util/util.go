package util

import "github.com/tarkanaciksoz/api-todo-app/internal/model"

func SetAndGetResponse(success bool, message string, data interface{}, code int) model.Response {
	return model.Response{Success: success, Message: message, Data: data, Code: code}
}

func SetAndGetTodoResponse(success bool, message string, data *model.Todo, code int) model.GetTodoResponse {
	return model.GetTodoResponse{Success: success, Message: message, Data: data, Code: code}
}

func SetAndGetTodosResponse(success bool, message string, data model.Todos, code int) model.GetTodosResponse {
	return model.GetTodosResponse{Success: success, Message: message, Data: data, Code: code}
}
