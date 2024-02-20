package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeModel struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name"`
	Role       string             `json:"role"`
	Department string             `json:"department"`
	UserId     string             `json:"userId"`
}
