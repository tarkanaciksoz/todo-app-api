package todo

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	mockMemory "github.com/tarkanaciksoz/api-todo-app/internal/mocks"
	"github.com/tarkanaciksoz/api-todo-app/internal/model"
)

type DBSuite struct {
	suite.Suite
	*require.Assertions
	ctrl *gomock.Controller

	mockDB *mockMemory.MockDB
}

func TestDBSuite(t *testing.T) {
	suite.Run(t, new(DBSuite))
}

func (s *DBSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())

	s.mockDB = mockMemory.NewMockDB(s.ctrl)
}

func (s *DBSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *DBSuite) TestMemoryGivenWhenGetIsCalled() {
	s.T().Run("TestMemoryGivenExistingTodoIdWhenGetIsCalledThenItShouldReturnTodoWithoutAnError", func(t *testing.T) {
		expectedResponse := &model.Todo{
			ID:     1,
			Value:  "buy some milk",
			Marked: 0,
		}

		s.mockDB.EXPECT().Get(1).Return(expectedResponse, nil).Times(1)

		actualResponse, actualErr := s.mockDB.Get(1)

		s.NoError(actualErr)
		s.Equal(expectedResponse, actualResponse)
	})

	s.T().Run("TestMemoryGivenUnExistingTodoIdWhenGetIsCalledThenItShouldReturnEmptyTodoStructAndAnEror", func(t *testing.T) {
		expectedResponse := &model.Todo{}
		expectedErr := errors.New("no todo found with id:1")

		s.mockDB.EXPECT().Get(1).Return(expectedResponse, expectedErr).Times(1)
		actualResponse, actualErr := s.mockDB.Get(1)

		s.EqualError(actualErr, expectedErr.Error())
		s.Equal(expectedResponse, actualResponse)
	})
}

func (s *DBSuite) TestMemoryGivenWhenListIsCalled() {
	s.T().Run("TestMemoryGivenEmptyTodoListWhenListIsCalledThenItShouldReturnEmptyTodoList", func(t *testing.T) {
		expectedResponse := []*model.Todo{}
		s.mockDB.EXPECT().List().Return(expectedResponse).Times(1)

		actualResponse := s.mockDB.List()
		s.Equal(expectedResponse, actualResponse)
	})

	s.T().Run("TestMemoryGivenTodoListWith2TodosWhenListIsCalledThenItShouldReturnTodoListWith2Todos", func(t *testing.T) {
		expectedResponse := []*model.Todo{
			{
				ID:     1,
				Value:  "buy some milk",
				Marked: 0,
			},
			{
				ID:     2,
				Value:  "enjoy the assignment",
				Marked: 0,
			},
		}
		s.mockDB.EXPECT().List().Return(expectedResponse).Times(1)

		actualResponse := s.mockDB.List()
		s.Equal(expectedResponse, actualResponse)
	})
}

func (s *DBSuite) TestMemoryGivenWhenCreateIsCalled() {
	s.T().Run("TestMemoryGivenTodoWhenCreateIsCalledThenItShouldReturnAnTodoAndNilError", func(t *testing.T) {
		createTodoRequest := &model.Todo{
			ID:     1,
			Value:  "buy some milk",
			Marked: 0,
		}
		expectedResponse := &model.Todo{
			ID:     1,
			Value:  "buy some milk",
			Marked: 0,
		}
		s.mockDB.EXPECT().Create(createTodoRequest).Return(expectedResponse, nil).Times(1)

		actualResponse, actualErr := s.mockDB.Create(createTodoRequest)

		s.NoError(actualErr)
		s.Equal(expectedResponse, actualResponse)
	})

	s.T().Run("TestMemoryGivenTodoWithExistedTodoIdWhenCreateIsCalledThenItShouldReturnTodoWithNewIdAndNilError", func(t *testing.T) {
		createTodoRequest := &model.Todo{
			ID:     1,
			Value:  "buy some milk",
			Marked: 0,
		}
		expectedResponse := &model.Todo{
			ID:     2,
			Value:  "buy some milk",
			Marked: 0,
		}
		s.mockDB.EXPECT().Create(createTodoRequest).Return(expectedResponse, nil).Times(1)

		actualResponse, actualErr := s.mockDB.Create(createTodoRequest)

		s.NoError(actualErr)
		s.Equal(expectedResponse, actualResponse)
	})
}

