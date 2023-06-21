package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tarkanaciksoz/api-todo-app/internal/model"
	"github.com/tarkanaciksoz/api-todo-app/internal/util"
)

type TodoHandler struct {
	ts Service
}

type Handler interface {
	GetTodo(rw http.ResponseWriter, r *http.Request)
	ListTodos(rw http.ResponseWriter, _ *http.Request)
	CreateTodo(rw http.ResponseWriter, r *http.Request)
	MarkTodo(rw http.ResponseWriter, r *http.Request)
	DeleteTodo(rw http.ResponseWriter, r *http.Request)
}

func NewTodoHandler(ts Service) Handler {
	return &TodoHandler{
		ts: ts,
	}
}

func (th *TodoHandler) GetTodo(rw http.ResponseWriter, r *http.Request) {
	th.ts.Log("Handle GetTodo method")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		th.ts.Log("Unable to resolve id " + vars["id"] + ": " + err.Error())
		json.NewEncoder(rw).Encode(util.SetAndGetResponse(false, "Unable to resolve id : "+vars["id"], nil, http.StatusBadRequest))
		return
	}

	todo, err := th.ts.Get(id)
	if err != nil {
		th.ts.Log(err.Error())
		json.NewEncoder(rw).Encode(util.SetAndGetResponse(false, err.Error(), nil, http.StatusBadRequest))
		return
	}

	json.NewEncoder(rw).Encode(util.SetAndGetTodoResponse(true, "Todo Listed Successfully", todo, http.StatusOK))
	th.ts.Log("GetTodo method successfully handled")
}

func (th *TodoHandler) ListTodos(rw http.ResponseWriter, _ *http.Request) {
	th.ts.Log("Handle ListTodos method")

	todos := th.ts.List()

	json.NewEncoder(rw).Encode(util.SetAndGetTodosResponse(true, "Todos Listed Successfully", todos, http.StatusOK))
	th.ts.Log("ListTodos method successfully handled")
}

func (th *TodoHandler) CreateTodo(rw http.ResponseWriter, r *http.Request) {
	th.ts.Log("Handle CreateTodo method")

	todo := &model.Todo{}
	err := todo.FromJSON(r.Body)
	if err != nil {
		th.ts.Log("Request Body Couldn't Resolved - Invalid JSON Data : " + err.Error())
		json.NewEncoder(rw).Encode(util.SetAndGetResponse(false, "Invalid JSON Data", nil, http.StatusBadRequest))
		return
	}

	todo, err = th.ts.Create(todo)
	if err != nil {
		th.ts.Log(err.Error())
		json.NewEncoder(rw).Encode(util.SetAndGetResponse(false, err.Error(), nil, http.StatusBadRequest))
		return
	}

	json.NewEncoder(rw).Encode(util.SetAndGetTodoResponse(true, "Todo Created Successfully", todo, http.StatusOK))
	th.ts.Log("CreateTodo method successfully handled")
}

func (th *TodoHandler) MarkTodo(rw http.ResponseWriter, r *http.Request) {
	th.ts.Log("Handle MarkTodo method")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		th.ts.Log("Unable to convert id " + vars["id"] + ": " + err.Error())
		json.NewEncoder(rw).Encode(util.SetAndGetResponse(false, "Unable to convert id : "+vars["id"], nil, http.StatusBadRequest))
		return
	}

	todo := &model.Todo{}
	err = todo.FromJSON(r.Body)
	if err != nil {
		th.ts.Log("Request Body Couldn't Resolved - Invalid JSON Data : " + err.Error())
		json.NewEncoder(rw).Encode(util.SetAndGetResponse(false, "Invalid JSON Data", nil, http.StatusBadRequest))
		return
	}
	todo.ID = id

	todo, err = th.ts.Mark(todo)
	if err != nil {
		th.ts.Log(err.Error())
		json.NewEncoder(rw).Encode(util.SetAndGetResponse(false, err.Error(), nil, http.StatusBadRequest))
		return
	}

	json.NewEncoder(rw).Encode(util.SetAndGetTodoResponse(true, "Todo Marked Successfully", todo, http.StatusOK))
	th.ts.Log("MarkTodo method successfully handled")
}

func (th *TodoHandler) DeleteTodo(rw http.ResponseWriter, r *http.Request) {
	th.ts.Log("Handle DeleteTodo method")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		th.ts.Log("Unable to convert id " + vars["id"] + ": " + err.Error())
		json.NewEncoder(rw).Encode(util.SetAndGetResponse(false, "Unable to convert id : "+vars["id"], nil, http.StatusBadRequest))
		return
	}

	err = th.ts.Delete(id)
	if err != nil {
		th.ts.Log(err.Error())
		json.NewEncoder(rw).Encode(util.SetAndGetResponse(false, err.Error(), nil, http.StatusBadRequest))
		return
	}

	json.NewEncoder(rw).Encode(util.SetAndGetResponse(true, "Todo Deleted Successfully", nil, http.StatusOK))
	th.ts.Log("DeleteTodo method successfully handled")
}
