package controller

import(
	"github.com/infinityc2/go-app/model"
	"github.com/infinityc2/go-app/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TodoAPI struct {
	TodoRepository repository.TodoRepository
}

func (a TodoAPI) TodoListHandler(c *gin.Context) {
	var todoModel model.TodoModel
	todo, err := a.TodoRepository.GetAllTodo()
	if err != nil {
		log.Fatal("error TodoListHandler", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	todoModel.Todo = todo
	c.JSON(http.StatusOK, todoModel)
}

func (a TodoAPI) AddTodoHandler(c *gin.Context) {
	var todo model.Todo
	err := c.ShouldBindJSON(&todo)
	if a.TodoRepository.AddTodo(todo).Error(); err != nil {
		log.Fatal("error AddTodoHandler", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (a TodoAPI) UpdateTodoHandler(c *gin.Context) {
	var todo model.Todo
	todoID := c.Param("todo_id")
	err := c.ShouldBindJSON(&todo)
	if a.TodoRepository.UpdateTodo(todoID, todo).Error(); err != nil {
		log.Fatal("error UpdateTodoHandler", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (a TodoAPI) DeleteTodoHandler(c *gin.Context) {
	todoID := c.Param("todo_id")
	err := a.TodoRepository.DeleteTodo(todoID) 
	if err != nil {
		log.Fatal("error DeleteTodoHandler", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}