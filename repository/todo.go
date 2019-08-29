package repository

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/infinityc2/go-app/model"
)

type TodoRepository interface {
	GetAllTodo() ([]model.Todo, error)
	AddTodo(todo model.Todo) error
	UpdateTodo(todoID string, todo model.Todo) error
	DeleteTodo(todoID string) error
}

type TodoHandler struct {
	ConnectionDB *mgo.Session
}

const(
	DBName 		= "dashboard"
	collection 	= "todo"
)

func (h TodoHandler) GetAllTodo() ([]model.Todo, error) {
	var todos []model.Todo
	err := h.ConnectionDB.DB(DBName).C(collection).Find(nil).All(&todos)
	return todos, err
}
func (h TodoHandler) AddTodo(todo model.Todo) error {
	return h.ConnectionDB.DB(DBName).C(collection).Insert(todo)
}
func (h TodoHandler) UpdateTodo(todoID string, todo model.Todo) error {
	objectID := bson.ObjectIdHex(todoID)
	update := bson.M{"$set": bson.M{"title": todo.Title}}
	return h.ConnectionDB.DB(DBName).C(collection).UpdateId(objectID, update)
}

func (h TodoHandler) DeleteTodo(todoID string) error {
	objectID := bson.ObjectIdHex(todoID)
	return h.ConnectionDB.DB(DBName).C(collection).RemoveId(objectID)
}