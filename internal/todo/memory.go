package todo

import (
	"errors"
	"sort"
	"strconv"

	"github.com/tarkanaciksoz/api-todo-app/internal/model"
)

type Memory struct {
	Todos map[int]*model.Todo
}

type DB interface {
	Get(id int) (*model.Todo, error)
	List() []*model.Todo
	Create(t *model.Todo) (*model.Todo, error)
	Mark(t *model.Todo) (*model.Todo, error)
	Delete(id int) error
}

func NewDB() DB {
	return Memory{
		Todos: make(map[int]*model.Todo),
	}
}

func (m Memory) Get(id int) (*model.Todo, error) {
	if !(id > 0) {
		return nil, errors.New("todo ID Must Be Greater Than Zero")
	}

	todo, exists := m.Todos[id]
	if !exists {
		return nil, errors.New("no todo found with id:" + strconv.Itoa(id))
	}

	return todo, nil
}

func (m Memory) List() []*model.Todo {
	todos := []*model.Todo{}

	keys := make([]int, 0, len(m.Todos))
	for k := range m.Todos {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		todos = append(todos, m.Todos[k])
	}

	return todos
}

func (m Memory) Create(todo *model.Todo) (*model.Todo, error) {
	id := todo.ID

	if !(id > 0) {
		return nil, errors.New("todo ID Must Be Greater Than Zero")
	}

	for {
		_, exists := m.Todos[id]
		if !exists {
			break
		}
		id++
	}

	todo.ID = id
	m.Todos[id] = todo
	return m.Todos[id], nil
}

func (m Memory) Mark(todo *model.Todo) (*model.Todo, error) {
	_, exists := m.Todos[todo.ID]
	if !exists {
		return nil, errors.New("no todo found with id:" + strconv.Itoa(todo.ID))
	}

	m.Todos[todo.ID] = todo

	return m.Todos[todo.ID], nil
}

func (m Memory) Delete(id int) error {
	_, exists := m.Todos[id]
	if !exists {
		return errors.New("no todo found with id:" + strconv.Itoa(id))
	}

	delete(m.Todos, id)
	return nil
}
