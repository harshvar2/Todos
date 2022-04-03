package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todos/domain"
	"todos/domain/mocks"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestFetchTodos(t *testing.T) {
	var mockResponse domain.Todo
	err := faker.FakeData(&mockResponse)
	assert.NoError(t, err)

	e := echo.New()

	mockUCase := new(mocks.TodoUsecase)
	mockUCase.On("FetchTodos").Return([]domain.Todo{}, nil)

	req, err := http.NewRequest(echo.GET, "/todos", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos")

	handler := TodoHandler{
		TodoUsecase: mockUCase,
	}
	err = handler.FetchTodos(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotNil(t, rec.Body)
	mockUCase.AssertExpectations(t)
}

func TestGetTodo(t *testing.T) {
	var mockResponse domain.Todo
	err := faker.FakeData(&mockResponse)
	assert.NoError(t, err)

	e := echo.New()

	mockUCase := new(mocks.TodoUsecase)
	mockUCase.On("GetTodo", mock.AnythingOfType("int")).Return(domain.Todo{}, nil)

	req, err := http.NewRequest(echo.GET, "/todos/:id", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	handler := TodoHandler{
		TodoUsecase: mockUCase,
	}
	err = handler.GetTodo(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotNil(t, rec.Body)
	mockUCase.AssertExpectations(t)
}

func TestCreateTodo(t *testing.T) {
	var mockResponse domain.Todo
	err := faker.FakeData(&mockResponse)
	assert.NoError(t, err)

	e := echo.New()

	mockUCase := new(mocks.TodoUsecase)
	mockUCase.On("CreateTodo", mock.AnythingOfType("domain.Todo")).Return(nil)

	var jsonStr = []byte(`{
        "id": 121,
        "name": "Todo",
        "dob": "19-09-1999",
        "address": "india",
        "createdAt": "2022-02-13T23:41:12+05:30",
        "updatedAt": "2022-02-13T23:41:12+05:30"
    }`)
	reqBody := bytes.NewBuffer(jsonStr)

	req, err := http.NewRequest(echo.POST, "/todos", reqBody)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos")

	handler := TodoHandler{
		TodoUsecase: mockUCase,
	}
	err = handler.CreateTodo(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotNil(t, rec.Body)
	mockUCase.AssertExpectations(t)
}

func TestUpdateTodo(t *testing.T) {
	var mockResponse domain.Todo
	err := faker.FakeData(&mockResponse)
	assert.NoError(t, err)

	e := echo.New()

	mockUCase := new(mocks.TodoUsecase)
	mockUCase.On("UpdateTodo", mock.AnythingOfType("domain.Todo")).Return(nil)

	var jsonStr = []byte(`{
        "id": 121,
        "name": "Todo",
        "dob": "19-09-1999",
        "address": "india"
    }`)
	reqBody := bytes.NewBuffer(jsonStr)

	req, err := http.NewRequest(echo.PUT, "/todos", reqBody)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos")
	handler := TodoHandler{
		TodoUsecase: mockUCase,
	}
	err = handler.UpdateTodo(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotNil(t, rec.Body)
	mockUCase.AssertExpectations(t)
}

func TestDeleteTodo(t *testing.T) {
	var mockResponse domain.Todo
	err := faker.FakeData(&mockResponse)
	assert.NoError(t, err)

	e := echo.New()

	mockUCase := new(mocks.TodoUsecase)
	mockUCase.On("DeleteTodo", mock.AnythingOfType("int")).Return(nil)

	var jsonStr = []byte(`{
        "id": 121,
        "name": "Todo",
        "dob": "19-09-1999",
        "address": "india"
    }`)
	reqBody := bytes.NewBuffer(jsonStr)

	req, err := http.NewRequest(echo.DELETE, "/todos/:id", reqBody)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	handler := TodoHandler{
		TodoUsecase: mockUCase,
	}
	err = handler.DeleteTodo(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotNil(t, rec.Body)
	mockUCase.AssertExpectations(t)
}
