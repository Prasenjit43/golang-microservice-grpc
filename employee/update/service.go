package update

import (
	"github.com/Prasenjit43/golang-microservice-grpc/employee/model"
	"github.com/Prasenjit43/golang-microservice-grpc/employee/persistance"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	repository persistance.IEmployeeDBContext
}

func InitService(repo persistance.IEmployeeDBContext) *Service {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Update(emp *model.EmployeeModel) (*mongo.UpdateResult, *model.ErrorDetail) {
	return s.repository.UpdateEmployee(emp)
}
