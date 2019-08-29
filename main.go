package main

import(
	"log"
	"github.com/gin-contrib/cors"
	"github.com/infinityc2/go-app/repository"
	"github.com/infinityc2/go-app/controller"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

const (
	endpoint 	= "mongodb://localhost:27017"
	port 		= "localhost:8000"
)

func main() {
	connectionDB, err := mgo.Dial(endpoint)
	if err != nil {
		log.Fatal("Cannot connect database", err.Error())
	}
	router := gin.Default()
	setupRouter(router, connectionDB)
	router.Run(port)
}

func setupRouter(route *gin.Engine, connectionDB *mgo.Session) {
	todoRepository := repository.TodoHandler{
		ConnectionDB: connectionDB,
	}
	todoAPI := controller.TodoAPI{
		TodoRepository: &todoRepository,
	}
	route.Use(cors.Default())
	route.GET("/todo", todoAPI.TodoListHandler)
	route.POST("/todo", todoAPI.AddTodoHandler)
	route.PUT("/todo/:todo_id", todoAPI.UpdateTodoHandler)
	route.DELETE("/todo/:todo_id", todoAPI.DeleteTodoHandler)
}