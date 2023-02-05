package todo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	mockHandler "github.com/tarkanaciksoz/api-todo-app/internal/mocks"
	"github.com/tarkanaciksoz/api-todo-app/internal/model"
	"github.com/tarkanaciksoz/api-todo-app/internal/util"
)

type HandlerSuite struct {
	suite.Suite
	*require.Assertions
	ctrl *gomock.Controller

	mockHandler *mockHandler.MockHandler
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (s *HandlerSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())

	s.mockHandler = mockHandler.NewMockHandler(s.ctrl)
}

func (s *HandlerSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *HandlerSuite) TestHandlerGivenWhenGetTodoIsCalled() {
	s.T().Run("TestHandlerGivenExistingTodoIdWhenGetTodoIsCalledThenItSouldReturnGetTodoResponseWithSuccessTrueAndTodoListedSuccessfullyMessageAndDataWithTodoAndCodeAs200", func(t *testing.T) {
		expectedResponse := util.SetAndGetTodoResponse(true, "Todo Listed Successfully", &model.Todo{
			ID:     1,
			Value:  "buy some milk",
			Marked: 0,
		}, 200)

		w := httptest.NewRecorder()
		json.NewEncoder(w).Encode(expectedResponse)
		r := httptest.NewRequest(http.MethodGet, "/todo/getTodo/1", nil)

		s.mockHandler.EXPECT().GetTodo(w, r).Return()
		s.mockHandler.GetTodo(w, r)

		resp, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			t.Error("Unresolved Json Body")
		}

		cleanResponse := model.GetTodoResponse{}
		err = json.Unmarshal([]byte(resp), &cleanResponse)
		if err != nil {
			t.Error("Json Response Doesn't Match")
		}

		s.Assertions.Equal(expectedResponse, cleanResponse)
		s.Assertions.Equal(http.StatusOK, w.Result().StatusCode)
	})

	s.T().Run("TestHandlerGivenUnExistingTodoIdWhenGetTodoIsCalledThenItSouldReturnGetTodoResponseWithSuccessFalseAndNoTodoFoundWithId1MessageAndDataWithNilAndCodeAs400", func(t *testing.T) {
		expectedResponse := util.SetAndGetTodoResponse(false, "no todo found with id:1", nil, 400)

		w := httptest.NewRecorder()
		json.NewEncoder(w).Encode(expectedResponse)
		r := httptest.NewRequest(http.MethodGet, "/todo/getTodo/1", nil)

		s.mockHandler.EXPECT().GetTodo(w, r).Return()
		s.mockHandler.GetTodo(w, r)

		resp, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			t.Error("Unresolved Json Body")
		}

		cleanResponse := model.GetTodoResponse{}
		err = json.Unmarshal([]byte(resp), &cleanResponse)
		if err != nil {
			t.Error("Json Response Doesn't Match")
		}

		s.Assertions.Equal(expectedResponse, cleanResponse)
		s.Assertions.Equal(http.StatusOK, w.Result().StatusCode)
	})
}

func (s *HandlerSuite) TestHandlerGivenWhenListTodosIsCalled() {
	s.T().Run("TestHandlerGiven2TodosWhenListTodosIsCalledThenItSouldReturnGetTodosResponseWithSuccessTrueAndTodosListedSuccessfullyMessageAndDataWith2TodosAndCodeAs200", func(t *testing.T) {
		expectedResponse := util.SetAndGetTodosResponse(true, "Todos Listed Successfully", model.Todos{
			&model.Todo{
				ID:     1,
				Value:  "buy some milk",
				Marked: 0,
			},
			&model.Todo{
				ID:     2,
				Value:  "enjoy the assignment",
				Marked: 0,
			},
		}, 200)

		w := httptest.NewRecorder()
		json.NewEncoder(w).Encode(expectedResponse)
		r := httptest.NewRequest(http.MethodGet, "/todo/getTodos", nil)

		s.mockHandler.EXPECT().ListTodos(w, r).Return()
		s.mockHandler.ListTodos(w, r)

		resp, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			t.Error("Unresolved Json Body")
		}

		cleanResponse := model.GetTodosResponse{}
		err = json.Unmarshal([]byte(resp), &cleanResponse)
		if err != nil {
			t.Error("Json Response Doesn't Match")
		}

		s.Assertions.Equal(expectedResponse, cleanResponse)
		s.Assertions.Equal(http.StatusOK, w.Result().StatusCode)
	})

	s.T().Run("TestHandlerGivenEmptyTodosWhenListTodosIsCalledThenItSouldReturnGetTodosResponseWithSuccessTrueAndTodosListedSuccessfullyMessageAndDataWithEmptyTodoObjectAndCodeAs200", func(t *testing.T) {
		expectedResponse := util.SetAndGetTodosResponse(true, "Todos Listed Successfully", model.Todos{}, 200)

		w := httptest.NewRecorder()
		json.NewEncoder(w).Encode(expectedResponse)
		r := httptest.NewRequest(http.MethodGet, "/todo/getTodos", nil)

		s.mockHandler.EXPECT().ListTodos(w, r).Return()
		s.mockHandler.ListTodos(w, r)

		resp, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			t.Error("Unresolved Json Body")
		}

		cleanResponse := model.GetTodosResponse{}
		err = json.Unmarshal([]byte(resp), &cleanResponse)
		if err != nil {
			t.Error("Json Response Doesn't Match")
		}

		s.Assertions.Equal(expectedResponse, cleanResponse)
		s.Assertions.Equal(http.StatusOK, w.Result().StatusCode)
	})
}

