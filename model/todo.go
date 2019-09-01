package model

import (
	"github.com/globalsign/mgo/bson"
)

// TodoModel ...
type TodoModel struct {
	Todo []Todo `json:"todos"`
}

// Todo Entity Todo create
type Todo struct {
	TodoID      bson.ObjectId 	`json:"todo_id" bson:"_id,omitempty"`
	Title  		string        	`json:"title" bson:"title"`
	Note   		string        	`json:"note" bson:"note"`
}