package http

import (
	"net/http"
	"strconv"
	"todos/domain"

	"github.com/labstack/echo/v4"
)

// todoHandler : http handler for todos
type todoHandler struct {
	TodoUsecase domain.TodoUsecase
}

// NewtodoHandler will initialize the todos/ resources endpoint
func NewtodoHandler(e *echo.Echo, us domain.TodoUsecase) {
	handler := &todoHandler{
		TodoUsecase: us,
	}
	e.GET("/todos", handler.FetchTodos)
	e.POST("/todos", handler.CreateTodo)
	e.GET("/todos/:id", handler.GetTodo)
	e.PUT("/todos", handler.UpdateTodo)
	e.DELETE("/todos/:id", handler.DeleteTodo)
}

// CreateTodo : Creates a new todo in the database
func (uh *todoHandler) CreateTodo(c echo.Context) error {
	var todo domain.Todo
	err := c.Echo().Binder.Bind(&todo, c)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = uh.TodoUsecase.CreateTodo(todo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Success")
}

// FetchTodos : Gets all the todo details from the database
func (uh *todoHandler) FetchTodos(c echo.Context) error {

	todos, err := uh.TodoUsecase.FetchTodos()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todos)
}

// GetTodo : Gets a todo details from the database
func (uh *todoHandler) GetTodo(c echo.Context) error {

	idString := c.Param("id")
	var id int
	if len(idString) == 0 {
		return c.JSON(http.StatusOK, domain.MIssingTodoID)
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusOK, domain.InvalidTodoID)
	}
	todoResponse, err := uh.TodoUsecase.GetTodo(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todoResponse)
}

// UpdateTodo : Updates a todo in the database
func (uh *todoHandler) UpdateTodo(c echo.Context) error {
	var todo domain.Todo
	err := c.Echo().Binder.Bind(&todo, c)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = uh.TodoUsecase.UpdateTodo(todo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Todo Details updated successfully")
}

// DeleteTodo : deletes a todo from the database
func (uh *todoHandler) DeleteTodo(c echo.Context) error {
	var todo domain.Todo
	err := c.Echo().Binder.Bind(&todo, c)
	if err != nil {
		c.Echo().Logger.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	idString := c.Param("id")
	var id int
	if len(idString) == 0 {
		return c.JSON(http.StatusOK, domain.MIssingTodoID)
	}

	id, err = strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusOK, domain.InvalidTodoID)
	}

	err = uh.TodoUsecase.DeleteTodo(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Successfully DeleteTodod Todo ID:"+idString)
}
