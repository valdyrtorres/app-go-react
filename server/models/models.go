package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ToDoList é exportável
// evita o comentário exported type should have comment or be unexported go-lint
type ToDoList struct {
 
  ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
  Task   string             `json:"task,omitempty"`
  Status bool               `json:"status,omitempty"`
}