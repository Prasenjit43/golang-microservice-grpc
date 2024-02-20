package persistance

import (
	"github.com/Prasenjit43/golang-microservice-grpc/employee/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type IEmployeeDBContext interface {
	AddEmployee(emp *model.EmployeeModel) (*mongo.InsertOneResult, *model.ErrorDetail)
	UpdateEmployee(emp *model.EmployeeModel)(*mongo.UpdateResult, *model.ErrorDetail)
}