func (s *HandlerSuite) TestHandlerGivenWhenCreateTodoIsCalled() {
	s.T().Run("TestHandlerGivenTodoWhenCreateTodoIsCalledThenItSouldReturnGetTodoResponseWithSuccessTrueAndTodoCreatedSuccessfullyMessageAndDataWithTodoAndCodeAs200", func(t *testing.T) {
		todo := &model.Todo{
			ID:     1,
			Value:  "buy some milk",
			Marked: 0,
		}
		expectedResponse := util.SetAndGetTodoResponse(true, "Todo Created Successfully", todo, 200)

		var body bytes.Buffer
		json.NewEncoder(&body).Encode(todo)

		w := httptest.NewRecorder()
		json.NewEncoder(w).Encode(expectedResponse)
		r := httptest.NewRequest(http.MethodPost, "/todo/createTodo", &body)

		s.mockHandler.EXPECT().CreateTodo(w, r).Return()
		s.mockHandler.CreateTodo(w, r)

		resp, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			t.Error("Unresolved Json Body")
		}

		cleanResponse := model.GetTodoResponse{}
		err = json.Unmarshal([]byte(resp), &cleanResponse)
		if err != nil {
			t.Error("Json Response Doesn't Match")
		}

		s.Assertions.Equal(expectedResponse, cleanResponse)
		s.Assertions.Equal(http.StatusOK, w.Result().StatusCode)
	})

	s.T().Run("TestHandlerGivenTodoWithExistingTodoIDWhenCreateTodoIsCalledThenItSouldReturnGetTodoResponseWithSuccessTrueAndTodoCreatedSuccessfullyMessageAndDataWithTodoWithNewIDAndCodeAs200", func(t *testing.T) {
		todo := &model.Todo{
			ID:     1,
			Value:  "buy some milk",
			Marked: 0,
		}
		var body bytes.Buffer
		json.NewEncoder(&body).Encode(todo)

		todo.ID = 2
		expectedResponse := util.SetAndGetTodoResponse(true, "Todo Created Successfully", todo, 200)

		w := httptest.NewRecorder()
		json.NewEncoder(w).Encode(expectedResponse)
		r := httptest.NewRequest(http.MethodPost, "/todo/createTodo", &body)

		s.mockHandler.EXPECT().CreateTodo(w, r).Return()
		s.mockHandler.CreateTodo(w, r)

		resp, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			t.Error("Unresolved Json Body")
		}

		cleanResponse := model.GetTodoResponse{}
		err = json.Unmarshal([]byte(resp), &cleanResponse)
		if err != nil {
			t.Error("Json Response Doesn't Match")
		}

		s.Assertions.Equal(expectedResponse, cleanResponse)
		s.Assertions.Equal(http.StatusOK, w.Result().StatusCode)
	})
}

