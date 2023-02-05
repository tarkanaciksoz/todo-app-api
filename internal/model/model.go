package model

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Middleware  *http.HandlerFunc
}

type Routes []Route

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

type GetTodoResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *Todo  `json:"data"`
	Code    int    `json:"code"`
}

type GetTodosResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    []*Todo `json:"data"`
	Code    int     `json:"code"`
}

type Config struct {
	AppEnv      string
	BindAddress string
}