func (s *DBSuite) TestMemoryGivenWhenMarkIsCalled() {
	s.T().Run("TestMemoryGivenTodoWithUnMarkedWhenMarkIsCalledThenItShouldReturnTodoWithMarkedAndNilError", func(t *testing.T) {
		markTodoRequest := &model.Todo{
			ID:     1,
			Marked: 0,
		}
		expectedResponse := &model.Todo{
			ID:     1,
			Value:  "buy some milk",
			Marked: 1,
		}
		s.mockDB.EXPECT().Mark(markTodoRequest).Return(expectedResponse, nil).Times(1)

		actualResponse, actualErr := s.mockDB.Mark(markTodoRequest)
		s.NoError(actualErr)
		s.Equal(expectedResponse, actualResponse)
	})

	s.T().Run("TestMemoryGivenUnExistingTodoWithUnMarkedWhenMarkCalledThenItShouldReturnTodoStructAndAnError", func(t *testing.T) {
		markTodoRequest := &model.Todo{
			ID:     1,
			Marked: 0,
		}
		expectedResponse := &model.Todo{}
		expectedErr := errors.New("no todo found with id:1")

		s.mockDB.EXPECT().Mark(markTodoRequest).Return(expectedResponse, expectedErr).Times(1)
		actualResponse, actualErr := s.mockDB.Mark(markTodoRequest)

		s.EqualError(actualErr, expectedErr.Error())
		s.Equal(expectedResponse, actualResponse)
	})

	s.T().Run("TestMemoryGivenTodoWithMarkedWhenMarkIsCalledThenItShouldReturnTodoWithUnMarkedAndNilError", func(t *testing.T) {
		markTodoRequest := &model.Todo{
			ID:     1,
			Marked: 1,
		}
		expectedResponse := &model.Todo{
			ID:     1,
			Value:  "buy some milk",
			Marked: 0,
		}
		s.mockDB.EXPECT().Mark(markTodoRequest).Return(expectedResponse, nil).Times(1)

		actualResponse, actualErr := s.mockDB.Mark(markTodoRequest)
		s.NoError(actualErr)
		s.Equal(expectedResponse, actualResponse)
	})

	s.T().Run("TestMemoryGivenUnExistingTodoWithMarkedWhenMarkIsCalledThenItShouldReturnTodoStructAndAnError", func(t *testing.T) {
		markTodoRequest := &model.Todo{
			ID:     1,
			Marked: 1,
		}
		expectedResponse := &model.Todo{}
		expectedErr := errors.New("no todo found with id:1")

		s.mockDB.EXPECT().Mark(markTodoRequest).Return(expectedResponse, expectedErr).Times(1)
		actualResponse, actualErr := s.mockDB.Mark(markTodoRequest)

		s.EqualError(actualErr, expectedErr.Error())
		s.Equal(expectedResponse, actualResponse)
	})
}

func (s *DBSuite) TestMemoryGivenWhenDeleteIsCalled() {
	s.T().Run("TestMemoryGivenExistingIdWhenDeleteIsCalledThenItShouldReturnNilError", func(t *testing.T) {
		s.mockDB.EXPECT().Delete(1).Return(nil).Times(1)
		actualErr := s.mockDB.Delete(1)
		s.NoError(actualErr)
	})

	s.T().Run("TestMemoryGivenUnExistingIdWhenDeleteIsCalledThenItShouldReturnAnError", func(t *testing.T) {
		expectedError := errors.New("no todo found with id:1")

		s.mockDB.EXPECT().Delete(1).Return(expectedError).Times(1)
		actualErr := s.mockDB.Delete(1)

		s.EqualError(actualErr, expectedError.Error())
	})
}