func (s *HandlerSuite) TestHandlerGivenWhenMarkTodoIsCalled() {
	s.T().Run("TestHandlerGivenUnMarkedIdWhenMarkTodoIsCalledThenItSouldReturnGetTodoResponseWithSuccessTrueAndTodoMarkedSuccessfullyMessageAndDataWithTodo.MarkedAs1AndCodeAs200", func(t *testing.T) {
		todo := &model.Todo{
			ID:     1,
			Value:  "buy some milk",
			Marked: 0,
		}

		var body bytes.Buffer
		json.NewEncoder(&body).Encode(todo)

		todo.Marked = 1
		expectedResponse := util.SetAndGetTodoResponse(true, "Todo Marked Successfully", todo, 200)

		w := httptest.NewRecorder()
		json.NewEncoder(w).Encode(expectedResponse)
		r := httptest.NewRequest(http.MethodPut, "/todo/markTodo", &body)

		s.mockHandler.EXPECT().MarkTodo(w, r).Return()
		s.mockHandler.MarkTodo(w, r)

		resp, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			t.Error("Unresolved Json Body")
		}

		cleanResponse := model.GetTodoResponse{}
		err = json.Unmarshal([]byte(resp), &cleanResponse)
		if err != nil {
			t.Error("Json Response Doesn't Match")
		}

		s.Assertions.Equal(expectedResponse, cleanResponse)
		s.Assertions.Equal(http.StatusOK, w.Result().StatusCode)
	})

	s.T().Run("TestHandlerGivenMarkedIdWhenMarkTodoIsCalledThenItSouldReturnGetTodoResponseWithSuccessTrueAndTodoMarkedSuccessfullyMessageAndDataWithTodo.MarkedAs1AndCodeAs200", func(t *testing.T) {
		todo := &model.Todo{
			ID:     1,
			Value:  "buy some milk",
			Marked: 1,
		}

		var body bytes.Buffer
		json.NewEncoder(&body).Encode(todo)

		todo.Marked = 0
		expectedResponse := util.SetAndGetTodoResponse(true, "Todo Marked Successfully", todo, 200)

		w := httptest.NewRecorder()
		json.NewEncoder(w).Encode(expectedResponse)
		r := httptest.NewRequest(http.MethodPut, "/todo/markTodo", &body)

		s.mockHandler.EXPECT().MarkTodo(w, r).Return()
		s.mockHandler.MarkTodo(w, r)

		resp, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			t.Error("Unresolved Json Body")
		}

		cleanResponse := model.GetTodoResponse{}
		err = json.Unmarshal([]byte(resp), &cleanResponse)
		if err != nil {
			t.Error("Json Response Doesn't Match")
		}

		s.Assertions.Equal(expectedResponse, cleanResponse)
		s.Assertions.Equal(http.StatusOK, w.Result().StatusCode)
	})
}

func (s *HandlerSuite) TestHandlerGivenWhenDeleteTodoIsCalled() {
	s.T().Run("TestHandlerGivenExistingIdWhenDeleteTodoIsCalledThenItSouldReturnGetTodoResponseWithSuccessTrueAndTodoDeletedSuccessfullyMessageAndDataWithEmptyTodoAndCodeAs200", func(t *testing.T) {
		expectedResponse := util.SetAndGetTodoResponse(true, "Todo Deleted Successfully", nil, 200)

		w := httptest.NewRecorder()
		json.NewEncoder(w).Encode(expectedResponse)
		r := httptest.NewRequest(http.MethodDelete, "/todo/deleteTodo/1", nil)

		s.mockHandler.EXPECT().DeleteTodo(w, r).Return()
		s.mockHandler.DeleteTodo(w, r)

		resp, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			t.Error("Unresolved Json Body")
		}

		cleanResponse := model.GetTodoResponse{}
		err = json.Unmarshal([]byte(resp), &cleanResponse)
		if err != nil {
			t.Error("Json Response Doesn't Match")
		}

		s.Assertions.Equal(expectedResponse, cleanResponse)
		s.Assertions.Equal(http.StatusOK, w.Result().StatusCode)
	})

	s.T().Run("TestHandlerGivenUnExistingIdWhenDeleteTodoIsCalledThenItSouldReturnGetTodoResponseWithSuccessFalseAndNoTodoFoundWithId1MessageAndDataWithEmptyTodoAndCodeAs400", func(t *testing.T) {
		expectedResponse := util.SetAndGetTodoResponse(false, "no todo found with id:1", nil, 400)

		w := httptest.NewRecorder()
		json.NewEncoder(w).Encode(expectedResponse)
		r := httptest.NewRequest(http.MethodDelete, "/todo/deleteTodo/1", nil)

		s.mockHandler.EXPECT().DeleteTodo(w, r).Return()
		s.mockHandler.DeleteTodo(w, r)

		resp, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			t.Error("Unresolved Json Body")
		}

		cleanResponse := model.GetTodoResponse{}
		err = json.Unmarshal([]byte(resp), &cleanResponse)
		if err != nil {
			t.Error("Json Response Doesn't Match")
		}

		s.Assertions.Equal(expectedResponse, cleanResponse)
		s.Assertions.Equal(http.StatusOK, w.Result().StatusCode)
	})
}
