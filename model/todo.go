package model

import (
	"github.com/globalsign/mgo/bson"
)

type TodoModel struct {
	Todo []Todo `json:"todos"`
}

type Todo struct {
	TodoID      bson.ObjectId 	`json:"todo_id" bson:"_id,omitempty"`
	Title  		string        	`json:"title" bson:"title"`
	Note   		string        	`json:"note" bson:"note"`
}