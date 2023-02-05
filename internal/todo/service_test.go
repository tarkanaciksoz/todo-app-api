package todo

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	mockService "github.com/tarkanaciksoz/api-todo-app/internal/mocks"
	"github.com/tarkanaciksoz/api-todo-app/internal/model"
)

type ServiceSuite struct {
	suite.Suite
	*require.Assertions
	ctrl *gomock.Controller

	mockService *mockService.MockService
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())

	s.mockService = mockService.NewMockService(s.ctrl)
}

func (s *ServiceSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *ServiceSuite) TestServiceGivenWhenServiceGetIsCalled() {
	s.T().Run("TestServiceGivenExistingTodoIdWhenGetIsCalledThenItShouldReturnTodoWithoutAnError", func(t *testing.T) {
		expectedResponse := &model.Todo{
			ID:     1,
			Value:  "buy some milk",
			Marked: 0,
		}

		s.mockService.EXPECT().Get(1).Return(expectedResponse, nil).Times(1)
		actualResponse, actualErr := s.mockService.Get(1)

		s.NoError(actualErr)
		s.Equal(expectedResponse, actualResponse)
	})

	s.T().Run("TestServiceGivenUnExistingTodoIdWhenGetIsCalledThenItShouldReturnEmptyTodoStructAndAnEror", func(t *testing.T) {
		expectedResponse := &model.Todo{}
		expectedErr := errors.New("no todo found with id:1")

		s.mockService.EXPECT().Get(1).Return(expectedResponse, expectedErr).Times(1)
		actualResponse, actualErr := s.mockService.Get(1)

		s.EqualError(actualErr, expectedErr.Error())
		s.Equal(expectedResponse, actualResponse)
	})
}

func (s *ServiceSuite) TestServiceGivenWhenListIsCalled() {
	s.T().Run("TestServiceGivenEmptyTodoListWhenListIsCalledThenItShouldReturnEmptyTodoList", func(t *testing.T) {
		expectedResponse := []*model.Todo{}
		s.mockService.EXPECT().List().Return(expectedResponse).Times(1)

		actualResponse := s.mockService.List()
		s.Equal(expectedResponse, actualResponse)
	})

	s.T().Run("TestServiceGivenTodoListWith2TodosWhenListIsCalledThenItShouldReturnTodoListWith2Todos", func(t *testing.T) {
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
		s.mockService.EXPECT().List().Return(expectedResponse).Times(1)

		actualResponse := s.mockService.List()
		s.Equal(expectedResponse, actualResponse)
	})
}

func (s *ServiceSuite) TestServiceGivenWhenCreateIsCalled() {
	s.T().Run("TestServiceGivenTodoWhenCreateIsCalledThenItShouldReturnAnTodoAndNilError", func(t *testing.T) {
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
		s.mockService.EXPECT().Create(createTodoRequest).Return(expectedResponse, nil).Times(1)

		actualResponse, actualErr := s.mockService.Create(createTodoRequest)

		s.NoError(actualErr)
		s.Equal(expectedResponse, actualResponse)
	})

	s.T().Run("TestServiceGivenTodoWithExistedTodoIdWhenCreateIsCalledThenItShouldReturnTodoWithNewIdAndNilError", func(t *testing.T) {
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
		s.mockService.EXPECT().Create(createTodoRequest).Return(expectedResponse, nil).Times(1)

		actualResponse, actualErr := s.mockService.Create(createTodoRequest)

		s.NoError(actualErr)
		s.Equal(expectedResponse, actualResponse)
	})
}

func (s *ServiceSuite) TestServiceGivenWhenMarkIsCalled() {
	s.T().Run("TestServiceGivenTodoWithUnMarkedWhenMarkIsCalledThenItShouldReturnTodoWithMarkedAndNilError", func(t *testing.T) {
		markTodoRequest := &model.Todo{
			ID:     1,
			Marked: 0,
		}
		expectedResponse := &model.Todo{
			ID:     1,
			Value:  "buy some milk",
			Marked: 1,
		}
		s.mockService.EXPECT().Mark(markTodoRequest).Return(expectedResponse, nil).Times(1)

		actualResponse, actualErr := s.mockService.Mark(markTodoRequest)
		s.NoError(actualErr)
		s.Equal(expectedResponse, actualResponse)
	})

	s.T().Run("TestServiceGivenUnExistingTodoWithUnMarkedWhenMarkCalledThenItShouldReturnTodoStructAndAnError", func(t *testing.T) {
		markTodoRequest := &model.Todo{
			ID:     1,
			Marked: 0,
		}
		expectedResponse := &model.Todo{}
		expectedErr := errors.New("no todo found with id:1")

		s.mockService.EXPECT().Mark(markTodoRequest).Return(expectedResponse, expectedErr).Times(1)
		actualResponse, actualErr := s.mockService.Mark(markTodoRequest)

		s.EqualError(actualErr, expectedErr.Error())
		s.Equal(expectedResponse, actualResponse)
	})

	s.T().Run("TestServiceGivenTodoWithMarkedWhenMarkIsCalledThenItShouldReturnTodoWithUnMarkedAndNilError", func(t *testing.T) {
		markTodoRequest := &model.Todo{
			ID:     1,
			Marked: 1,
		}
		expectedResponse := &model.Todo{
			ID:     1,
			Value:  "buy some milk",
			Marked: 0,
		}
		s.mockService.EXPECT().Mark(markTodoRequest).Return(expectedResponse, nil).Times(1)

		actualResponse, actualErr := s.mockService.Mark(markTodoRequest)
		s.NoError(actualErr)
		s.Equal(expectedResponse, actualResponse)
	})

	s.T().Run("TestServiceGivenUnExistingTodoWithMarkedWhenMarkIsCalledThenItShouldReturnTodoStructAndAnError", func(t *testing.T) {
		markTodoRequest := &model.Todo{
			ID:     1,
			Marked: 1,
		}
		expectedResponse := &model.Todo{}
		expectedErr := errors.New("no todo found with id:1")

		s.mockService.EXPECT().Mark(markTodoRequest).Return(expectedResponse, expectedErr).Times(1)
		actualResponse, actualErr := s.mockService.Mark(markTodoRequest)

		s.EqualError(actualErr, expectedErr.Error())
		s.Equal(expectedResponse, actualResponse)
	})
}

func (s *ServiceSuite) TestServiceGivenWhenDeleteIsCalled() {
	s.T().Run("TestServiceGivenExistingIdWhenDeleteIsCalledThenItShouldReturnNilError", func(t *testing.T) {
		s.mockService.EXPECT().Delete(1).Return(nil).Times(1)
		actualErr := s.mockService.Delete(1)
		s.NoError(actualErr)
	})

	s.T().Run("TestServiceGivenUnExistingIdWhenDeleteIsCalledThenItShouldReturnAnError", func(t *testing.T) {
		expectedError := errors.New("no todo found with id:1")

		s.mockService.EXPECT().Delete(1).Return(expectedError).Times(1)
		actualErr := s.mockService.Delete(1)

		s.EqualError(actualErr, expectedError.Error())
	})
}
