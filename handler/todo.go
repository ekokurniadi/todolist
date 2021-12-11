package handler

import (
	"net/http"
	"strconv"

	"github.com/ekokurniadi/tokopedia-go-submittion/formatter"
	"github.com/ekokurniadi/tokopedia-go-submittion/helper"
	"github.com/ekokurniadi/tokopedia-go-submittion/input"
	"github.com/ekokurniadi/tokopedia-go-submittion/service"
	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	service service.TodoService
}

func NewTodoHandler(service service.TodoService) *todoHandler {
	return &todoHandler{service}
}
func (h *todoHandler) GetTodo(c *gin.Context) {
	var input input.InputIDTodo
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	todoDetail, err := h.service.TodoServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail Todo", http.StatusOK, "success", formatter.FormatTodo(todoDetail))
	c.JSON(http.StatusOK, response)
}

func (h *todoHandler) GetTodos(c *gin.Context) {
	todos, err := h.service.TodoServiceGetAll()
	if err != nil {
		response := helper.ApiResponse("Failed to get Todos", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of Todos", http.StatusOK, "success", formatter.FormatTodos(todos))
	c.JSON(http.StatusOK, response)
}
func (h *todoHandler) GetTodosInComplete(c *gin.Context) {
	todos, err := h.service.TodoServiceGetAllInComplete()
	if err != nil {
		response := helper.ApiResponse("Failed to get Todos", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of Todos", http.StatusOK, "success", formatter.FormatTodos(todos))
	c.JSON(http.StatusOK, response)
}

func (h *todoHandler) CreateTodo(c *gin.Context) {
	var input input.TodoInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create Todo failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newTodo, err := h.service.TodoServiceCreate(input)
	if err != nil {
		response := helper.ApiResponse("Create Todo failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Create Todo", http.StatusOK, "success", formatter.FormatTodo(newTodo))
	c.JSON(http.StatusOK, response)
}
func (h *todoHandler) UpdateTodo(c *gin.Context) {
	var inputID input.InputIDTodo
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Todos", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData input.TodoInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update Todo failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updatedTodo, err := h.service.TodoServiceUpdate(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to get Todos", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Update Todo", http.StatusOK, "success", formatter.FormatTodo(updatedTodo))
	c.JSON(http.StatusOK, response)
}
func (h *todoHandler) DeleteTodo(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var inputID input.InputIDTodo
	inputID.ID = id
	_, err := h.service.TodoServiceGetByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Todos", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.TodoServiceDeleteByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Todos", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Delete Todo", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *todoHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
