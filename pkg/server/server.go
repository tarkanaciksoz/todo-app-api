package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/tarkanaciksoz/api-todo-app/internal/model"
	"github.com/tarkanaciksoz/api-todo-app/internal/todo"
	"github.com/tarkanaciksoz/api-todo-app/internal/util"
)

func Init(logger *log.Logger) *mux.Router {
	db := todo.NewDB()
	todoService := todo.NewTodoService(logger, db)
	todoHandler := todo.NewTodoHandler(todoService)

	mappedRoutes := make(map[string]model.Routes)
	mappedRoutes[http.MethodGet] = model.Routes{
		model.Route{
			Name:        "GET TODO",
			Method:      http.MethodGet,
			Pattern:     "/todo/getTodo/{id:[0-9]+}",
			HandlerFunc: todoHandler.GetTodo,
		},
		model.Route{
			Name:        "GET TODO LIST",
			Method:      http.MethodGet,
			Pattern:     "/todo/getTodos",
			HandlerFunc: todoHandler.ListTodos,
		},
	}

	mappedRoutes[http.MethodPost] = model.Routes{
		model.Route{
			Name:        "CREATE NEW TODO",
			Method:      http.MethodPost,
			Pattern:     "/todo/createTodo",
			HandlerFunc: todoHandler.CreateTodo,
		},
	}

	mappedRoutes[http.MethodPut] = model.Routes{
		model.Route{
			Name:        "MARK - UNMARK TODO",
			Method:      http.MethodPut,
			Pattern:     "/todo/markTodo/{id:[0-9]+}",
			HandlerFunc: todoHandler.MarkTodo,
		},
	}

	mappedRoutes[http.MethodDelete] = model.Routes{
		model.Route{
			Name:        "DELETE TODO",
			Method:      http.MethodDelete,
			Pattern:     "/todo/deleteTodo/{id:[0-9]+}",
			HandlerFunc: todoHandler.DeleteTodo,
		},
	}

	router := mux.NewRouter()
	router.NotFoundHandler = MethodNotFoundHandler(logger)
	router.MethodNotAllowedHandler = MethodNotAllowedHandler(logger)

	for method, routes := range mappedRoutes {
		methodRout := router.Methods(method).Subrouter()
		methodRout.Use(ApplicationRecovery)
		methodRout.Use(Middleware)
		for _, route := range routes {
			methodRout.HandleFunc(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

func MethodNotFoundHandler(logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		logger.Printf("No Route Found With Pattern : %s", r.URL.Path)

		rw.Header().Add("Content-Type", "application/json")
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		rw.Header().Add("Access-Control-Allow-Headers", "*")
		rw.Header().Add("Access-Control-Allow-Credentials", "true")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(util.SetAndGetResponse(false, "Method Not Found", nil, http.StatusNotFound))
	})
}

func MethodNotAllowedHandler(logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		logger.Printf("Method %s Not Allowed For Pattern : %s", r.Method, r.URL.Path)

		rw.Header().Add("Content-Type", "application/json")
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		rw.Header().Add("Access-Control-Allow-Headers", "*")
		rw.Header().Add("Access-Control-Allow-Credentials", "true")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(util.SetAndGetResponse(false, "Method Not Allowed", nil, http.StatusBadRequest))
	})
}

func ApplicationRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				_, err := fmt.Fprintln(os.Stderr, "Recovered from application error occurred")
				if err != nil {
					rw.WriteHeader(http.StatusOK)
					json.NewEncoder(rw).Encode(util.SetAndGetResponse(false, "Internal Server Error at Application Recovery", nil, http.StatusInternalServerError))
					return
				}
				_, _ = fmt.Fprintln(os.Stderr, err)

				rw.WriteHeader(http.StatusOK)
				json.NewEncoder(rw).Encode(util.SetAndGetResponse(false, "Internal Server Error", nil, http.StatusInternalServerError))
				return
			}
		}()

		ctx := r.Context()
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
