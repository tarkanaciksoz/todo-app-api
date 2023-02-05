package todo

import (
	"log"

	"github.com/tarkanaciksoz/api-todo-app/internal/model"
)

type TodoService struct {
	L  *log.Logger
	DB DB
}

type Service interface {
	Get(id int) (*model.Todo, error)
	List() []*model.Todo
	Create(t *model.Todo) (*model.Todo, error)
	Mark(t *model.Todo) (*model.Todo, error)
	Delete(id int) error
}

func NewTodoService(l *log.Logger, db DB) TodoService {
	return TodoService{
		L:  l,
		DB: db,
	}
}

func (ts TodoService) Get(id int) (*model.Todo, error) {
	return ts.DB.Get(id)
}

func (ts TodoService) List() []*model.Todo {
	return ts.DB.List()
}

func (ts TodoService) Create(todo *model.Todo) (*model.Todo, error) {
	return ts.DB.Create(todo)
}

func (ts TodoService) Mark(todo *model.Todo) (*model.Todo, error) {
	return ts.DB.Mark(todo)
}

func (ts TodoService) Delete(id int) error {
	return ts.DB.Delete(id)
}
